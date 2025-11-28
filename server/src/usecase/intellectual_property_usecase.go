package usecase

import (
	"context"
	"fmt"

	"kujicole/domain/model"
	"kujicole/domain/repository"
	"github.com/pkg/errors"
)

type intellectualPropertyUseCase struct {
	repos *repository.Repositories
}

var _ IntellectualPropertyUseCase = &intellectualPropertyUseCase{}

func NewIntellectualPropertyUseCase(repos *repository.Repositories) *intellectualPropertyUseCase {
	return &intellectualPropertyUseCase{
		repos: repos,
	}
}

func (u *intellectualPropertyUseCase) ListIntellectualPropertyCategories(ctx context.Context) ([]*model.IntellectualPropertyCategory, error) {
	return u.repos.IntellectualProperty.ListCategories(ctx)
}

func (u *intellectualPropertyUseCase) ListIntellectualPropertyRankGroups(ctx context.Context) ([]*model.IntellectualPropertyRankGroup, error) {
	return u.repos.IntellectualProperty.ListRankGroups(ctx)
}

func (u *intellectualPropertyUseCase) ListIntellectualPropertyRankGroupsByCategoryID(ctx context.Context, categoryID model.ID) ([]*model.IntellectualPropertyRankGroup, error) {
	return u.repos.IntellectualProperty.ListRankGroupsByCategoryID(ctx, categoryID)
}

func (u *intellectualPropertyUseCase) ListIntellectualProperties(ctx context.Context) ([]*model.IntellectualProperty, error) {
	return u.repos.IntellectualProperty.ListProperties(ctx)
}

func (u *intellectualPropertyUseCase) ListIntellectualPropertiesByRankGroupID(ctx context.Context, rankGroupID model.ID) ([]*model.IntellectualProperty, error) {
	return u.repos.IntellectualProperty.ListPropertiesByRankGroupID(ctx, rankGroupID)
}

func (u *intellectualPropertyUseCase) GetCategoryByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyCategory, error) {
	return u.repos.IntellectualProperty.GetCategoryByID(ctx, id)
}

func (u *intellectualPropertyUseCase) GetIntellectualPropertyCategoryByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyCategory, error) {
	return u.repos.IntellectualProperty.GetCategoryByID(ctx, id)
}

func (u *intellectualPropertyUseCase) GetRankGroupByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyRankGroup, error) {
	return u.repos.IntellectualProperty.GetRankGroupByID(ctx, id)
}

func (u *intellectualPropertyUseCase) DrawIntellectualProperty(ctx context.Context, input model.DrawIntellectualPropertyInput) ([]*model.IntellectualProperty, error) {
	// 在庫チェック
	// INPUTのCATEGORYIDから対象のcATEGORYと紐づくRANKgROUPと紐づくPROPERTYを取得して、在庫チェックを行う
	ipCategoryAggregate, err := u.repos.IntellectualProperty.GetIPCategoryAggregateByCategoryIDWithPessimisticLock(ctx, input.IPCategoryID)
	if err != nil {
		return nil, err
	}

	if !ipCategoryAggregate.CheckStock(input.DrawCount) {
		return nil, fmt.Errorf("在庫不足")
	}

	// 購入トランザクション作成
	ipCategoryPurchaseTransaction, err := model.NewIPCategoryPurchaseTransaction(ipCategoryAggregate.Category, input.DrawCount)
	if err != nil {
		return nil, err
	}

	// 決済処理

	// 決済が完了したら購入トランザクション更新
	ipCategoryPurchaseTransaction.UpdateStatusPaid()

	// 上記までで別APIに分割予定-------------------------------------

	// まだ引いていない対象カテゴリのトランザクションを取得
	// ipCategoryPurchaseTransaction, err := u.repos.IntellectualProperty.GetPurchaseTransactionsByCategoryIDAndStatus(ctx, input.IPCategoryID, "pending")
	// if err != nil {
	// 	return nil, err
	// }

	// 抽選処理（※購入履歴も作成して返す予定）
	results, err := ipCategoryAggregate.Draw(input.DrawCount)
	if err != nil {
		return nil, err
	}

	// 抽選が完了したらトランザクションを更新
	ipCategoryPurchaseTransaction.UpdateStatusDrawSuccess()

	err = repository.RunInTx(ctx, func(ctx context.Context) error {
		// 購入トランザクション作成
		// err, := u.repos.IntellectualProperty.CreateIPCategoryPurchaseTransaction(ctx, input.IPCategoryID, input.DrawCount)
		// if err != nil {
		// 	return err
		// }

		// aggregateの在庫を更新
		err = u.repos.IntellectualProperty.UpdateIPPropertiesStock(ctx, results)
		if err != nil {
			return errors.WithStack(err)
		}

		// 購入履歴作成
		// err = u.repos.IntellectualProperty.CreatePurchaseTransaction(ctx, input.IPCategoryID, input.DrawCount)
		// if err != nil {
		// 	return errors.WithStack(err)
		// }

		return nil
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 本当はくじを引ける最大個数を渡してクライアントが選択した個数分引いてくじを開くという流れになる予定だが、今回は簡略化のためにそのまま抽選結果を返す
	return results, nil
}

package mysql

import (
	"context"
	"fmt"

	"kujicole/domain/model"
	"kujicole/domain/repository"
	"kujicole/infra/bob/models"
	"kujicole/infra/mysql/convert"
	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/sm"
)

type intellectualPropertyRepository struct {
	db *bob.DB
}

var _ repository.IntellectualPropertyRepository = &intellectualPropertyRepository{}

func NewIntellectualPropertyRepository(db *bob.DB) *intellectualPropertyRepository {
	return &intellectualPropertyRepository{db: db}
}

func (r *intellectualPropertyRepository) ListCategories(ctx context.Context) ([]*model.IntellectualPropertyCategory, error) {
	categories, err := models.IntellectualPropertyCategories.Query().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertyCategorySliceToModelCategorySlice(categories)
}

func (r *intellectualPropertyRepository) ListRankGroups(ctx context.Context) ([]*model.IntellectualPropertyRankGroup, error) {
	rankGroups, err := models.IntellectualPropertyRankGroups.Query().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertyRankGroupSliceToModelRankGroupSlice(rankGroups)
}

func (r *intellectualPropertyRepository) ListRankGroupsByCategoryID(ctx context.Context, categoryID model.ID) ([]*model.IntellectualPropertyRankGroup, error) {
	rankGroups, err := models.IntellectualPropertyRankGroups.Query(
		sm.Where(models.IntellectualPropertyRankGroups.Columns.IPCategoryID.EQ(mysql.Arg(categoryID.String()))),
		sm.OrderBy(models.IntellectualPropertyRankGroups.Columns.EmissionRate).Asc(),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertyRankGroupSliceToModelRankGroupSlice(rankGroups)
}

func (r *intellectualPropertyRepository) ListProperties(ctx context.Context) ([]*model.IntellectualProperty, error) {
	properties, err := models.IntellectualProperties.Query().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertySliceToModelPropertySlice(properties)
}

func (r *intellectualPropertyRepository) ListPropertiesByRankGroupID(ctx context.Context, rankGroupID model.ID) ([]*model.IntellectualProperty, error) {
	properties, err := models.IntellectualProperties.Query(
		sm.Where(models.IntellectualProperties.Columns.IPRankGroupID.EQ(mysql.Arg(rankGroupID.String()))),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertySliceToModelPropertySlice(properties)
}

func (r *intellectualPropertyRepository) GetCategoryByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyCategory, error) {
	category, err := models.IntellectualPropertyCategories.Query(
		sm.Where(models.IntellectualPropertyCategories.Columns.ID.EQ(mysql.Arg(id.String()))),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertyCategoryToModelCategory(category)
}

func (r *intellectualPropertyRepository) GetRankGroupByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyRankGroup, error) {
	rankGroup, err := models.IntellectualPropertyRankGroups.Query(
		sm.Where(models.IntellectualPropertyRankGroups.Columns.ID.EQ(mysql.Arg(id.String()))),
	).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobIntellectualPropertyRankGroupToModelRankGroup(rankGroup)
}

func (r *intellectualPropertyRepository) GetIntellectualPropertyAggregateByCategoryID(ctx context.Context, categoryID model.ID) (*model.IPIntellectualPropertyAggregate, error) {
	// カテゴリーを取得
	category, err := r.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// カテゴリーに紐づくランクグループを取得
	rankGroups, err := r.ListRankGroupsByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// 各ランクグループに紐づくプロパティを取得
	var allProperties []*model.IntellectualProperty
	for _, rankGroup := range rankGroups {
		properties, err := r.ListPropertiesByRankGroupID(ctx, rankGroup.ID)
		if err != nil {
			return nil, err
		}
		allProperties = append(allProperties, properties...)
	}

	// 集約オブジェクトを作成
	aggregate := model.NewIPIntellectualPropertyAggregate(category, rankGroups, allProperties)
	return aggregate, nil
}

func (r *intellectualPropertyRepository) GetIntellectualPropertyAggregateByCategoryIDWithPessimisticLock(ctx context.Context, categoryID model.ID) (*model.IPIntellectualPropertyAggregate, error) {
	// カテゴリーを取得（読み取り専用、更新しないためロック不要）
	modelCategory, err := r.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// カテゴリーに紐づくランクグループを取得（読み取り専用、更新しないためロック不要）
	rankGroups, err := r.ListRankGroupsByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	// 各ランクグループに紐づくプロパティを取得（FOR UPDATEでロック、Drawメソッドでstockを更新するため）
	var allProperties []*model.IntellectualProperty
	for _, rankGroup := range rankGroups {
		propertiesBob, err := models.IntellectualProperties.Query(
			sm.Where(models.IntellectualProperties.Columns.IPRankGroupID.EQ(mysql.Arg(rankGroup.ID.String()))),
			sm.ForUpdate(),
		).All(ctx, r.db)
		if err != nil {
			return nil, err
		}
		properties, err := convert.BobIntellectualPropertySliceToModelPropertySlice(propertiesBob)
		if err != nil {
			return nil, err
		}
		allProperties = append(allProperties, properties...)
	}

	// 集約オブジェクトを作成
	aggregate := model.NewIPIntellectualPropertyAggregate(modelCategory, rankGroups, allProperties)
	return aggregate, nil
}

func (r *intellectualPropertyRepository) UpdateIntellectualPropertiesStock(ctx context.Context, drawnProperties []*model.IntellectualProperty) error {
	// Drawメソッドで選ばれたプロパティの在庫を更新
	// N+1問題が発生するので、修正予定
	for _, property := range drawnProperties {
		// プロパティの在庫を更新
		propertyBob, err := models.IntellectualProperties.Query(
			sm.Where(models.IntellectualProperties.Columns.ID.EQ(mysql.Arg(property.ID.String()))),
		).One(ctx, r.db)
		if err != nil {
			return fmt.Errorf("プロパティ取得エラー: %w", err)
		}

		newStock := propertyBob.Stock - 1
		if newStock < 0 {
			return fmt.Errorf("在庫が負の値になります: プロパティID %s", property.ID.String())
		}

		// プロパティの在庫を更新
		setter := &models.IntellectualPropertySetter{
			Stock: omit.From(newStock),
			UpdatedAt: omit.From(property.UpdatedAt),
		}
		err = propertyBob.Update(ctx, r.db, setter)
		if err != nil {
			return fmt.Errorf("プロパティ更新エラー: %w", err)
		}
	}

	return nil
}


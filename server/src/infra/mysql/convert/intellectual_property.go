package convert

import (
	"time"

	"kujicole/domain/model"
	"kujicole/infra/bob/enums"
	"kujicole/infra/bob/models"
	"kujicole/util"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/shopspring/decimal"
	"github.com/pkg/errors"
)

func BobIntellectualPropertyCategoryToModelCategory(category *models.IntellectualPropertyCategory) (*model.IntellectualPropertyCategory, error) {
	categoryID, err := model.IDFromString(category.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var precautions *string
	if !category.Precautions.IsNull() {
		val := category.Precautions.MustGet()
		precautions = &val
	}

	var imgURL *string
	if !category.ImgURL.IsNull() {
		val := category.ImgURL.MustGet()
		imgURL = &val
	}

	return &model.IntellectualPropertyCategory{
		ID:             categoryID,
		Name:           category.Name,
		Price:          category.Price.InexactFloat64(),
		SalesStartDate: category.SalesStartDate,
		SalesEndDate:   category.SalesEndDate,
		Fee:            category.Fee.InexactFloat64(),
		Precautions:   precautions,
		IsHidden:       category.IsHidden,
		ImgURL:         imgURL,
		CreatedAt:      category.CreatedAt,
		UpdatedAt:      category.UpdatedAt,
	}, nil
}

func BobIntellectualPropertyCategorySliceToModelCategorySlice(categories []*models.IntellectualPropertyCategory) ([]*model.IntellectualPropertyCategory, error) {
	return util.MapWithErr(categories, BobIntellectualPropertyCategoryToModelCategory)
}

func BobIntellectualPropertyRankGroupToModelRankGroup(rankGroup *models.IntellectualPropertyRankGroup) (*model.IntellectualPropertyRankGroup, error) {
	rankGroupID, err := model.IDFromString(rankGroup.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	categoryID, err := model.IDFromString(rankGroup.IPCategoryID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Convert enum
	var rank model.IntellectualPropertyRank
	switch rankGroup.Rank {
	case enums.IntellectualPropertyRankGroupsRankS:
		rank = model.IntellectualPropertyRankS
	case enums.IntellectualPropertyRankGroupsRankA:
		rank = model.IntellectualPropertyRankA
	case enums.IntellectualPropertyRankGroupsRankB:
		rank = model.IntellectualPropertyRankB
	case enums.IntellectualPropertyRankGroupsRankC:
		rank = model.IntellectualPropertyRankC
	case enums.IntellectualPropertyRankGroupsRankD:
		rank = model.IntellectualPropertyRankD
	case enums.IntellectualPropertyRankGroupsRankE:
		rank = model.IntellectualPropertyRankE
	default:
		rank = model.IntellectualPropertyRank(rankGroup.Rank.String())
	}

	var explanation *string
	if !rankGroup.Explanation.IsNull() {
		val := rankGroup.Explanation.MustGet()
		explanation = &val
	}

	var comments *string
	if !rankGroup.Comments.IsNull() {
		val := rankGroup.Comments.MustGet()
		comments = &val
	}

	var imgURL *string
	if !rankGroup.ImgURL.IsNull() {
		val := rankGroup.ImgURL.MustGet()
		imgURL = &val
	}

	return &model.IntellectualPropertyRankGroup{
		ID:           rankGroupID,
		IPCategoryID: categoryID,
		Name:         rankGroup.Name,
		Rank:         rank,
		EmissionRate: int(rankGroup.EmissionRate),
		Explanation:  explanation,
		IsHidden:     rankGroup.IsHidden,
		Comments:     comments,
		ImgURL:       imgURL,
		CreatedAt:    rankGroup.CreatedAt,
		UpdatedAt:    rankGroup.UpdatedAt,
	}, nil
}

func BobIntellectualPropertyRankGroupSliceToModelRankGroupSlice(rankGroups []*models.IntellectualPropertyRankGroup) ([]*model.IntellectualPropertyRankGroup, error) {
	return util.MapWithErr(rankGroups, BobIntellectualPropertyRankGroupToModelRankGroup)
}

func BobIntellectualPropertyToModelProperty(property *models.IntellectualProperty) (*model.IntellectualProperty, error) {
	propertyID, err := model.IDFromString(property.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ipRankGroupID, err := model.IDFromString(property.IPRankGroupID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var imgURL *string
	if !property.ImgURL.IsNull() {
		val := property.ImgURL.MustGet()
		imgURL = &val
	}

	return &model.IntellectualProperty{
		ID:        propertyID,
		IPRankGroupID: ipRankGroupID,
		Name:      property.Name,
		Stock:     int(property.Stock),
		IsHidden:  property.IsHidden,
		ImgURL:    imgURL,
		CreatedAt: property.CreatedAt,
		UpdatedAt: property.UpdatedAt,
	}, nil
}

func BobIntellectualPropertySliceToModelPropertySlice(properties []*models.IntellectualProperty) ([]*model.IntellectualProperty, error) {
	return util.MapWithErr(properties, BobIntellectualPropertyToModelProperty)
}

func ModelPurchaseHistoryToBobPurchaseHistorySetter(purchaseHistory *model.IntellectualPropertyPurchaseHistory) *models.PurchaseHistorySetter {
	return &models.PurchaseHistorySetter{
		ID:                     omit.From(purchaseHistory.ID.String()),
		PurchaseTransactionID:  omit.From(purchaseHistory.PurchaseTransactionID.String()),
		IntellectualPropertyID: omit.From(purchaseHistory.IntellectualPropertyID.String()),
		CreatedAt:              omit.From(purchaseHistory.CreatedAt),
		UpdatedAt:              omit.From(purchaseHistory.UpdatedAt),
	}
}

func BobPurchaseHistoryToModelPurchaseHistory(historyBob *models.PurchaseHistory) (*model.IntellectualPropertyPurchaseHistory, error) {
	historyID, err := model.IDFromString(historyBob.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	purchaseTransactionID, err := model.IDFromString(historyBob.PurchaseTransactionID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	intellectualPropertyID, err := model.IDFromString(historyBob.IntellectualPropertyID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.IntellectualPropertyPurchaseHistory{
		ID:                     historyID,
		PurchaseTransactionID:  purchaseTransactionID,
		IntellectualPropertyID: intellectualPropertyID,
		CreatedAt:              historyBob.CreatedAt,
		UpdatedAt:              historyBob.UpdatedAt,
	}, nil
}

func BobPurchaseHistorySliceToModelPurchaseHistorySlice(histories []*models.PurchaseHistory) ([]*model.IntellectualPropertyPurchaseHistory, error) {
	return util.MapWithErr(histories, BobPurchaseHistoryToModelPurchaseHistory)
}

func BobPurchaseTransactionToModelPurchaseTransaction(transactionBob *models.PurchaseTransaction) (*model.IntellectualPropertyPurchaseTransaction, error) {
	transactionID, err := model.IDFromString(transactionBob.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := model.IDFromString(transactionBob.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ipCategoryID, err := model.IDFromString(transactionBob.IPCategoryID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Convert enum
	var paymentMethod model.PurchaseTransactionPaymentMethod
	switch transactionBob.PaymentMethod {
	case enums.PurchaseTransactionsPaymentMethodCreditCard:
		paymentMethod = model.PurchaseTransactionPaymentMethodCreditCard
	case enums.PurchaseTransactionsPaymentMethodBankTransfer:
		paymentMethod = model.PurchaseTransactionPaymentMethodBankTransfer
	case enums.PurchaseTransactionsPaymentMethodOther:
		paymentMethod = model.PurchaseTransactionPaymentMethodOther
	default:
		paymentMethod = model.PurchaseTransactionPaymentMethod(transactionBob.PaymentMethod.String())
	}

	var status model.PurchaseTransactionStatus
	switch transactionBob.Status {
	case enums.PurchaseTransactionsStatusPaymentPending:
		status = model.PurchaseTransactionStatusPaymentPending
	case enums.PurchaseTransactionsStatusPaymentSuccess:
		status = model.PurchaseTransactionStatusPaymentSuccess
	case enums.PurchaseTransactionsStatusPaymentFailed:
		status = model.PurchaseTransactionStatusPaymentFailed
	case enums.PurchaseTransactionsStatusDrawSuccess:
		status = model.PurchaseTransactionStatusDrawSuccess
	case enums.PurchaseTransactionsStatusDrawFailed:
		status = model.PurchaseTransactionStatusDrawFailed
	default:
		status = model.PurchaseTransactionStatus(transactionBob.Status.String())
	}

	var providerTransactionID *string
	if !transactionBob.ProviderTransactionID.IsNull() {
		val := transactionBob.ProviderTransactionID.MustGet()
		providerTransactionID = &val
	}

	var paidAt *time.Time
	if !transactionBob.PaidAt.IsNull() {
		val := transactionBob.PaidAt.MustGet()
		paidAt = &val
	}

	return &model.IntellectualPropertyPurchaseTransaction{
		ID:                    transactionID,
		UserID:                userID,
		IPCategoryID:          ipCategoryID,
		PurchaseQuantity:      int(transactionBob.PurchaseQuantity),
		PurchasePrice:         transactionBob.PurchasePrice.InexactFloat64(),
		PaymentMethod:         paymentMethod,
		ProviderTransactionID: providerTransactionID,
		Status:                status,
		PaidAt:                paidAt,
		CreatedAt:             transactionBob.CreatedAt,
		UpdatedAt:             transactionBob.UpdatedAt,
	}, nil
}

func BobPurchaseTransactionSliceToModelPurchaseTransactionSlice(transactions []*models.PurchaseTransaction) ([]*model.IntellectualPropertyPurchaseTransaction, error) {
	return util.MapWithErr(transactions, BobPurchaseTransactionToModelPurchaseTransaction)
}

func ModelPurchaseTransactionToBobPurchaseTransactionSetter(transaction *model.IntellectualPropertyPurchaseTransaction) *models.PurchaseTransactionSetter {
	setter := &models.PurchaseTransactionSetter{
		ID:              omit.From(transaction.ID.String()),
		UserID:          omit.From(transaction.UserID.String()),
		IPCategoryID:    omit.From(transaction.IPCategoryID.String()),
		PurchaseQuantity: omit.From(int32(transaction.PurchaseQuantity)),
		PurchasePrice:   omit.From(decimal.NewFromFloat(transaction.PurchasePrice)),
		PaymentMethod:   omit.From(enums.PurchaseTransactionsPaymentMethod(transaction.PaymentMethod.String())),
		Status:          omit.From(enums.PurchaseTransactionsStatus(transaction.Status.String())),
		UpdatedAt:       omit.From(transaction.UpdatedAt),
	}
	if transaction.PaidAt != nil {
		setter.PaidAt = omitnull.From(*transaction.PaidAt)
	}
	if transaction.ProviderTransactionID != nil {
		setter.ProviderTransactionID = omitnull.From(*transaction.ProviderTransactionID)
	}
	return setter
}




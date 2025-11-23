package convert

import (
	"kujicole/domain/model"
	"kujicole/infra/bob/enums"
	"kujicole/infra/bob/models"
	"kujicole/util"
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


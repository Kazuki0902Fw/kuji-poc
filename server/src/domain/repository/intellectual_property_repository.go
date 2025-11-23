package repository

import (
	"context"

	"kujicole/domain/model"
)

type IntellectualPropertyRepository interface {
	ListCategories(ctx context.Context) ([]*model.IntellectualPropertyCategory, error)
	ListRankGroups(ctx context.Context) ([]*model.IntellectualPropertyRankGroup, error)
	ListRankGroupsByCategoryID(ctx context.Context, categoryID model.ID) ([]*model.IntellectualPropertyRankGroup, error)
	ListProperties(ctx context.Context) ([]*model.IntellectualProperty, error)
	ListPropertiesByRankGroupID(ctx context.Context, rankGroupID model.ID) ([]*model.IntellectualProperty, error)
	GetCategoryByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyCategory, error)
	GetRankGroupByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyRankGroup, error)
	GetIntellectualPropertyAggregateByCategoryID(ctx context.Context, categoryID model.ID) (*model.IPIntellectualPropertyAggregate, error)
	GetIntellectualPropertyAggregateByCategoryIDWithPessimisticLock(ctx context.Context, categoryID model.ID) (*model.IPIntellectualPropertyAggregate, error)
	UpdateIntellectualPropertiesStock(ctx context.Context, drawnProperties []*model.IntellectualProperty) error
}


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
	GetIPCategoryAggregateByCategoryID(ctx context.Context, categoryID model.ID) (*model.IPCategoryAggregate, error)
	GetIPCategoryAggregateByCategoryIDWithPessimisticLock(ctx context.Context, categoryID model.ID) (*model.IPCategoryAggregate, error)
	UpdateIPPropertiesStock(ctx context.Context, drawnProperties []*model.IntellectualProperty) error
}


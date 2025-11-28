package usecase

import (
	"context"

	"kujicole/domain/model"
)

type UseCases interface {
	AuthUseCase
	UserUseCase
	IntellectualPropertyUseCase
}

type AuthUseCase interface {
	Login(ctx context.Context, mailAddress string, password string) (*model.AuthToken, error)
	Logout(ctx context.Context, userID model.ID) error
	RefreshAccessToken(ctx context.Context, refreshToken string) (*model.AuthToken, error)
}

type UserUseCase interface {
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, userID model.ID) (*model.User, error)
}

type IntellectualPropertyUseCase interface {
	ListIntellectualPropertyCategories(ctx context.Context) ([]*model.IntellectualPropertyCategory, error)
	ListIntellectualPropertyRankGroups(ctx context.Context) ([]*model.IntellectualPropertyRankGroup, error)
	ListIntellectualPropertyRankGroupsByCategoryID(ctx context.Context, categoryID model.ID) ([]*model.IntellectualPropertyRankGroup, error)
	ListIntellectualProperties(ctx context.Context) ([]*model.IntellectualProperty, error)
	ListIntellectualPropertiesByRankGroupID(ctx context.Context, rankGroupID model.ID) ([]*model.IntellectualProperty, error)
	GetCategoryByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyCategory, error)
	GetIntellectualPropertyCategoryByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyCategory, error)
	GetRankGroupByID(ctx context.Context, id model.ID) (*model.IntellectualPropertyRankGroup, error)
	DrawIntellectualProperty(ctx context.Context, input model.DrawIntellectualPropertyInput) ([]*model.IntellectualProperty, error)
}

var _ UseCases = &useCases{}

type useCases struct {
	*authUseCase
	*userUseCase
	*intellectualPropertyUseCase
}

func NewUseCases(
	authUseCase *authUseCase,
	userUseCase *userUseCase,
	intellectualPropertyUseCase *intellectualPropertyUseCase,
) *useCases {
	return &useCases{
		authUseCase:   authUseCase,
		userUseCase:                userUseCase,
		intellectualPropertyUseCase: intellectualPropertyUseCase,
	}
}

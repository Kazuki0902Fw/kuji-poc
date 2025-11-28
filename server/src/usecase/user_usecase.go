package usecase

import (
	"context"

	"kujicole/domain/model"
	"kujicole/domain/repository"
)

type userUseCase struct {
	repos *repository.Repositories
}

var _ UserUseCase = &userUseCase{}

func NewUserUseCase(repos *repository.Repositories) *userUseCase {
	return &userUseCase{
		repos: repos,
	}
}

func (u *userUseCase) ListUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.repos.User.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, userID model.ID) (*model.User, error) {
	user, err := u.repos.User.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

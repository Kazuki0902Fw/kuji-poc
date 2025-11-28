package usecase

import (
	"context"

	"kujicole/domain/model"
	"kujicole/domain/repository"
	"kujicole/domain/service"
	"github.com/pkg/errors"
)

var _ AuthUseCase = &authUseCase{}

type authUseCase struct {
	repos            *repository.Repositories
	authTokenService service.AuthTokenService
}

func NewAuthUseCase(repos *repository.Repositories, authTokenService service.AuthTokenService) *authUseCase {
	return &authUseCase{
		repos:            repos,
		authTokenService: authTokenService,
	}
}

func (u *authUseCase) Login(ctx context.Context, mailAddress string, password string) (*model.AuthToken, error) {
	// NOTE: validation error のテスト用
	// if true {
	// 	return nil, errors.WithStack(domainErr.NewValidationError([]domainErr.ValidationErrorField{
	// 		{
	// 			Field:   "mailAddress",
	// 			Message: "test error",
	// 		},
	// 		{
	// 			Field:   "password",
	// 			Message: "test error",
	// 		},
	// 	}))
	// }

	user, err := u.repos.User.GetByMailAddressAndPassword(ctx, mailAddress, password)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	token, err := u.authTokenService.Create(ctx, user.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return token, nil
}

func (u *authUseCase) Logout(ctx context.Context, userID model.ID) error {
	err := u.authTokenService.DeleteRefreshTokenByUserID(ctx, userID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (u *authUseCase) RefreshAccessToken(ctx context.Context, refreshToken string) (*model.AuthToken, error) {
	token, err := u.authTokenService.UpdateByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return token, nil
}

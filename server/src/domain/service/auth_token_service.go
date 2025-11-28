package service

import (
	"context"

	"kujicole/domain/model"
)

type AuthTokenService interface {
	Create(ctx context.Context, userID model.ID) (*model.AuthToken, error)
	UpdateByRefreshToken(ctx context.Context, refreshToken string) (*model.AuthToken, error)
	VerifyAccessToken(ctx context.Context, token string) (*model.ID, error)
	DeleteRefreshTokenByUserID(ctx context.Context, userID model.ID) error
}

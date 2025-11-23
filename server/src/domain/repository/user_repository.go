package repository

import (
	"context"

	"kujicole/domain/model"
)

type UserRepository interface {
	// GetByLoginIDAndPassword(ctx context.Context, loginID string, password string) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
}

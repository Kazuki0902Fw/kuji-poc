package repository

import (
	"context"

	"kujicole/domain/model"
)

type UserRepository interface {
	GetByMailAddressAndPassword(ctx context.Context, mailAddress string, password string) (*model.User, error)
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, userID model.ID) (*model.User, error)
}

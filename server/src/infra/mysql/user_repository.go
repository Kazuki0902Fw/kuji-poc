package mysql

import (
	"context"

	"kujicole/domain/model"
	"kujicole/domain/repository"
	"kujicole/infra/bob/models"
	"kujicole/infra/mysql/convert"
	"github.com/stephenafamo/bob"
)

type userRepository struct {
	db *bob.DB
}

var _ repository.UserRepository = &userRepository{}

func NewUserRepository(db *bob.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	users, err := models.Users.Query().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return convert.BobUserSliceToModelUserSlice(users)
}

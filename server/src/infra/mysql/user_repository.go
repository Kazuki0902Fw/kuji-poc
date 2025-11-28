package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"kujicole/domain/model"
	"kujicole/domain/repository"
	"kujicole/infra/bob/models"
	"kujicole/infra/mysql/convert"
	"kujicole/util"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/sm"
	"github.com/pkg/errors"
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

func (r *userRepository) GetByMailAddressAndPassword(ctx context.Context, mailAddress string, password string) (*model.User, error) {
	userRecord, err := models.Users.Query(
		sm.Where(models.Users.Columns.MailAddress.EQ(mysql.Arg(mailAddress))),
	).One(ctx, r.db)

	if err == sql.ErrNoRows {
		return nil, errors.WithStack(fmt.Errorf("user not found by mail_address. mail_address: %s", mailAddress))
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = util.ComparePassword(userRecord.PasswordHash, password)
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf("password is incorrect. mail_address: %s", mailAddress))
	}

	user, err := convert.BobUserToModelUser(userRecord)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, userID model.ID) (*model.User, error) {
	userRecord, err := models.Users.Query(
		sm.Where(models.Users.Columns.ID.EQ(mysql.Arg(userID))),
	).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, errors.WithStack(fmt.Errorf("user not found by id. id: %s", userID))
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return convert.BobUserToModelUser(userRecord)
}

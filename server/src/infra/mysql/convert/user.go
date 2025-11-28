package convert

import (
	"kujicole/domain/model"
	"kujicole/infra/bob/enums"
	"kujicole/infra/bob/models"
	"kujicole/util"
	"github.com/pkg/errors"
)

func BobUserToModelUser(user *models.User) (*model.User, error) {
	userID, err := model.IDFromString(user.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var gender model.Gender
	for _, v := range enums.AllUsersGender() {
		if v.String() == user.Gender.String() {
			gender = model.Gender(v.String())
			break
		}
	}
	return &model.User{
		ID:    		      userID,
		Nickname:       user.Nickname,
		PasswordHash:   user.PasswordHash,
		Birthdate:      user.Birthdate,
		Gender:         gender,
		IsAdmin:        user.IsAdmin,
		IsMailmagazine: user.IsMailmagazine,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}, nil
}

func BobUserSliceToModelUserSlice(users []*models.User) ([]*model.User, error) {
	return util.MapWithErr(users, BobUserToModelUser)
}

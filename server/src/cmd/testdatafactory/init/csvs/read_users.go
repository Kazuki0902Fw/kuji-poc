package csvs

import (
	"strconv"
	"time"

	"kujicole/domain/model"
	"kujicole/infra/bob/enums"
	"kujicole/infra/bob/models"
	"kujicole/util"
	"github.com/aarondl/opt/omit"
)

func ReadUsers(
	usersCSVBytes []byte,
) ([]*models.UserSetter, error) {
	records, err := readCSV(usersCSVBytes)
	if err != nil {
		return nil, err
	}

	hash, _ := util.HashPassword("password")

	users := make([]*models.UserSetter, 0)
	for ix, record := range records {
		if ix == 0 {
			continue
		}

		userID := model.NewID().String()
		now := time.Now()

		// Parse birthdate
		birthdate, err := time.Parse("2006-01-02", record[2])
		if err != nil {
			return nil, err
		}

		// Parse gender
		gender := enums.UsersGender(record[3])
		if !gender.Valid() {
			return nil, err
		}

		// Parse boolean values
		isAdmin, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, err
		}
		isMailmagazine, err := strconv.ParseBool(record[5])
		if err != nil {
			return nil, err
		}

		users = append(users, &models.UserSetter{
			ID:             omit.From(userID),
			Nickname:       omit.From(record[1]),
			Password:       omit.From(hash),
			Birthdate:      omit.From(birthdate),
			Gender:         omit.From(gender),
			IsAdmin:        omit.From(isAdmin),
			IsMailmagazine: omit.From(isMailmagazine),
			CreatedAt:      omit.From(now),
			UpdatedAt:      omit.From(now),
		})
	}
	return users, nil
}

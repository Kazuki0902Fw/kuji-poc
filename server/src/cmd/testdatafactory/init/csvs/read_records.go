package csvs

import (
	"embed"

	"kujicole/infra/bob/models"
)

//go:embed *.csv
var files embed.FS

type Records struct {
	Users                        []*models.UserSetter
	IntellectualPropertyCategories []*models.IntellectualPropertyCategorySetter
	IntellectualPropertyRankGroups []*models.IntellectualPropertyRankGroupSetter
	IntellectualProperties        []*models.IntellectualPropertySetter
}

func ReadRecords() (*Records, error) {
	// Read users
	usersCSVBytes, err := files.ReadFile("users.csv")
	if err != nil {
		return nil, err
	}
	users, err := ReadUsers(usersCSVBytes)
	if err != nil {
		return nil, err
	}

	// Read intellectual property categories
	categoriesCSVBytes, err := files.ReadFile("intellectual_property_categories.csv")
	if err != nil {
		return nil, err
	}
	categories, err := ReadIntellectualPropertyCategories(categoriesCSVBytes)
	if err != nil {
		return nil, err
	}

	// Create map for category IDs (index -> ID)
	categoryIDMap := make(map[int]string)
	for i, category := range categories {
		if category.ID.IsValue() {
			categoryIDMap[i] = category.ID.MustGet()
		}
	}

	// Read intellectual property rank groups
	rankGroupsCSVBytes, err := files.ReadFile("intellectual_property_rank_groups.csv")
	if err != nil {
		return nil, err
	}
	rankGroups, err := ReadIntellectualPropertyRankGroups(rankGroupsCSVBytes, categoryIDMap)
	if err != nil {
		return nil, err
	}

	// Create map for rank group IDs (index -> ID)
	rankGroupIDMap := make(map[int]string)
	for i, rankGroup := range rankGroups {
		if rankGroup.ID.IsValue() {
			rankGroupIDMap[i] = rankGroup.ID.MustGet()
		}
	}

	// Read intellectual properties
	propertiesCSVBytes, err := files.ReadFile("intellectual_properties.csv")
	if err != nil {
		return nil, err
	}
	properties, err := ReadIntellectualProperties(propertiesCSVBytes, rankGroupIDMap)
	if err != nil {
		return nil, err
	}

	return &Records{
		Users:                        users,
		IntellectualPropertyCategories: categories,
		IntellectualPropertyRankGroups: rankGroups,
		IntellectualProperties:        properties,
	}, nil
}

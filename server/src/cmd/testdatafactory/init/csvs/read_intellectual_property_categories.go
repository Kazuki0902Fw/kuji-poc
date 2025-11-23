package csvs

import (
	"strconv"
	"time"

	"kujicole/domain/model"
	"kujicole/infra/bob/models"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/shopspring/decimal"
)

func ReadIntellectualPropertyCategories(
	categoriesCSVBytes []byte,
) ([]*models.IntellectualPropertyCategorySetter, error) {
	records, err := readCSV(categoriesCSVBytes)
	if err != nil {
		return nil, err
	}

	categories := make([]*models.IntellectualPropertyCategorySetter, 0)
	for ix, record := range records {
		if ix == 0 {
			continue
		}

		categoryID := model.NewID().String()
		now := time.Now()

		// Parse price
		price, err := decimal.NewFromString(record[2])
		if err != nil {
			return nil, err
		}

		// Parse sales_start_date
		salesStartDate, err := time.Parse("2006-01-02", record[3])
		if err != nil {
			return nil, err
		}

		// Parse sales_end_date
		salesEndDate, err := time.Parse("2006-01-02", record[4])
		if err != nil {
			return nil, err
		}

		// Parse fee
		fee, err := decimal.NewFromString(record[5])
		if err != nil {
			return nil, err
		}

		// Parse is_hidden
		isHidden, err := strconv.ParseBool(record[7])
		if err != nil {
			return nil, err
		}

		category := &models.IntellectualPropertyCategorySetter{
			ID:             omit.From(categoryID),
			Name:           omit.From(record[1]),
			Price:          omit.From(price),
			SalesStartDate: omit.From(salesStartDate),
			SalesEndDate:   omit.From(salesEndDate),
			Fee:            omit.From(fee),
			IsHidden:       omit.From(isHidden),
			CreatedAt:      omit.From(now),
			UpdatedAt:      omit.From(now),
		}

		// Optional fields
		if record[6] != "" {
			category.Precautions = omitnull.From(record[6])
		}
		if record[8] != "" {
			category.ImgURL = omitnull.From(record[8])
		}

		categories = append(categories, category)
	}
	return categories, nil
}


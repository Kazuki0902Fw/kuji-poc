package csvs

import (
	"strconv"
	"time"

	"kujicole/domain/model"
	"kujicole/infra/bob/models"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
)

func ReadIntellectualProperties(
	propertiesCSVBytes []byte,
	ipRankGroupIDMap map[int]string,
) ([]*models.IntellectualPropertySetter, error) {
	records, err := readCSV(propertiesCSVBytes)
	if err != nil {
		return nil, err
	}

	properties := make([]*models.IntellectualPropertySetter, 0)
	for ix, record := range records {
		if ix == 0 {
			continue
		}

		propertyID := model.NewID().String()
		now := time.Now()

		// Get ip_rank_group_id from map
		// 各ランクグループに順番にプロパティを割り当てる
		// プロパティの行インデックス（ヘッダーを除く）をランクグループ数で割った余りで分配
		rowIndex := ix - 1 // ヘッダーを除いた行インデックス（0始まり）
		rankGroupIndex := rowIndex % len(ipRankGroupIDMap)
		ipRankGroupID, ok := ipRankGroupIDMap[rankGroupIndex]
		if !ok {
			// If not in map, use the value from CSV (if provided)
			if record[1] != "" {
				ipRankGroupID = record[1]
			} else {
				continue // Skip if no rank group ID
			}
		}

		// Parse stock
		stock, err := strconv.ParseInt(record[3], 10, 32)
		if err != nil {
			return nil, err
		}

		// Parse is_hidden
		isHidden, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, err
		}

		property := &models.IntellectualPropertySetter{
			ID:            omit.From(propertyID),
			IPRankGroupID: omit.From(ipRankGroupID),
			Name:          omit.From(record[2]),
			Stock:         omit.From(int32(stock)),
			IsHidden:      omit.From(isHidden),
			CreatedAt:     omit.From(now),
			UpdatedAt:     omit.From(now),
		}

		// Optional fields
		if record[5] != "" {
			property.ImgURL = omitnull.From(record[5])
		}

		properties = append(properties, property)
	}
	return properties, nil
}


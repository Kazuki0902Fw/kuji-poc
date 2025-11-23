package csvs

import (
	"strconv"
	"time"

	"kujicole/domain/model"
	"kujicole/infra/bob/enums"
	"kujicole/infra/bob/models"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
)

func ReadIntellectualPropertyRankGroups(
	rankGroupsCSVBytes []byte,
	ipCategoryIDMap map[int]string,
) ([]*models.IntellectualPropertyRankGroupSetter, error) {
	records, err := readCSV(rankGroupsCSVBytes)
	if err != nil {
		return nil, err
	}

	rankGroups := make([]*models.IntellectualPropertyRankGroupSetter, 0)
	for ix, record := range records {
		if ix == 0 {
			continue
		}

		rankGroupID := model.NewID().String()
		now := time.Now()

		// Get ip_category_id from map
		// カテゴリ1（インデックス0）: 最初の5行（S, A, B, C, D）を紐づける
		// カテゴリ2（インデックス1）: 残りの行を紐づける
		var categoryIndex int
		rowIndex := ix - 1 // ヘッダーを除いた行インデックス（0始まり）
		if rowIndex < 5 {
			// 最初の5行（行2-6）をカテゴリ1に紐づける
			categoryIndex = 0
		} else {
			// 残りの行（行7-11）をカテゴリ2に紐づける
			categoryIndex = 1
		}
		
		ipCategoryID, ok := ipCategoryIDMap[categoryIndex]
		if !ok {
			// If not in map, use the value from CSV (if provided)
			if record[1] != "" {
				ipCategoryID = record[1]
			} else {
				continue // Skip if no category ID
			}
		}

		// Parse rank
		rank := enums.IntellectualPropertyRankGroupsRank(record[3])
		if !rank.Valid() {
			return nil, err
		}

		// Parse emission_rate
		emissionRate, err := strconv.ParseInt(record[4], 10, 32)
		if err != nil {
			return nil, err
		}

		// Parse is_hidden
		isHidden, err := strconv.ParseBool(record[6])
		if err != nil {
			return nil, err
		}

		rankGroup := &models.IntellectualPropertyRankGroupSetter{
			ID:           omit.From(rankGroupID),
			IPCategoryID: omit.From(ipCategoryID),
			Name:         omit.From(record[2]),
			Rank:         omit.From(rank),
			EmissionRate: omit.From(int32(emissionRate)),
			IsHidden:     omit.From(isHidden),
			CreatedAt:    omit.From(now),
			UpdatedAt:    omit.From(now),
		}

		// Optional fields
		// record[5] = explanation
		if record[5] != "" {
			rankGroup.Explanation = omitnull.From(record[5])
		}
		// record[7] = comments
		if record[7] != "" {
			rankGroup.Comments = omitnull.From(record[7])
		}
		// record[8] = img_url
		if record[8] != "" {
			rankGroup.ImgURL = omitnull.From(record[8])
		}

		rankGroups = append(rankGroups, rankGroup)
	}
	return rankGroups, nil
}


package model

import (
	"fmt"
	"math/rand"
)

type IPCategoryAggregate struct {
	Category *IntellectualPropertyCategory
	RankGroups []*IntellectualPropertyRankGroup
}

const UPPER_LIMIT = 100

func NewIPCategoryAggregate(category *IntellectualPropertyCategory, rankGroups []*IntellectualPropertyRankGroup, properties []*IntellectualProperty) *IPCategoryAggregate {
	// Propertiesを各RankGroupに割り当てる
	rankGroupMap := make(map[ID]*IntellectualPropertyRankGroup)
	for _, rankGroup := range rankGroups {
		rankGroupMap[rankGroup.ID] = rankGroup
		// 初期化（nilの可能性があるため）
		if rankGroup.Properties == nil {
			rankGroup.Properties = make([]*IntellectualProperty, 0)
		}
	}
	
	// 各Propertyを対応するRankGroupに割り当て
	for _, property := range properties {
		if rankGroup, ok := rankGroupMap[property.IPRankGroupID]; ok {
			rankGroup.Properties = append(rankGroup.Properties, property)
		}
	}
	
	return &IPCategoryAggregate{
		Category: category,
		RankGroups: rankGroups,
	}
}

func (ipa *IPCategoryAggregate) GetTotalStock() int {
	totalStock := 0
	for _, rankGroup := range ipa.RankGroups {
		for _, property := range rankGroup.Properties {
			totalStock += property.Stock
		}
	}
	return totalStock
}

func (ipa *IPCategoryAggregate) CheckStock(drawCount int) bool {
	return ipa.GetTotalStock() >= drawCount
}

// ストックがあるプロパティを取得
func (ipa *IPCategoryAggregate) findPropertiesByRankGroupIdInStock(rankGroupId ID) []*IntellectualProperty {
	properties := make([]*IntellectualProperty, 0)
	for _, rankGroup := range ipa.RankGroups {
		if rankGroup.ID.String() == rankGroupId.String() {
			for _, property := range rankGroup.Properties {
				if property.Stock > 0 {
					properties = append(properties, property)
				}
			}
			break
		}
	}
	return properties
}

// utilに移動してもいいかも
func (ipa *IPCategoryAggregate) getRandomIndex(upperLimit int) int {
	return rand.Intn(upperLimit)
}

func (ipa *IPCategoryAggregate) Draw(drawCount int) ([]*IntellectualProperty, error) {
	results := make([]*IntellectualProperty, 0, drawCount)
	for i := 0; i < drawCount; i++ {
		// くじコレの仕様として在庫数/100の確率で当たりを引けるようにしてるっぽいので、100の乱数を生成
		randomDrawIndex := ipa.getRandomIndex(UPPER_LIMIT) + 1
		fmt.Printf("randomDrawIndex: %d\n", randomDrawIndex)
		// 各ランクの排出確率に対して当たったかどうかのフラグ
		hit := false
		// くじを引くごとに在庫が足りてるかチェック（※Drawメソッドが走る前に在庫チェックを行ってるので通らないはず）
		if ipa.GetTotalStock() == 0 || (drawCount - i) > ipa.GetTotalStock() {
			return nil, fmt.Errorf("乱数のインデックスが在庫数を超えた")
		}
		// 乱数のインデックスから排出確率が低いランクグループから順にチェックしていき、乱数のインデックスが排出確率を超えたらそのランクグループのプロパティを抽選結果に追加
		for _, rankGroup := range ipa.RankGroups {
			fmt.Printf("rankGroup.EmissionRate: %d\n", rankGroup.EmissionRate)
			fmt.Printf("rankGroup.Rank: %s\n", rankGroup.Rank)
			// TODO: この辺の処理はいい感じにmodel分割してメソッドを定義したい
			if rankGroup.EmissionRate >= randomDrawIndex {
				rankGroupProperties := ipa.findPropertiesByRankGroupIdInStock(rankGroup.ID)
				if len(rankGroupProperties) == 0 {
					fmt.Printf("在庫があるプロパティが見つからないので次のランクグループのチェックに進む: ランクグループID %s\n", rankGroup.ID.String())
					// 在庫があるプロパティが見つからない場合次のランクグループのチェックに進む
					continue
				}
				randomPropertyIndex := rand.Intn(len(rankGroupProperties))
				property := rankGroupProperties[randomPropertyIndex]
				results = append(results, property)
				property.decreaseStock()				
				hit = true
				break
			}
		}
		if !hit {
			// 排出確率を逆順にして、在庫があるプロパティを抽選結果に追加
			for i := len(ipa.RankGroups) - 1; i >= 0; i-- {
				rankGroup := ipa.RankGroups[i]
				rankGroupProperties := ipa.findPropertiesByRankGroupIdInStock(rankGroup.ID)
				if len(rankGroupProperties) == 0 {
					fmt.Printf("在庫があるプロパティが見つからないので次のランクグループのチェックに進む: ランクグループID %s\n", rankGroup.ID.String())
					// 在庫があるプロパティが見つからない場合次のランクグループのチェックに進む
					continue
				}
				randomPropertyIndex := rand.Intn(len(rankGroupProperties))
				property := rankGroupProperties[randomPropertyIndex]
				results = append(results, property)
				property.decreaseStock()
				break // プロパティを見つけたらループを抜ける
			}
		}
	}
	return results, nil
}

package model

import (
	"fmt"
	"math/rand"
)

type IPIntellectualPropertyAggregate struct {
	Category *IntellectualPropertyCategory
	RankGroups []*IntellectualPropertyRankGroup
	Properties []*IntellectualProperty
	TotalStock int
}

const UPPER_LIMIT = 100

func NewIPIntellectualPropertyAggregate(category *IntellectualPropertyCategory, rankGroups []*IntellectualPropertyRankGroup, properties []*IntellectualProperty) *IPIntellectualPropertyAggregate {
	totalStock := 0
	for _, property := range properties {
		totalStock += property.Stock
	}
	return &IPIntellectualPropertyAggregate{
		Category: category,
		RankGroups: rankGroups,
		Properties: properties,
		TotalStock: totalStock,
	}
}

func (a *IPIntellectualPropertyAggregate) CheckStock(drawCount int) bool {
	return a.TotalStock >= drawCount
}

// ストックがあるプロパティを取得
func (a *IPIntellectualPropertyAggregate) findPropertiesByRankGroupIdInStock(rankGroupId ID) []*IntellectualProperty {
	properties := make([]*IntellectualProperty, 0)
	for _, property := range a.Properties {
		if property.IPRankGroupID.String() == rankGroupId.String() && property.Stock > 0 {
			properties = append(properties, property)
		}
	}
	return properties
}

func (a *IPIntellectualPropertyAggregate) getRandomIndex(upperLimit int) int {
	return rand.Intn(upperLimit)
}

func (a *IPIntellectualPropertyAggregate) Draw(drawCount int) ([]*IntellectualProperty, error) {
	results := make([]*IntellectualProperty, 0, drawCount)
	for i := 0; i < drawCount; i++ {
		// くじコレの仕様として在庫数/100の確率で当たりを引けるようにしてるっぽいので、100の乱数を生成
		randomDrawIndex := a.getRandomIndex(UPPER_LIMIT) + 1
		fmt.Printf("randomDrawIndex: %d\n", randomDrawIndex)
		// 各ランクの排出確率に対して当たったかどうかのフラグ
		hit := false
		// くじを引くごとに在庫が足りてるかチェック（※Drawメソッドが走る前に在庫チェックを行ってるので通らないはず）
		if a.TotalStock == 0 || (drawCount - i) > a.TotalStock {
			return nil, fmt.Errorf("乱数のインデックスが在庫数を超えた")
		}
		// 乱数のインデックスから排出確率が低いランクグループから順にチェックしていき、乱数のインデックスが排出確率を超えたらそのランクグループのプロパティを抽選結果に追加
		for _, rankGroup := range a.RankGroups {
			fmt.Printf("rankGroup.EmissionRate: %d\n", rankGroup.EmissionRate)
			fmt.Printf("rankGroup.Rank: %s\n", rankGroup.Rank)
			// TODO: この辺の処理はいい感じにmodel分割してメソッドを定義したい
			if rankGroup.EmissionRate >= randomDrawIndex {
				rankGroupProperties := a.findPropertiesByRankGroupIdInStock(rankGroup.ID)
				if len(rankGroupProperties) == 0 {
					fmt.Printf("在庫があるプロパティが見つからないので次のランクグループのチェックに進む: ランクグループID %s\n", rankGroup.ID.String())
					// 在庫があるプロパティが見つからない場合次のランクグループのチェックに進む
					continue
				}
				randomPropertyIndex := rand.Intn(len(rankGroupProperties))
				property := rankGroupProperties[randomPropertyIndex]
				results = append(results, property)
				property.Stock--				
				a.TotalStock--
				hit = true
				break
			}
		}
		if !hit {
			// 排出確率を逆順にして、在庫があるプロパティを抽選結果に追加
			for i := len(a.RankGroups) - 1; i >= 0; i-- {
				rankGroup := a.RankGroups[i]
				rankGroupProperties := a.findPropertiesByRankGroupIdInStock(rankGroup.ID)
				if len(rankGroupProperties) == 0 {
					fmt.Printf("在庫があるプロパティが見つからないので次のランクグループのチェックに進む: ランクグループID %s\n", rankGroup.ID.String())
					// 在庫があるプロパティが見つからない場合次のランクグループのチェックに進む
					continue
				}
				randomPropertyIndex := rand.Intn(len(rankGroupProperties))
				property := rankGroupProperties[randomPropertyIndex]
				results = append(results, property)
				property.Stock--
				a.TotalStock--
			}
		}
	}
	return results, nil
}

package queries

import (
	"fmt"
)

type Query string

const (
	OpenMenu     Query = "0"
	OpenMyHeroes Query = "1"
	OpenMyHero   Query = "2"

	HeroCreationStart       Query = "3"
	HeroCreationSelectClass Query = "4"

	OpenShop           Query = "5"
	ShopBuyItem        Query = "6"
	ShopBuyItemConfirm Query = "7"

	ChallengeBoss            Query = "8"
	OpenFightOptions         Query = "24"
	StartFight               Query = "9"
	OpenActiveFight          Query = "10"
	FightActionSimpleStrike  Query = "11"
	FightActionStrongStrike  Query = "12"
	FightActionPreciseStrike Query = "13"
	FightActionInventory     Query = "14"
	FightActionRunAway       Query = "15"

	OpenHeroEquipment        Query = "16"
	HeroEquipmentOpenItem    Query = "17"
	HeroEquipmentEquipItem   Query = "18"
	HeroEquipmentUnequipItem Query = "19"
	HeroEquipmentConsumeItem Query = "20"

	OpenLevelUpBonusesOverview Query = "21"
	SpendLevelUpBonus          Query = "22"

	HeroDelete Query = "23"

	Back Query = "9999"

	SpecialDelimeterInQueryCallback = "#_#"
)

func New(query Query, additionalFields ...string) Query {
	for _, field := range additionalFields {
		query += Query(fmt.Sprintf("%s%s", SpecialDelimeterInQueryCallback, field))
	}

	return query
}

func (i Query) With(additionalFields ...string) Query {
	return New(i, additionalFields...)
}

func (i Query) WithID(id int64) Query {
	return i + Query(fmt.Sprintf("%s%v", SpecialDelimeterInQueryCallback, id))
}

func FightQueries() []Query {
	return []Query{
		FightActionSimpleStrike,
		FightActionStrongStrike,
		FightActionPreciseStrike,
		FightActionRunAway,

		OpenActiveFight,
		FightActionInventory,
		OpenHeroEquipment,
		HeroEquipmentOpenItem,
		HeroEquipmentEquipItem,
		HeroEquipmentUnequipItem,
		HeroEquipmentConsumeItem,
	}
}

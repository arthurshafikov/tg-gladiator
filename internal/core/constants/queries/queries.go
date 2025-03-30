package queries

import (
	"fmt"
)

type Query string

// @todo prior to launch replace these text values with numbers 0-999
const (
	OpenMenu     Query = "OpenMenu"
	OpenMyHeroes Query = "OpenMyHeroes"
	OpenMyHero   Query = "OpenMyHero"

	HeroCreationStart       Query = "HeroCreationStart"
	HeroCreationSelectClass Query = "HeroCreationSelectClass"

	OpenShop           Query = "OpenShop"
	ShopBuyItem        Query = "ShopBuyItem"
	ShopBuyItemConfirm Query = "ShopBuyItemConfirm"

	ChallengeBoss            Query = "ChallengeBoss"
	StartTournamentFight     Query = "StartTournamentFight"
	FightActionSimpleStrike  Query = "FightActionSimpleStrike"
	FightActionStrongStrike  Query = "FightActionStrongStrike"
	FightActionPreciseStrike Query = "FightActionPreciseStrike"
	FightActionRunAway       Query = "FightActionRunAway"

	OpenHeroEquipment        Query = "OpenHeroEquipment"
	HeroEquipmentOpenItem    Query = "HeroEquipmentOpenItem"
	HeroEquipmentEquipItem   Query = "HeroEquipmentEquipItem"
	HeroEquipmentUnequipItem Query = "HeroEquipmentUnequipItem"

	OpenLevelUpBonusesOverview Query = "OpenLevelUpBonusesOverview"
	SpendLevelUpBonus          Query = "SpendLevelUpBonus"

	HeroDelete Query = "HeroDelete"

	Back Query = "Back"

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
	}
}

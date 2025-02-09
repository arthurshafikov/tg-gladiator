package models

import (
	"fmt"
	"time"
)

type Item struct {
	ID                         int64  `json:"id"`
	Name                       string `json:"name"`
	AttackBonus                int    `json:"attack_bonus"`
	DefenseBonus               int    `json:"defense_bonus"`
	CriticalChancePercentBonus int    `json:"critical_chance_percent_bonus"`
	EvasionPercentBonus        int    `json:"evasion_percent_bonus"`
	BasePrice                  int    `json:"base_price"`
}

// @todo should it be here?
func (i *Item) GetShortCharacteristicsText() string {
	text := ""

	if i.AttackBonus != 0 {
		if i.AttackBonus > 0 {
			text += fmt.Sprintf("+%v💪 ", i.AttackBonus)
		} else {
			text += fmt.Sprintf("%v💪 ", i.AttackBonus)
		}
	}

	if i.DefenseBonus != 0 {
		if i.DefenseBonus > 0 {
			text += fmt.Sprintf("+%v🛡 ", i.DefenseBonus)
		} else {
			text += fmt.Sprintf("%v🛡 ", i.DefenseBonus)
		}
	}

	if i.CriticalChancePercentBonus != 0 {
		if i.CriticalChancePercentBonus > 0 {
			text += fmt.Sprintf("+%v%%💥 ", i.CriticalChancePercentBonus)
		} else {
			text += fmt.Sprintf("%v%%💥 ", i.CriticalChancePercentBonus)
		}
	}

	if i.EvasionPercentBonus != 0 {
		if i.EvasionPercentBonus > 0 {
			text += fmt.Sprintf("+%v%%🍀 ", i.EvasionPercentBonus)
		} else {
			text += fmt.Sprintf("%v%%🍀 ", i.EvasionPercentBonus)
		}
	}

	return text
}

type HeroShop struct {
	ID        int64     `json:"id"`
	HeroID    int64     `json:"hero_id"`
	UpdatesAt time.Time `json:"updates_at"`
	CreatedAt time.Time `json:"created_at"`
}

type HeroShopItem struct {
	Item       Item  `gorm:"->"`
	ItemID     int64 `json:"item_id"`
	HeroShopID int64 `json:"hero_shop_id"`
	Price      int   `json:"price"`
}

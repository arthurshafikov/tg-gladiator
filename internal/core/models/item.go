package models

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

type Item struct {
	ID                         int64              `json:"id"`
	Category                   enums.ItemCategory `json:"category"`
	EquipsOn                   enums.ItemEquipsOn `json:"equips_on"`
	Name                       string             `json:"name"`
	AttackBonus                int                `json:"attack_bonus"`
	DefenseBonus               int                `json:"defense_bonus"`
	CriticalChancePercentBonus int                `json:"critical_chance_percent_bonus"`
	EvasionPercentBonus        int                `json:"evasion_percent_bonus"`
	BasePrice                  int                `json:"base_price"`
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

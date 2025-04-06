package models

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

const ItemFieldCategory = "category"

type Item struct {
	ID                         int64              `json:"id"`
	Category                   enums.ItemCategory `json:"category"`
	EquipsOn                   enums.ItemEquipsOn `json:"equips_on"`
	Name                       string             `json:"name"`
	AttackBonus                int                `json:"attack_bonus"`
	DefenseBonus               int                `json:"defense_bonus"`
	CriticalChancePercentBonus int                `json:"critical_chance_percent_bonus"`
	EvasionPercentBonus        int                `json:"evasion_percent_bonus"`
	PotionEffectType           string             `json:"potion_effect_type"`
	PotionEffectValue          int                `json:"potion_effect_value"`
	PotionEffectDuration       int                `json:"potion_effect_duration"`
	BasePrice                  int                `json:"base_price"`
}

func (i *Item) IsPotion() bool {
	return i.Category == enums.ItemCategoryPotion
}

func (i *Item) IsEquippable() bool {
	return string(i.EquipsOn) != ""
}

// @todo should it be here?
func (i Item) GetShortCharacteristicsText() string {
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

	if i.PotionEffectValue != 0 {
		var icon string
		switch i.PotionEffectType {
		case enums.PotionEffectTypeHeal:
			icon = "❤️"
			// case enums.PotionEffectTypeAttackBuff:
			// 	icon = "💪"
			// case enums.PotionEffectTypeDefenseBuff:
			// 	icon = "🛡️"
			// case enums.PotionEffectTypeEvasionBuff:
			// 	icon = "🍀"
		}

		var durationText string
		if i.PotionEffectDuration > 0 {
			durationText = fmt.Sprintf(
				" на %v ходов", // @todo убрать в config messages и там подставлять текст
				i.PotionEffectDuration,
			)
		}

		if i.PotionEffectValue > 0 {
			text += fmt.Sprintf("+%v%s%s", i.PotionEffectValue, icon, durationText)
		} else {
			text += fmt.Sprintf("%v%s%s", i.PotionEffectValue, icon, durationText)
		}
	}

	return text
}

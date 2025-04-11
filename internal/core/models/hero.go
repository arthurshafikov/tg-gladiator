package models

import (
	"math"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

const (
	HeroFieldCurrentGold                = "current_gold"
	HeroFieldXP                         = "xp"
	HeroFieldLevel                      = "level"
	HeroFieldLevelUpBonusesLeft         = "level_up_bonuses_left"
	LevelUpRequiredXPKoefficient        = 50
	HeroMaxEnergy                       = 15
	HeroMinutesToRestoreEnergy          = 5
	HeroFieldHealthBonus                = "hp_bonus"
	HeroFieldAttackBonus                = "attack_bonus"
	HeroFieldDefenseBonus               = "defense_bonus"
	HeroFieldCriticalChancePercentBonus = "critical_chance_percent_bonus"
	HeroFieldEvasionChancePercentBonus  = "evasion_chance_percent_bonus"
	HeroFieldEnergy                     = "current_energy"
)

type Hero struct {
	ID            int64           `gorm:"->" json:"id"`
	ChatID        int64           `json:"chat_id"`
	Name          string          `json:"name"`
	Class         enums.HeroClass `json:"class"`
	XP            int             `json:"xp"`
	Level         int             `json:"level"`
	CurrentHP     int             `json:"current_hp"`
	CurrentEnergy int             `json:"current_energy"`
	CurrentGold   int             `json:"current_gold"`

	AttackBonus                int `json:"attack_bonus"`
	DefenseBonus               int `json:"defense_bonus"`
	HpBonus                    int `json:"hp_bonus"`
	CriticalChancePercentBonus int `json:"critical_chance_percent_bonus"`
	EvasionChancePercentBonus  int `json:"evasion_chance_percent_bonus"`
	LevelUpBonusesLeft         int `json:"level_up_bonuses_left"`

	CreatedAt time.Time `gorm:"->" json:"created_at"`

	Equipment []Item `gorm:"-"`
}

func (Hero) TableName() string {
	return "heroes"
}

func (h *Hero) GetID() int64 {
	return h.ID
}

func (h *Hero) GetType() enums.FighterType {
	return enums.FighterTypeHero
}

func (h *Hero) GetName(enemyNames ...map[string]string) string {
	return h.Name
}

func (h *Hero) GetLevel() int {
	return h.Level
}

func (h *Hero) GetHP() int {
	return h.CurrentHP + h.HpBonus
}

func (h *Hero) HasAttackRange() bool {
	return h.GetMaxAttack() != h.GetMinAttack()
}

func (h *Hero) GetMinAttack() int {
	minAttack := h.Class.GetCharacteristics().Attack

	if len(h.Equipment) > 0 {
		for _, item := range h.Equipment {
			minAttack += item.AttackBonus
		}
	}

	return minAttack + h.AttackBonus
}

func (h *Hero) GetMaxAttack() int {
	return h.GetMinAttack()
}

func (h *Hero) GetDefense() int {
	defense := h.Class.GetCharacteristics().Defense

	if len(h.Equipment) > 0 {
		for _, item := range h.Equipment {
			defense += item.DefenseBonus
		}
	}

	return defense + h.DefenseBonus
}

func (h *Hero) GetCriticalChancePercent() int {
	criticalChancePercent := h.Class.GetCharacteristics().CriticalChancePercent

	if len(h.Equipment) > 0 {
		for _, item := range h.Equipment {
			criticalChancePercent += item.CriticalChancePercentBonus
		}
	}

	return criticalChancePercent + h.CriticalChancePercentBonus
}

func (h *Hero) GetEvasionChancePercent() int {
	evasionChancePercent := h.Class.GetCharacteristics().EvasionChancePercent

	if len(h.Equipment) > 0 {
		for _, item := range h.Equipment {
			evasionChancePercent += item.EvasionPercentBonus
		}
	}

	return evasionChancePercent + h.EvasionChancePercentBonus
}

func (h *Hero) GetMinGoldReward() int {
	return 0
}

func (h *Hero) GetMaxGoldReward() int {
	return h.GetMinGoldReward()
}

func (h *Hero) GetMinXPReward() int {
	return 0 // @todo
}

func (h *Hero) GetMaxXPReward() int {
	return h.GetMinXPReward()
}

func (h *Hero) HasLevelUpBonuses() bool {
	return h.LevelUpBonusesLeft > 0
}

func (h *Hero) GetXPForNextLevelLeft() int {
	requiredTotalXP := 0
	for level := 1; level <= h.Level; level++ {
		requiredTotalXP += level * level * LevelUpRequiredXPKoefficient
	}

	return int(math.Abs(float64(h.XP - requiredTotalXP)))
}

func (h *Hero) CalculateLevelForXP(xp int) int {
	for level := 1; ; level++ {
		requiredXPForTheNextLevel := level * level * LevelUpRequiredXPKoefficient

		xp = xp - requiredXPForTheNextLevel

		if xp < 0 {
			return level
		}
	}
}

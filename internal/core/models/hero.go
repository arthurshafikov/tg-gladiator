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

func (h *Hero) GetName() string {
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
		requiredXPForTheNextLevel := level * level * 50

		xp = xp - requiredXPForTheNextLevel

		if xp < 0 {
			return level
		}
	}
}

// For level - 2, required total XP - 50
// For level - 3, required total XP - 250
// For level - 4, required total XP - 700
// For level - 5, required total XP - 1500
// For level - 6, required total XP - 2750
// For level - 7, required total XP - 4550
// For level - 8, required total XP - 7000
// For level - 9, required total XP - 10200
// For level - 10, required total XP - 14250
// For level - 11, required total XP - 19250
// For level - 12, required total XP - 25300
// For level - 13, required total XP - 32500
// For level - 14, required total XP - 40950
// For level - 15, required total XP - 50750
// For level - 16, required total XP - 62000
// For level - 17, required total XP - 74800
// For level - 18, required total XP - 89250
// For level - 19, required total XP - 105450
// For level - 20, required total XP - 123500
// For level - 21, required total XP - 143500
// For level - 22, required total XP - 165550
// For level - 23, required total XP - 189750
// For level - 24, required total XP - 216200
// For level - 25, required total XP - 245000
// For level - 26, required total XP - 276250
// For level - 27, required total XP - 310050
// For level - 28, required total XP - 346500
// For level - 29, required total XP - 385700
// For level - 30, required total XP - 427750
// For level - 31, required total XP - 472750
// For level - 32, required total XP - 520800
// For level - 33, required total XP - 572000
// For level - 34, required total XP - 626450
// For level - 35, required total XP - 684250
// For level - 36, required total XP - 745500
// For level - 37, required total XP - 810300
// For level - 38, required total XP - 878750
// For level - 39, required total XP - 950950
// For level - 40, required total XP - 1027000

package models

import (
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

const HeroFieldCurrentGold = "current_gold"

type Hero struct {
	ID            int64           `gorm:"->" json:"id"`
	ChatID        int64           `json:"chat_id"`
	Name          string          `json:"name"`
	Class         enums.HeroClass `json:"class"`
	CurrentHP     int             `json:"current_hp"`
	CurrentEnergy int             `json:"current_energy"`
	CurrentGold   int             `json:"current_gold"`
	CreatedAt     time.Time       `gorm:"->" json:"created_at"`

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

func (h *Hero) GetHP() int {
	return h.CurrentHP
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

	return minAttack
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

	return defense
}

func (h *Hero) GetCriticalChancePercent() int {
	criticalChancePercent := h.Class.GetCharacteristics().CriticalChancePercent

	if len(h.Equipment) > 0 {
		for _, item := range h.Equipment {
			criticalChancePercent += item.CriticalChancePercentBonus
		}
	}

	return criticalChancePercent
}

func (h *Hero) GetEvasionChancePercent() int {
	evasionChancePercent := h.Class.GetCharacteristics().EvasionChancePercent

	if len(h.Equipment) > 0 {
		for _, item := range h.Equipment {
			evasionChancePercent += item.EvasionPercentBonus
		}
	}

	return evasionChancePercent
}

func (h *Hero) GetMinReward() int {
	return 0
}

func (h *Hero) GetMaxReward() int {
	return h.GetMinReward()
}

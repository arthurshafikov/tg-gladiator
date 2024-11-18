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

func (h *Hero) GetMinAttack() int {
	return h.Class.GetCharacteristics().Attack
}

func (h *Hero) GetMaxAttack() int {
	return h.Class.GetCharacteristics().Attack
}

func (h *Hero) GetDefense() int {
	return h.Class.GetCharacteristics().Defense
}

func (h *Hero) GetCriticalChancePercent() int {
	return h.Class.GetCharacteristics().CriticalChancePercent
}

func (h *Hero) GetEvasionChancePercent() int {
	return h.Class.GetCharacteristics().EvasionChancePercent
}

func (h *Hero) GetMinReward() int {
	return 0
}

func (h *Hero) GetMaxReward() int {
	return 0
}

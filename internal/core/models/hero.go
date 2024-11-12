package models

import (
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
)

type Hero struct {
	ID            int64           `gorm:"->" json:"id"`
	ChatID        int64           `json:"chat_id"`
	Class         enums.HeroClass `json:"class"`
	CurrentHP     int             `json:"current_hp"`
	CurrentEnergy int             `json:"current_energy"`
	CurrentGold   int             `json:"current_gold"`
	CreatedAt     time.Time       `gorm:"->" json:"created_at"`
}

func (Hero) TableName() string {
	return "heroes"
}

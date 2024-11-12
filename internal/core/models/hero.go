package models

import (
	"fmt"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
)

type Hero struct {
	ID            int64     `gorm:"->" json:"id"`
	ChatID        int64     `json:"chat_id"`
	Class         string    `json:"class"`
	CurrentHP     int64     `json:"current_hp"`
	CurrentEnergy int64     `json:"current_energy"`
	CurrentGold   int64     `json:"current_gold"`
	CreatedAt     time.Time `gorm:"->" json:"created_at"`
}

func (h *Hero) GetName(messages *config.Messages) string {
	return fmt.Sprintf(
		messages.HeroName,
		messages.Classes[h.Class],
		h.CurrentHP,
		h.CurrentGold,
	)
}

func (Hero) TableName() string {
	return "heroes"
}

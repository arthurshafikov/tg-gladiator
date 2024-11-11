package models

import (
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
)

type Chat struct {
	ID             int64
	Name           string                    `json:"name"`
	Username       string                    `json:"username"`
	Interaction    *interactions.Interaction `json:"interaction"`
	Language       string                    `json:"language"`
	LatestActiveAt time.Time                 `json:"latest_active_at"`
	CreatedAt      time.Time                 `json:"created_at"`
}

package models

import (
	"time"
)

type AnalyticEvent struct {
	ID        string    `json:"id" db:"id" gorm:"->"`
	ChatID    int64     `json:"chat_id" db:"chat_id"`
	Type      string    `json:"type" db:"type"`
	Payload   *string   `json:"payload" db:"payload" gorm:"type:jsonb"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

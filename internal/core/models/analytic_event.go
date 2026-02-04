package models

import "time"

type AnalyticEvent struct {
	ID        int       `json:"id" db:"id"`
	ChatID    int64     `json:"chat_id" db:"chat_id"`
	Type      string    `json:"type" db:"type"`
	Payload   any       `json:"payload" db:"payload"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

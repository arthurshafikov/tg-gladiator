package models

import (
	"time"
)

type AnalyticEventType string

const (
	AnalyticEventTypeChatUpdateProcessed AnalyticEventType = "chat_update_processed"
	AnalyticEventTypeFightStarted AnalyticEventType = "fight_started" // payload location: location from enum location_types
	AnalyticEventTypeItemPurchased AnalyticEventType = "item_purchased" // payload item_id: id of the purchased item
)

func (t AnalyticEventType) ToString() string {
	return string(t)
}

type AnalyticEvent struct {
	ID        string    `json:"id" db:"id" gorm:"->"`
	ChatID    int64     `json:"chat_id" db:"chat_id"`
	Type      AnalyticEventType `json:"type" db:"type"`
	Payload   *string   `json:"payload" db:"payload" gorm:"type:jsonb"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

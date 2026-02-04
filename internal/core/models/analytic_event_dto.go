package models

import (
	"time"
)

type CreateAnalyticEventDTO struct {
	ChatID    int64
	Type      AnalyticEventType
	Payload   map[string]any
	Timestamp time.Time
}



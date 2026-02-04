package models

import (
	"time"
)

type EventName string

const (
	EventHeroLevelUp EventName = "HeroLevelUp"
	EventAnalytic    EventName = "EventAnalytic"
)

type CreateAnalyticEventDTO struct {
	EventName EventName
	ChatID    int64
	Payload   map[string]any
	Timestamp time.Time
}

func (e EventName) ToString() string {
	return string(e)
}

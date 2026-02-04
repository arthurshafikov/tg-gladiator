package listeners

import (
	"context"
	"fmt"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
)

type AnalyticEventListener struct {
	service services.AnalyticEvent
}

func NewAnalyticEventListener(service services.AnalyticEvent) *AnalyticEventListener {
	return &AnalyticEventListener{service: service}
}

func (l *AnalyticEventListener) Handle(ctx context.Context, params ...any) error {
	if len(params) < 2 {
		return fmt.Errorf("params len < 2 AnalyticEventListener")
	}
	eventName, ok := params[0].(models.EventName)
	if !ok {
		return fmt.Errorf("params[0] is not EventName: %#v", params)
	}
	chatID, ok := params[1].(int64)
	if !ok {
		return fmt.Errorf("params[1] is not int64 (chatID): %#v", params)
	}
	var payload any
	if len(params) > 2 {
		payload = params[2]
	}

	return l.service.Create(ctx, models.CreateAnalyticEventDTO{
		EventName: eventName,
		ChatID:    chatID,
		Payload:   payload,
		Timestamp: time.Now(),
	})
}

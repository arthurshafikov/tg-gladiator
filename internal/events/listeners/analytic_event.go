package listeners

import (
	"context"
	"fmt"

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
	if len(params) != 1 {
		return fmt.Errorf("invalid params length: expected 1, got %d", len(params))
	}

	dto, ok := params[0].(models.CreateAnalyticEventDTO)
	if !ok {
		return fmt.Errorf("invalid params type: expected models.CreateAnalyticEventDTO, got %T", params)
	}

	return l.service.Create(ctx, dto)
}

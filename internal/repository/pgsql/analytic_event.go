package pgsql

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
)

type AnalyticEvent struct {
	BaseRepo
}

func NewAnalyticEventRepository(baseRepo BaseRepo) *AnalyticEvent {
	return &AnalyticEvent{
		BaseRepo: baseRepo,
	}
}

func (r *AnalyticEvent) Create(ctx context.Context, event models.AnalyticEvent) error {
	return r.getDBInstance(ctx).
		Table("analytic_events").
		Create(&event).Error
}

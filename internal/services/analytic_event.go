package services

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type AnalyticEventService struct {
	repo repository.AnalyticEvent
}

func NewAnalyticEventService(repo repository.AnalyticEvent) *AnalyticEventService {
	return &AnalyticEventService{repo: repo}
}

func (s *AnalyticEventService) Create(ctx context.Context, dto models.CreateAnalyticEventDTO) error {
	return s.repo.Create(ctx, models.AnalyticEvent{
		ChatID:    dto.ChatID,
		Type:      dto.EventName.ToString(),
		Payload:   dto.Payload,
		Timestamp: dto.Timestamp,
	})
}

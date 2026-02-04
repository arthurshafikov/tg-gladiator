package services

import (
	"context"
	"encoding/json"

	"github.com/arthurshafikov/tg-gladiator/internal/core/helpers"
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
	var payloadBytes []byte
	var err error
	if dto.Payload != nil {
		if payloadBytes, err = json.Marshal(dto.Payload); err != nil {
			return err
		}
	}

	analyticEvent := models.AnalyticEvent{
		ChatID:    dto.ChatID,
		Type:      dto.Type,
		Timestamp: dto.Timestamp,
	}

	if len(payloadBytes) > 0 {
		analyticEvent.Payload = helpers.GetPointer(string(payloadBytes))
	}

	return s.repo.Create(ctx, analyticEvent)
}

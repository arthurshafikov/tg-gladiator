package services

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type InteractionService struct {
	logger Logger
	repo   repository.Chat
}

func newInteractionService(logger Logger, repo repository.Chat) *InteractionService {
	return &InteractionService{
		logger: logger,
		repo:   repo,
	}
}

func (s *InteractionService) Find(ctx *types.Context) (*interactions.Interaction, error) {
	chat, err := s.repo.Find(ctx.GetContext(), ctx.GetChat().ID)
	if err != nil {
		if !errors.Is(err, errors.ErrNotFound) {
			s.logger.Error(err)

			return nil, errors.ErrServerError
		}

		return nil, errors.ErrNotFound
	}

	return chat.Interaction, nil
}

func (s *InteractionService) Set(ctx *types.Context, interaction interactions.Interaction) error {
	_, err := s.Find(ctx)
	if err != nil {
		return err
	}

	fields := &models.Chat{}
	if interaction == "" {
		fields.Interaction = nil
	} else {
		fields.Interaction = &interaction
	}

	if _, err = s.repo.UpdateMap(ctx.GetContext(), ctx.GetChat().ID, map[string]interface{}{
		"interaction": interaction,
	}); err != nil {
		if !errors.Is(err, errors.ErrNotFound) {
			s.logger.Error(err)

			return errors.ErrServerError
		}

		return errors.ErrNotFound
	}

	return nil
}

func (s *InteractionService) Remove(ctx *types.Context) error {
	if _, err := s.repo.UpdateMap(ctx.GetContext(), ctx.GetChat().ID, map[string]interface{}{
		"interaction": nil,
	}); err != nil {
		if !errors.Is(err, errors.ErrNotFound) {
			s.logger.Error(err)

			return errors.ErrServerError
		}

		return errors.ErrNotFound
	}

	return nil
}

func (s *InteractionService) RemoveIfEquals(ctx *types.Context, interaction *interactions.Interaction) error {
	currentInteraction, err := s.Find(ctx)
	if err != nil {
		return err
	}

	if currentInteraction != nil && interaction != nil && string(*currentInteraction) == string(*interaction) {
		return s.Remove(ctx)
	}

	return nil
}

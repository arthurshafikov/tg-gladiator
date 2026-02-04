package services

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/events"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/sirupsen/logrus"
)

type FightService struct {
	repo repository.Fight

	enemyService  *EnemyService
	heroService   *HeroService
	eventsHandler EventsHandler
}

func newFightService(repo repository.Fight, enemyService *EnemyService, heroService *HeroService, eventsHandler EventsHandler) *FightService {
	return &FightService{
		repo:          repo,
		enemyService:  enemyService,
		heroService:   heroService,
		eventsHandler: eventsHandler,
	}
}

func (s *FightService) FindActiveByHeroID(ctx *types.Context, heroID int64) (*models.Fight, error) {
	return s.findBy(ctx, &models.Fight{
		HeroID: heroID,
		Status: enums.FightStatusActive,
	})
}

func (s *FightService) FindMyActive(ctx *types.Context) (*models.Fight, error) {
	heroes, err := s.heroService.GetMy(ctx)
	if err != nil {
		return nil, err
	}

	heroIDs := make([]int64, 0, len(*heroes))
	for _, hero := range *heroes {
		heroIDs = append(heroIDs, hero.ID)
	}

	return s.findBy(ctx, map[string]interface{}{
		models.FightFieldHeroID: heroIDs,
		models.FightFieldStatus: enums.FightStatusActive,
	})
}

func (s *FightService) create(ctx *types.Context, fields models.Fight) (*models.Fight, error) {
	fight, err := s.repo.Create(ctx.GetContext(), fields)
	if err != nil {
		if errors.Is(err, errors.ErrAlreadyExists) {
			return nil, errors.ErrAlreadyExists
		}

		logrus.Error(err)

		return nil, err
	}

	s.eventsHandler.Dispatch(
		events.AnalyticEvent,
		models.CreateAnalyticEventDTO{
			ChatID:    ctx.GetChat().ID,
			Type:      models.AnalyticEventTypeFightStarted,
			Timestamp: fight.CreatedAt,
			Payload: map[string]any{
				"fight_id": fight.ID,
			},
		},
	)

	return s.findBy(ctx, &models.Fight{
		ID: fight.ID,
	})
}

func (s *FightService) update(ctx *types.Context, fightID int64, fields *models.Fight) (*models.Fight, error) {
	if err := s.repo.Update(ctx.GetContext(), fightID, fields); err != nil {
		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return s.findBy(ctx, &models.Fight{
		ID: fightID,
	})
}

func (s *FightService) updateMap(ctx *types.Context, fightID int64, fields map[string]interface{}) (*models.Fight, error) {
	if err := s.repo.UpdateMap(ctx.GetContext(), fightID, fields); err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrNotFound
		}

		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return s.findBy(ctx, &models.Fight{
		ID: fightID,
	})
}

func (s *FightService) findBy(ctx *types.Context, fields interface{}) (*models.Fight, error) {
	var fight *models.Fight
	var err error

	if _, ok := fields.(*models.Fight); ok {
		fight, err = s.repo.FindBy(ctx.GetContext(), fields.(*models.Fight))
	} else if _, ok := fields.(map[string]interface{}); ok {
		fight, err = s.repo.FindByMap(ctx.GetContext(), fields.(map[string]interface{}))
	} else {
		logrus.Error(fmt.Sprintf(
			"undefined fields model in fightService findBy: %#v",
			fields,
		))

		return nil, errors.ErrServerError
	}

	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return nil, errors.ErrNotFound
		}

		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	if fight.OpponentType == enums.OpponentTypeMob {
		enemy, err := s.enemyService.find(ctx, fight.OpponentID)
		if err != nil {
			return nil, err
		}

		fight.Opponent = enemy
	} else {
		logrus.Error(fmt.Sprintf("unsupported opponent type: %s", fight.OpponentType))

		return nil, errors.ErrServerError
	}

	fight.Hero, err = s.heroService.FindMy(ctx, fight.HeroID)
	if err != nil {
		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return fight, nil
}

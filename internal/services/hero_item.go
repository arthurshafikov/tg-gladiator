package services

import (
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/sirupsen/logrus"
)

type HeroItemService struct {
	repo repository.HeroItem
}

func newHeroItemService(repo repository.HeroItem) *HeroItemService {
	return &HeroItemService{
		repo: repo,
	}
}

func (s *HeroItemService) Create(ctx *types.Context, heroID int64, itemID int64) (*models.HeroItem, error) {
	heroItem, err := s.repo.Create(ctx.GetContext(), models.HeroItem{
		HeroID: heroID,
		ItemID: itemID,
	})
	if err != nil {
		logrus.Error(err)

		return nil, errors.ErrServerError
	}

	return heroItem, nil
}

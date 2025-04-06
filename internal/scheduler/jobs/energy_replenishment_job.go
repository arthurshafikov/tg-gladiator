package jobs

import (
	"context"
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/services"
	"github.com/sirupsen/logrus"
)

type EnergyReplenishmentJob struct {
	ctx context.Context

	services *services.Services
}

func NewEnergyReplenishmentJob(
	ctx context.Context,
	services *services.Services,
) *EnergyReplenishmentJob {
	return &EnergyReplenishmentJob{
		ctx:      ctx,
		services: services,
	}
}

func (j *EnergyReplenishmentJob) Handle() {
	if err := j.services.Heroes.ReplenishHeroesEnergy(j.ctx); err != nil {
		logrus.Error(fmt.Errorf("EnergyReplenishmentJob ReplenishHeroesEnergy err: %w", err))
	}
}

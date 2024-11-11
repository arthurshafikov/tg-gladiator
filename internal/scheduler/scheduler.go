package scheduler

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram"
	"github.com/go-co-op/gocron/v2"
)

type Deps struct {
	Logger               Logger
	Services             *services.Services
	Config               *config.Config
	NotificationsHandler telegram.NotificationsHandler
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
}

type Scheduler struct {
	ctx context.Context

	scheduler gocron.Scheduler

	logger               Logger
	services             *services.Services
	config               *config.Config
	notificationsHandler telegram.NotificationsHandler
}

func NewScheduler(ctx context.Context, deps Deps) (*Scheduler, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		ctx:       ctx,
		scheduler: scheduler,

		logger:               deps.Logger,
		services:             deps.Services,
		config:               deps.Config,
		notificationsHandler: deps.NotificationsHandler,
	}, nil
}

func (s *Scheduler) Start() error {
	s.initJobs()

	s.scheduler.Start()
	s.logger.Info("Scheduler started...")

	<-s.ctx.Done()
	s.logger.Info("Shutting down the scheduler...")
	err := s.scheduler.Shutdown()
	s.logger.Info("Scheduler has been stopped...")

	return err
}

func (s *Scheduler) initJobs() {
}

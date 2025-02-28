package events

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants"
	"github.com/arthurshafikov/tg-gladiator/internal/events/listeners"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram"
	"golang.org/x/sync/errgroup"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type RegisteredEventsMap map[string][]listeners.Listener

type Handler struct {
	ctx    context.Context
	group  *errgroup.Group
	logger Logger

	registeredEvents *RegisteredEventsMap
}

type Deps struct {
	NotificationsHandler telegram.NotificationsHandler
	Services             *services.Services
	Config               *config.Config
}

func NewHandler(
	ctx context.Context,
	group *errgroup.Group,
	logger Logger,
) *Handler {
	return &Handler{
		ctx:    ctx,
		group:  group,
		logger: logger,

		registeredEvents: &RegisteredEventsMap{},
	}
}

func (h *Handler) InitEvents(deps *Deps) {
	h.subscribe(constants.EventHeroLevelUp, listeners.NewHeroLevelUpListener(
		deps.NotificationsHandler,
		deps.Services.Chat,
		&deps.Config.MessagesBag,
	))
}

func (h *Handler) Dispatch(eventName string, params ...any) {
	listeners, ok := (*h.registeredEvents)[eventName]
	if !ok {
		h.logger.Error("no listeners registered for eventName " + eventName)
		return
	}

	h.group.Go(func() error {
		for _, listener := range listeners {
			if err := listener.Handle(h.ctx, params...); err != nil {
				h.logger.Error(err)
			}
		}

		return nil
	})
}

func (h *Handler) subscribe(eventName string, listener listeners.Listener) {
	(*h.registeredEvents)[eventName] = append((*h.registeredEvents)[eventName], listener)
}

package services

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
)

type Logger interface {
	Error(args ...interface{})
}

type EventsHandler interface {
	Dispatch(eventName string, params ...any)
}

type Chat interface {
	Find(ctx context.Context, id int64) (*models.Chat, error)
	FirstOrCreate(ctx context.Context, chat models.Chat) (*models.Chat, bool, error)
	UpdateLatestActiveAt(ctx *types.Context)
	UpdateUsername(ctx *types.Context, username string) error
}

type Interactions interface {
	Find(ctx *types.Context) (*interactions.Interaction, error)
	Set(ctx *types.Context, interaction interactions.Interaction) error
	Remove(ctx *types.Context) error
	RemoveIfEquals(ctx *types.Context, interaction *interactions.Interaction) error
}

type Services struct {
	Chat
	Interactions
}

type Deps struct {
	Repository *repository.Repository
	Logger
	Config        *config.Config
	EventsHandler EventsHandler
}

func NewServices(deps Deps) *Services {
	interactionService := newInteractionService(deps.Logger, deps.Repository.Chat)

	return &Services{
		Chat: newChatService(
			deps.Logger,
			deps.Repository.Chat,
		),
		Interactions: interactionService,
	}
}

package listeners

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
)

type Listener interface {
	Handle(ctx context.Context, params ...any) error
}

type NotificationsHandler interface {
	NotifyAboutHeroLevelUp(ctx *types.Context, hero *models.Hero) error
}

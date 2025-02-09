package services

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
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
	FindByChatID(ctx context.Context, chatID int64) (*models.Chat, error)
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

type Hero interface {
	Create(ctx *types.Context, name, class string) (*models.Hero, error)
	DeleteMy(ctx *types.Context, id int64) error
	FindMy(ctx *types.Context, id int64) (*models.Hero, error)
}

type Heroes interface {
	GetMy(ctx *types.Context) (*[]models.Hero, error)
}

type HeroShopItem interface {
	BuyItem(ctx *types.Context, heroID, itemID int64) (*models.HeroShopItem, error)
	FindMy(ctx *types.Context, heroID, itemID int64) (*models.HeroShopItem, error)
}

type TournamentFight interface {
	Create(ctx *types.Context, heroID int64) (*models.Fight, error)
	FindActiveByHeroID(ctx *types.Context, heroID int64) (*models.Fight, error)
	FindMyActive(ctx *types.Context) (*models.Fight, error)
	MakeTurn(
		ctx *types.Context,
		fight *models.Fight,
		actionType enums.FightActionType,
	) (*models.Fight, *models.FightEvents, error)
	RunAwayAsHero(ctx *types.Context, fightID int64) error
}

type Shop interface {
	GetShopItemsFor(ctx *types.Context, heroID int64) ([]models.HeroShopItem, error)
}

type Services struct {
	Chat
	Interactions
	Hero
	Heroes
	HeroShopItem
	TournamentFight
	Shop
}

type Deps struct {
	Repository *repository.Repository
	Logger
	Config        *config.Config
	EventsHandler EventsHandler
}

func NewServices(deps Deps) *Services {
	validatorService := newValidatorService(deps.Logger)

	heroService := newHeroService(deps.Logger, deps.Repository.Hero, validatorService)

	enemyService := newEnemyService(deps.Logger, deps.Repository.Enemy)

	tournamentFightEventsService := newTournamentFightEventsService(deps.Logger)

	heroItemService := newHeroItemService(deps.Repository.HeroItem)

	heroShopItemService := newHeroShopItemService(
		deps.Repository.HeroShopItem,
		deps.Repository.HeroShop,
		heroService,
		heroItemService,
	)

	return &Services{
		Chat: newChatService(
			deps.Logger,
			deps.Repository.Chat,
		),
		Interactions: newInteractionService(deps.Logger, deps.Repository.Chat),
		Hero:         heroService,
		Heroes:       newHeroesService(deps.Logger, deps.Repository.Hero),
		HeroShopItem: heroShopItemService,
		TournamentFight: newTournamentFightService(
			deps.Logger,
			deps.Repository.Fight,
			heroService,
			enemyService,
			tournamentFightEventsService,
		),
		Shop: newShopService(
			deps.Repository.Item,
			deps.Repository.HeroShop,
			deps.Repository.HeroShopItem,
		),
	}
}

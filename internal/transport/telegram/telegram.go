package telegram

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/commands"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/events"
	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	services      *services.Services
	logger        services.Logger
	config        *config.Config
	eventsHandler services.EventsHandler

	commandHandler     CommandsHandler
	queryHandler       QueryHandler
	interactionHandler InteractionHandler

	helper handlers.TelegramHandlerHelper

	chatMutexes sync.Map // map[chatID]*sync.Mutex
}

type BaseHandler interface {
	HandleShowActiveFightOverview(ctx *types.Context) error
}

type CommandsHandler interface {
	BaseHandler
	HandleStart(ctx *types.Context, message *tgbotapi.Message) error
	HandleHelp(ctx *types.Context) error
}

type QueryHandler interface {
	BaseHandler
	HandleOpenMenu(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleOpenMyHeroes(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleOpenMyHero(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleHeroCreationStart(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleHeroCreationSelectClass(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleChallengeBoss(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleOpenFightOptions(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleStartFight(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleOpenActiveFight(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleFightActionSimpleStrike(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleFightActionStrongStrike(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleFightActionPreciseStrike(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleFightActionRunAway(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleOpenShop(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleShopBuyItem(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleShopBuyItemConfirm(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleOpenHeroEquipment(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleHeroEquipmentOpenItem(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleHeroEquipmentUnequipItem(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleHeroEquipmentEquipItem(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleHeroEquipmentConsumeItem(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleOpenLevelUpBonusesOverview(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
	HandleSpendLevelUpBonus(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error

	HandleHeroDelete(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error
}

type InteractionHandler interface {
	BaseHandler
	HandleHeroCreationEnterName(ctx *types.Context, message *tgbotapi.Message, payload []string) error
	HandleHeroDelete(ctx *types.Context, message *tgbotapi.Message, payload []string) error
}

type NotificationsHandler interface {
	NotifyAboutHeroLevelUp(ctx *types.Context, hero *models.Hero) error
}

type Deps struct {
	Services      *services.Services
	Logger        services.Logger
	Config        *config.Config
	EventsHandler services.EventsHandler

	CommandsHandler    CommandsHandler
	QueryHandler       QueryHandler
	InteractionHandler InteractionHandler

	Helper handlers.TelegramHandlerHelper
}

func NewBot(deps *Deps) *Bot {
	return &Bot{
		services:      deps.Services,
		logger:        deps.Logger,
		config:        deps.Config,
		eventsHandler: deps.EventsHandler,

		commandHandler:     deps.CommandsHandler,
		queryHandler:       deps.QueryHandler,
		interactionHandler: deps.InteractionHandler,

		helper: deps.Helper,
	}
}

func (b *Bot) Start(ctx context.Context) error { // nolint:gocognit,gocyclo
	u := b.helper.NewUpdateChannel(0)
	u.Timeout = 60

	updates, err := b.helper.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	const workerCount = 32
	updateChan := make(chan tgbotapi.Update, workerCount*2)

	// Запускаем воркеры
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for update := range updateChan {
				b.ProcessUpdate(ctx, update)
			}
		}()
	}

	// Читаем апдейты и отправляем в канал
	for update := range updates {
		updateChan <- update
	}

	close(updateChan)
	wg.Wait()

	return nil
}

//nolint:gocognit
func (b *Bot) ProcessUpdate(ctx context.Context, update tgbotapi.Update) {
	var chatID int64
	switch {
	case update.Message != nil:
		chatID = update.Message.Chat.ID
	case update.CallbackQuery != nil:
		chatID = update.CallbackQuery.Message.Chat.ID
	case update.PreCheckoutQuery != nil:
		chatID = int64(update.PreCheckoutQuery.From.ID)
	default:
		b.logger.Error(fmt.Errorf("could not determine chat id in request: %#v", update))

		return
	}

	messages := b.config.MessagesBag.GetMessages("RU")

	chat, err := b.authorize(ctx, chatID)
	if err != nil && (update.CallbackQuery != nil || update.Message.Command() != commands.Start) {
		b.handleError(chatID, err, messages)

		return
	}

	if chat != nil {
		defer b.eventsHandler.Dispatch(
			events.AnalyticEvent,
			models.CreateAnalyticEventDTO{
				ChatID:    chat.ID,
				Type:      models.AnalyticEventTypeChatUpdateProcessed,
				Timestamp: time.Now(),
			},
		)
	}

	mutexIface, _ := b.chatMutexes.LoadOrStore(chatID, &sync.Mutex{})
	mutex := mutexIface.(*sync.Mutex)
	locked := mutex.TryLock()
	if !locked {
		if update.CallbackQuery != nil {
			b.helper.AnswerCallbackQueryNil(update.CallbackQuery.ID)
		}

		if update.Message != nil {
			b.handleError(
				chatID,
				errors.ErrAnotherRequestInProgress,
				messages,
			)
		}

		return
	}
	defer mutex.Unlock()

	newCtx := types.NewContext(ctx, chat, messages)

	// update latest active at
	if chat != nil && (chat.LatestActiveAt.IsZero() ||
		chat.LatestActiveAt.Before(time.Now().Add(time.Minute*-10))) { // log maximum once per 10 minutes
		defer b.services.Chat.UpdateLatestActiveAt(newCtx)
	}

	// handle changed username
	if chat != nil {
		if update.Message != nil && update.Message.From.UserName != chat.Username {
			defer b.handleChangedUsername(newCtx, update.Message.From.UserName)
		} else if update.CallbackQuery != nil && update.CallbackQuery.From.UserName != chat.Username {
			defer b.handleChangedUsername(newCtx, update.CallbackQuery.From.UserName)
		}
	}

	// handle interactions
	if chat != nil { //nolint:nestif
		interaction, err := b.services.Interactions.Find(newCtx)
		if err != nil {
			b.handleError(chatID, err, messages)

			return
		}
		if interaction != nil {
			if update.Message != nil && !update.Message.IsCommand() {
				if err := b.handleInteraction(newCtx, interaction, update.Message); err != nil {
					b.handleError(chatID, err, messages)

					return
				}

				if err := b.services.Interactions.RemoveIfEquals(newCtx, interaction); err != nil {
					b.handleError(chatID, err, messages)
				}

				return
			}
		}
	}

	if update.CallbackQuery != nil {
		if err := b.handleCallbackQuery(newCtx, update.CallbackQuery); err != nil {
			b.handleCallbackQueryError(update.CallbackQuery.ID, chatID, err, messages)
		}

		return
	}

	// Handle commands
	if update.Message.IsCommand() {
		if err := b.handleCommand(newCtx, update.Message); err != nil {
			b.handleError(chatID, err, messages)
		}

		return
	}

	if err := b.handleMessage(newCtx); err != nil {
		b.handleError(chatID, err, messages)
	}
}

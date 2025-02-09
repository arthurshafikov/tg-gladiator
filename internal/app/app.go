package app

import (
	"context"
	"flag"
	"os"
	"time"

	configPkg "github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/events"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/arthurshafikov/tg-gladiator/internal/repository/pgsql"
	"github.com/arthurshafikov/tg-gladiator/internal/scheduler"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers/commands"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers/interactions"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers/notifications"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers/queries"
	"github.com/evalphobia/logrus_sentry"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	envFolderPath    string
	configFolderPath string
)

func init() {
	flag.StringVar(&envFolderPath, "env", "", "Path to .env file folder")
	flag.StringVar(&configFolderPath, "cfgFolder", "", "Path to configs folder")
}

func Run() {
	flag.Parse()

	time.Local = time.UTC

	ctx := context.Background()
	config := configPkg.NewConfig(envFolderPath, configFolderPath)
	logger := logrus.New()
	logrus.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.Info("Starting the app...")

	if config.Sentry.DSN != "" {
		sentryHook, err := logrus_sentry.NewSentryHook(config.Sentry.DSN, []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		})
		if err != nil {
			logger.Panic(err)
		}
		sentryHook.Timeout = 10 * time.Second
		sentryHook.StacktraceConfiguration.Enable = true
		sentryHook.SetEnvironment(config.Env)
		logger.Hooks.Add(sentryHook)
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Panic(err)
		}
	}()

	g := &errgroup.Group{}
	eventsHandler := events.NewHandler(ctx, g, logger)

	db := pgsql.ConnectToDatabase(ctx, &config.DBConfig, config.App.Debug)
	repository := repository.NewRepository(db)

	services := services.NewServices(services.Deps{
		Repository:    repository,
		Logger:        logger,
		Config:        config,
		EventsHandler: eventsHandler,
	})

	botAPI, err := tgbotapi.NewBotAPI(config.TelegramBotConfig.APIKey)
	if err != nil {
		logger.Error(err)

		return
	}
	telegramHelper := handlers.NewHelper(logger, botAPI, config)

	telegramHandlerParams := handlers.HandlerParams{
		Services: services,
		Logger:   logger,
		Config:   config,

		Helper: telegramHelper,
	}

	baseHandler := handlers.NewBaseHandler(telegramHandlerParams)

	telegramBot := telegram.NewBot(&telegram.Deps{
		Services: services,
		Logger:   logger,
		Config:   config,

		CommandsHandler:    commands.NewHandler(baseHandler),
		QueryHandler:       queries.NewHandler(baseHandler),
		InteractionHandler: interactions.NewHandler(baseHandler),

		Helper: telegramHelper,
	})

	notificationsHandler := notifications.NewHandler(baseHandler)
	eventsHandler.InitEvents(&events.Deps{
		NotificationsHandler: notificationsHandler,
		Services:             services,
		Config:               config,
	})

	scheduler, err := scheduler.NewScheduler(ctx, scheduler.Deps{
		Logger:               logger,
		Services:             services,
		Config:               config,
		NotificationsHandler: notificationsHandler,
	})
	if err != nil {
		logger.Error(err)

		return
	}
	g.Go(func() error {
		defer func() {
			if err := recover(); err != nil {
				logger.Panic(err)
			}
		}()

		return scheduler.Start()
	})

	g.Go(func() error {
		defer func() {
			if err := recover(); err != nil {
				logger.Panic(err)
			}
		}()

		return telegramBot.Start(ctx)
	})

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}
}

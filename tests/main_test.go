package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/arthurshafikov/tg-gladiator/internal/config"
	"github.com/arthurshafikov/tg-gladiator/internal/repository"
	"github.com/arthurshafikov/tg-gladiator/internal/repository/pgsql"
	"github.com/arthurshafikov/tg-gladiator/internal/services"
	mock_services "github.com/arthurshafikov/tg-gladiator/internal/services/mocks"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers"
	mock_handlers "github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers/mocks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

var TablesList = []string{
	"service_day_times",
	"service_days",
	"appointments",
	"services",
	"client_masters",
	"chats",
}

var r *require.Assertions

type APITestSuite struct {
	suite.Suite

	logger        *mock_services.MockLogger
	config        *config.Config
	eventsHandler *mock_services.MockEventsHandler
	services      *services.Services
	repos         *repository.Repository
	db            *gorm.DB

	bot            *telegram.Bot
	telegramHelper *mock_handlers.MockTelegramHandlerHelper

	ctx       context.Context
	ctxCancel context.CancelFunc
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {
	r = s.Require()
	s.ctx, s.ctxCancel = context.WithCancel(context.Background())

	ctrl := gomock.NewController(s.T())
	s.logger = mock_services.NewMockLogger(ctrl)
	s.eventsHandler = mock_services.NewMockEventsHandler(ctrl)
	s.config = config.NewConfig("config", "../configs")

	s.db = pgsql.ConnectToDatabase(s.ctx, &s.config.DBConfig, false)

	s.repos = repository.NewRepository(s.db)
	s.services = services.NewServices(services.Deps{
		Repository:    s.repos,
		Logger:        s.logger,
		Config:        s.config,
		EventsHandler: s.eventsHandler,
	})

	s.telegramHelper = mock_handlers.NewMockTelegramHandlerHelper(ctrl)

	_ = handlers.NewBaseHandler(handlers.HandlerParams{
		Services: s.services,
		Logger:   s.logger,
		Config:   s.config,

		Helper: s.telegramHelper,
	})

	s.bot = telegram.NewBot(&telegram.Deps{
		Services: s.services,
		Logger:   s.logger,
		Config:   s.config,

		// CommandsHandler:    commands.NewHandler(baseHandler),
		// QueryHandler:       queries.NewQueryHandler(baseHandler),
		// InteractionHandler: interactions.NewHandler(baseHandler),

		Helper: s.telegramHelper,
	})
}

func (s *APITestSuite) SetupTest() {
	// existed, err := s.services.Users.CreateIfNotExists(s.ctx, chatID, user.Username)
	// r.NoError(err)
	// r.False(existed)
}

func (s *APITestSuite) TearDownTest() {
	r.NoError(s.resetDatabase())
}

func (s *APITestSuite) TearDownSuite() {
	s.ctxCancel()
}

func (s *APITestSuite) resetDatabase() error {
	for _, tableName := range TablesList {
		if err := s.db.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", tableName)).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *APITestSuite) MockMenuCall(chatID int64) {
	gomock.InOrder(
		s.telegramHelper.EXPECT().NewMessage(chatID, gomock.Any()).Times(1).Return(tgbotapi.MessageConfig{}),
		s.telegramHelper.EXPECT().NewRow(gomock.Any(), gomock.Any()).Times(4).Return([]tgbotapi.InlineKeyboardButton{}),
		s.telegramHelper.EXPECT().Send(gomock.Any()).Times(1).Return(nil),
	)
}

package telegram

// import (
// 	"context"
// 	"testing"

// 	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
// 	"github.com/arthurshafikov/tg-gladiator/internal/services"
// 	mock_services "github.com/arthurshafikov/tg-gladiator/internal/services/mocks"
// 	mock_handlers "github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/handlers/mocks"
// 	mock_telegram "github.com/arthurshafikov/tg-gladiator/internal/transport/telegram/mocks"
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// 	"golang.org/x/sync/errgroup"
// )

// var (
// 	updateConfig = tgbotapi.UpdateConfig{
// 		Offset:  0,
// 		Timeout: 60,
// 	}
// 	updateChannel = make(chan tgbotapi.Update)
// 	msg           = &tgbotapi.Message{
// 		Text: "some msg",
// 		Chat: &tgbotapi.Chat{
// 			ID:   chatID,
// 			Type: telegram.TelegramChatTypePrivate,
// 		},
// 	}
// )

// type MockBag struct {
// 	Chats  *mock_services.MockChats
// 	Logger *mock_services.MockLogger

// 	Helper *mock_handlers.MockTelegramHandlerHelper

// 	CommandsHandler *mock_telegram.MockCommandsHandler
// 	QueryHandler    *mock_telegram.MockQueryHandler
// }

// func getBotWithMockBag(t *testing.T) (*Bot, *MockBag) {
// 	t.Helper()
// 	ctrl := gomock.NewController(t)
// 	chatsMock := mock_services.NewMockChats(ctrl)
// 	helperMock := mock_handlers.NewMockTelegramHandlerHelper(ctrl)
// 	loggerMock := mock_services.NewMockLogger(ctrl)
// 	services := &services.Services{
// 		Chats: chatsMock,
// 	}
// 	commandHandler := mock_telegram.NewMockCommandsHandler(ctrl)
// 	queryHandler := mock_telegram.NewMockQueryHandler(ctrl)

// 	return &Bot{
// 			helper:   helperMock,
// 			logger:   loggerMock,
// 			services: services,

// 			commandHandler: commandHandler,
// 			queryHandler:   queryHandler,
// 		}, &MockBag{
// 			Chats:  chatsMock,
// 			Logger: loggerMock,
// 			Helper: helperMock,

// 			CommandsHandler: commandHandler,
// 			QueryHandler:    queryHandler,
// 		}
// }

// func TestStart(t *testing.T) {
// 	g := errgroup.Group{}
// 	bot, mockBag := getBotWithMockBag(t)

// 	gomock.InOrder(
// 		mockBag.Helper.EXPECT().NewUpdateChannel(0).Return(updateConfig),
// 		mockBag.Helper.EXPECT().GetUpdatesChan(updateConfig).Return(updateChannel, nil),
// 	)
// 	g.Go(func() error {
// 		defer close(updateChannel)

// 		// t.Run("start", func(t *testing.T) {
// 		// 	gomock.InOrder(
// 		// 		mockBag.Chat.EXPECT().GetChat(bot.ctx, chatID).Return(expectedUser, nil),
// 		// 		mockBag.Chat.EXPECT().SetChatType(bot.ctx, chatID, telegram.TelegramChatTypePrivate).Return(nil),
// 		// 		mockBag.Interactions.EXPECT().Get(bot.ctx, chatID).Return(0, nil),
// 		// 		mockBag.CommandsHandler.EXPECT().HandleSetLastNote(language, msg).Return(nil),
// 		// 	)
// 		// 	msg := msg
// 		// 	msg.Chat.Type = telegram.TelegramChatTypePrivate
// 		// 	updateChannel <- tgbotapi.Update{Message: msg}
// 		// })

// 		return nil
// 	})

// 	err := bot.Start(context.Background())
// 	require.NoError(t, err)

// 	require.NoError(t, g.Wait())
// }

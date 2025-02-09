package queries

import (
	"fmt"
	"slices"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenShop(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	shopItems, err := h.Services.Shop.GetShopItemsFor(ctx, heroID)
	if err != nil {
		return err
	}

	slices.SortFunc(shopItems, func(a, b models.HeroShopItem) int {
		if a.Price < b.Price {
			return -1
		} else if a.Price > b.Price {
			return 1
		}

		return 0
	})

	shopItemsText := ctx.Messages().ShopIntro

	for _, shopItem := range shopItems {
		shopItemsText += fmt.Sprintf(
			"\n\n- %s (%s)💰%v",
			shopItem.Item.Name,
			shopItem.Item.GetShortCharacteristicsText(),
			shopItem.Price,
		)
	}

	// @todo updates at in text?
	shopItemsText += fmt.Sprintf("\n\n%s - 💰%v", ctx.Messages().ShopYourBalance, hero.CurrentGold)

	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, shopItemsText)

	shopItemsButtons := make([]telegram.KeyboardButton, 0, len(shopItems))
	for _, shopItem := range shopItems {
		shopItemsButtons = append(shopItemsButtons, telegram.KeyboardButton{
			CallbackQuery: queries.OpenMenu, // @todo ShopBuyItem
			Text:          shopItem.Item.Name,
		})
	}

	msg.ReplyMarkup = h.Helper.CreateKeyboard(shopItemsButtons, 2)
	msg.ReplyMarkup.InlineKeyboard = append(
		msg.ReplyMarkup.InlineKeyboard,
		h.GetBackButtonKeyboardRow(
			ctx,
			queries.OpenMyHero.WithID(heroID),
		),
	)

	return h.Helper.Send(msg)
}

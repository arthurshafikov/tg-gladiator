package queries

import (
	"fmt"
	"slices"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
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

	return h.openShop(ctx, heroID, query)
}

func (h *Handler) HandleShopBuyItem(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 2); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	itemID, err := h.GetIDFromString(payload[1])
	if err != nil {
		return err
	}

	shopItem, err := h.Services.HeroShopItem.FindMy(ctx, heroID, itemID)
	if err != nil {
		return err
	}

	msgText := fmt.Sprintf(
		"%s\n\n%s",
		fmt.Sprintf(
			ctx.Messages().ShopBuyItemConfirmation,
			shopItem.Item.Name,
			shopItem.Price,
		),
		fmt.Sprintf(
			"%s - 💰%v",
			ctx.Messages().ShopYourBalance,
			hero.CurrentGold,
		),
	)
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, msgText)

	msg.ReplyMarkup = h.Helper.CreateKeyboard([]telegram.KeyboardButton{
		{
			CallbackQuery: queries.ShopBuyItemConfirm.WithID(heroID).WithID(itemID),
			Text:          ctx.Messages().ShopBuyItemConfirm,
		},
		{
			CallbackQuery: queries.OpenShop.WithID(heroID),
			Text:          ctx.Messages().ShopBackButton,
		},
	}, 1)

	return h.Helper.Send(msg)
}

func (h *Handler) HandleShopBuyItemConfirm(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 2); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	itemID, err := h.GetIDFromString(payload[1])
	if err != nil {
		return err
	}

	shopItem, err := h.Services.HeroShopItem.BuyItem(ctx, heroID, itemID)
	if err != nil {
		return err
	}

	msgText := fmt.Sprintf(
		ctx.Messages().ShopBuyItemSuccess,
		shopItem.Item.Name,
	)
	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, msgText)
	if err := h.Helper.Send(msg); err != nil {
		return err
	}

	return h.openShop(ctx, heroID)
}

func (h *Handler) openShop(ctx *types.Context, heroID int64, query ...*tgbotapi.CallbackQuery) error {
	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	shopItems, err := h.Services.Shop.GetShopItemsFor(ctx, heroID)
	if err != nil {
		return err
	}

	shopItemsByCategories := make(map[enums.ItemCategory][]models.HeroShopItem, len(enums.ItemCategoryAll()))

	for _, shopItem := range shopItems {
		shopItemsByCategories[shopItem.Item.Category] = append(
			shopItemsByCategories[shopItem.Item.Category],
			shopItem,
		)
	}

	// for category, shopItemsByCategoryElement := range shopItemsByCategories {
	// 	slices.SortFunc(shopItemsByCategoryElement[category], func(a, b models.HeroShopItem) int {
	// 		if a.Price < b.Price {
	// 			return -1
	// 		} else if a.Price > b.Price {
	// 			return 1
	// 		}

	// 		return 0
	// 	})
	// }

	shopItemsText := ctx.Messages().ShopIntro

	for _, itemCategory := range enums.ItemCategoryAll() {
		if len(shopItemsByCategories[itemCategory]) < 1 {
			continue
		}

		shopItemsText += fmt.Sprintf(
			"\n\n%s:",
			ctx.Messages().ItemCategories[itemCategory],
		)

		shopItems := shopItemsByCategories[itemCategory]
		slices.SortFunc(shopItems, func(a, b models.HeroShopItem) int {
			if a.Price < b.Price {
				return -1
			} else if a.Price > b.Price {
				return 1
			}

			return 0
		})

		for _, shopItem := range shopItems {
			shopItemsText += fmt.Sprintf(
				"\n- %s (%s)💰%v",
				shopItem.Item.Name,
				shopItem.Item.GetShortCharacteristicsText(),
				shopItem.Price,
			)
		}
	}

	// @todo updates at in text?
	shopItemsText += fmt.Sprintf("\n\n%s - 💰%v", ctx.Messages().ShopYourBalance, hero.CurrentGold)

	shopItemsButtons := make([]telegram.KeyboardButton, 0, len(shopItems)+1)
	for _, shopItem := range shopItems {
		shopItemsButtons = append(shopItemsButtons, telegram.KeyboardButton{
			CallbackQuery: queries.ShopBuyItem.WithID(heroID).WithID(shopItem.ItemID),
			Text:          shopItem.Item.Name,
		})
	}
	shopItemsButtons = append(shopItemsButtons, telegram.KeyboardButton{
		CallbackQuery: queries.OpenMyHero.WithID(heroID),
		Text:          ctx.Messages().DefaultBackBtn,
	})

	keyboard := h.Helper.CreateKeyboard(shopItemsButtons, 2)

	if len(query) > 0 {
		msg := h.Helper.NewEditMessage(ctx.GetChatID(), query[0].Message.MessageID, shopItemsText)
		msg.ReplyMarkup = keyboard

		return h.Helper.Send(msg)
	} else {
		msg := h.Helper.NewMessage(ctx.GetChatID(), shopItemsText)
		msg.ReplyMarkup = keyboard

		return h.Helper.Send(msg)
	}
}

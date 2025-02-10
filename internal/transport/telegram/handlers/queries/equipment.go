package queries

import (
	"fmt"

	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/queries"
	"github.com/arthurshafikov/tg-gladiator/internal/core/constants/telegram"
	"github.com/arthurshafikov/tg-gladiator/internal/core/helpers"
	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *Handler) HandleOpenHeroEquipment(ctx *types.Context, query *tgbotapi.CallbackQuery, payload []string) error {
	if err := h.ValidatePayloadLength(payload, 1); err != nil {
		return err
	}

	heroID, err := h.GetIDFromString(payload[0])
	if err != nil {
		return err
	}

	page := 1
	if len(payload) > 1 {
		pageInt64, err := h.GetIDFromString(payload[1])
		if err != nil {
			return err
		}

		page = int(pageInt64)
	}

	hero, err := h.Services.Hero.FindMy(ctx, heroID)
	if err != nil {
		return err
	}

	heroItemsPaginated, err := h.Services.HeroItem.GetPaginatedForHero(ctx, hero.ID, page)
	if err != nil {
		return err
	}

	// heroEquipment, err := getHeroEquipment
	heroEquipment := map[enums.ItemEquipsOn]models.Item{
		enums.ItemEquipsOnBody: {
			Name: "Шлем!",
		},
		enums.ItemEquipsOnFinger: {
			Name: "Кольцо",
		},
	}

	var equipmentText string
	for _, itemEquipsOn := range enums.ItemEquipsOnAll() {
		itemEquippedText := "-"

		if _, ok := heroEquipment[itemEquipsOn]; ok {
			itemEquippedText = heroEquipment[itemEquipsOn].Name
		}

		equipmentText += fmt.Sprintf(
			"\n%s: %s",
			ctx.Messages().ItemEquipsOn[itemEquipsOn],
			itemEquippedText,
		)
	}

	msgText := fmt.Sprintf(
		"%s:\n%s\n\n%s:",
		ctx.Messages().HeroEquipmentOverview,
		equipmentText,
		ctx.Messages().HeroEquipmentChooseItem,
	)

	msg := h.Helper.NewEditMessage(ctx.GetChatID(), query.Message.MessageID, msgText)

	itemsKeyboardButtons := []telegram.KeyboardButton{}

	for _, heroItem := range heroItemsPaginated.Rows {
		itemsKeyboardButtons = append(
			itemsKeyboardButtons,
			telegram.KeyboardButton{
				CallbackQuery: queries.OpenMyHero.WithID(hero.ID),
				Text:          ctx.Messages().GetItemNameWithIcon(heroItem.Item),
			},
		)
	}

	keyboard := h.Helper.CreateKeyboard(itemsKeyboardButtons, 2)

	if heroItemsPaginated.HasNextPage() || heroItemsPaginated.HasPreviousPage() {
		paginationKeyboardRow := []tgbotapi.InlineKeyboardButton{}

		if heroItemsPaginated.HasPreviousPage() {
			paginationKeyboardRow = append(
				paginationKeyboardRow,
				tgbotapi.InlineKeyboardButton{
					Text: ctx.Messages().PreviousPage,
					CallbackData: helpers.GetPointer(
						string(queries.OpenHeroEquipment.WithID(hero.ID).WithID(int64(page) - 1)),
					),
				},
			)
		}
		if heroItemsPaginated.HasNextPage() {
			paginationKeyboardRow = append(
				paginationKeyboardRow,
				tgbotapi.InlineKeyboardButton{
					Text: ctx.Messages().NextPage,
					CallbackData: helpers.GetPointer(
						string(queries.OpenHeroEquipment.WithID(hero.ID).WithID(int64(page) + 1)),
					),
				},
			)
		}

		keyboard.InlineKeyboard = append(
			keyboard.InlineKeyboard,
			paginationKeyboardRow,
		)
	}

	keyboard.InlineKeyboard = append(
		keyboard.InlineKeyboard,
		h.GetBackButtonKeyboardRow(ctx, queries.OpenMyHero.WithID(hero.ID)),
	)
	msg.ReplyMarkup = keyboard

	return h.Helper.Send(msg)
}

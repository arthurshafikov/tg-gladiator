package errors

import (
	"errors"
	"fmt"
)

var (
	ErrForbidden       = errors.New("403_forbidden")
	ErrNotFound        = errors.New("404_not_found")
	ErrServerError     = errors.New("500_server_error")
	ErrTooManyRequests = errors.New("429_too_many_requests")

	ErrAlreadyExists          = errors.New("already_exists")
	ErrUndefinedCommand       = errors.New("undefined_command")
	ErrUndefinedInteraction   = errors.New("undefined_interaction")
	ErrUndefinedCallbackQuery = errors.New("undefined_callbackQuery")

	ErrInvalidHeroClass            = errors.New("invalid_hero_class")
	ErrHasActiveFight              = errors.New("has_active_fight")
	ErrShopNotExists               = errors.New("shop_not_exists")
	ErrHeroShopItemNotExists       = errors.New("hero_shop_item_not_exists")
	ErrInsufficientGold            = errors.New("insufficient_gold")
	ErrInsufficientEnergy          = errors.New("insufficient_energy")
	ErrAlreadyHaveEquippedItem     = errors.New("already_have_equipped_item")
	ErrTooMuchItemsEquippedAlready = errors.New("too_much_items_equipped_already")
	ErrCannotConsumeEmptyQuantity  = errors.New("cannot_consume_empty_quantity")
	ErrEmpty                       = errors.New("empty")
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func ErrTooLong(maximumSymbols int) error {
	return fmt.Errorf("too_long:%v", maximumSymbols)
}

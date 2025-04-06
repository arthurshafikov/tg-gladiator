package models

import (
	"fmt"
	"math"
	"time"
)

type HeroShop struct {
	ID        int64     `json:"id"`
	HeroID    int64     `json:"hero_id"`
	UpdatesAt time.Time `json:"updates_at"`
	CreatedAt time.Time `json:"created_at"`
}

type HeroShopItem struct {
	Item       Item  `gorm:"->"`
	ItemID     int64 `json:"item_id"`
	HeroShopID int64 `json:"hero_shop_id"`
	Price      int   `json:"price"`
}

func (hs *HeroShop) GetUpdatesInText() string {
	updatesIn := time.Until(hs.UpdatesAt)

	updatesInText := ""

	hours := math.Floor(updatesIn.Hours())
	if hours > 0 {
		updatesInText += fmt.Sprintf("%v час(а) ", hours) // @todo config messages
	} else {
		hours = 0
	}

	minutes := math.Floor(updatesIn.Minutes() - hours*60)
	if minutes < 1 {
		minutes = 1
	}
	updatesInText += fmt.Sprintf("%v минут", minutes)

	return updatesInText
}

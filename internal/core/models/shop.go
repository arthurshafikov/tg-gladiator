package models

import (
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

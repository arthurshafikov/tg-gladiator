package models

type HeroItem struct {
	HeroID     int64 `json:"hero_id"`
	ItemID     int64 `json:"item_id"`
	IsEquipped bool  `json:"is_equipped"`
}

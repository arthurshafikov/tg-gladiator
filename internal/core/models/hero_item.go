package models

import "github.com/arthurshafikov/tg-gladiator/internal/core/constants/enums"

const HeroItemFieldQuantity = "quantity"

type HeroItem struct {
	Item       Item  `gorm:"->"`
	HeroID     int64 `json:"hero_id"`
	ItemID     int64 `json:"item_id"`
	IsEquipped bool  `json:"is_equipped"`
	Quantity   int   `json:"quantity"`
}

type HeroEquipment map[enums.ItemEquipsOn][]Item

func (he HeroEquipment) GetAllItems() []Item {
	result := make([]Item, 0, len(he))
	for _, items := range he {
		result = append(result, items...)
	}

	return result
}

type PaginatedHeroItem struct {
	Rows       []HeroItem
	Page       int
	PerPage    int
	TotalCount int64
}

func (p *PaginatedHeroItem) HasNextPage() bool {
	totalPages := p.TotalCount / int64(p.PerPage)
	if p.TotalCount%int64(p.PerPage) != 0 {
		totalPages++
	}

	return int64(p.Page) < totalPages
}

func (p *PaginatedHeroItem) HasPreviousPage() bool {
	return p.Page > 1
}

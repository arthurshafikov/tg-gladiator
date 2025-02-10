package models

type HeroItem struct {
	Item       Item  `gorm:"->"`
	HeroID     int64 `json:"hero_id"`
	ItemID     int64 `json:"item_id"`
	IsEquipped bool  `json:"is_equipped"`
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

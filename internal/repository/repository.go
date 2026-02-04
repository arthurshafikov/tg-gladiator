package repository

import (
	"context"

	"github.com/arthurshafikov/tg-gladiator/internal/core/models"
	"github.com/arthurshafikov/tg-gladiator/internal/core/types"
	"github.com/arthurshafikov/tg-gladiator/internal/repository/pgsql"
	"gorm.io/gorm"
)

type BaseRepo interface {
	WrapInTransaction(ctx *types.Context, txFunc types.TXFunction) error
}

type Chat interface {
	Find(ctx context.Context, id int64) (*models.Chat, error)
	FindBy(ctx context.Context, fields *models.Chat) (*models.Chat, error)
	Create(ctx context.Context, chat models.Chat) (*models.Chat, error)
	Update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error)
	UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Chat, error)
}

type Hero interface {
	Create(ctx context.Context, chat models.Hero) (*models.Hero, error)
	DeleteByID(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*models.Hero, error)
	// FindBy(ctx context.Context, fields *models.Chat) (*models.Chat, error)
	GetBy(ctx context.Context, fields *models.Hero) (*[]models.Hero, error)
	// Update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error)
	UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Hero, error)
	ReplenishHeroesEnergy(ctx context.Context) error
}

type Fight interface {
	Create(ctx context.Context, fight models.Fight) (*models.Fight, error)
	FindBy(ctx context.Context, fields *models.Fight) (*models.Fight, error)
	FindByMap(ctx context.Context, fields map[string]interface{}) (*models.Fight, error)
	Update(ctx context.Context, id int64, fields *models.Fight) error
	UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) error
}

type Enemy interface {
	Find(ctx context.Context, id int64) (*models.Enemy, error)
	FindBy(ctx context.Context, fields *models.Enemy) (*models.Enemy, error)
	GetByMap(ctx context.Context, where *types.WhereConditions) (*[]models.Enemy, error)
	GetNextBossForHero(ctx context.Context, heroID int64) (*models.Enemy, error)
	GetBossesAmountLeft(ctx context.Context, heroID int64) (int, error)
	// Update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error)
	// UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Chat, error)
}

type Item interface {
	GetBy(ctx context.Context, where *types.WhereConditions) ([]models.Item, error)
}

type HeroItem interface {
	Create(ctx context.Context, heroItem models.HeroItem) (*models.HeroItem, error)
	DeleteBy(ctx context.Context, fields *models.HeroItem) error
	FindBy(ctx context.Context, fields *models.HeroItem) (*models.HeroItem, error)
	GetBy(
		ctx context.Context,
		fields *models.HeroItem,
		pagination ...*types.Pagination,
	) (*models.PaginatedHeroItem, error)
	UpdateMap(ctx context.Context, heroItem *models.HeroItem, fields map[string]interface{}) (*models.HeroItem, error)
}

type HeroShop interface {
	Create(ctx context.Context, heroShop models.HeroShop) (*models.HeroShop, error)
	FindBy(ctx context.Context, heroShop *models.HeroShop) (*models.HeroShop, error)
	Update(ctx context.Context, id int64, fields *models.HeroShop) (*models.HeroShop, error)
}

type HeroShopItem interface {
	Create(ctx context.Context, heroShopItem models.HeroShopItem) (*models.HeroShopItem, error)
	DeleteBy(ctx context.Context, fields *models.HeroShopItem) error
	FindBy(ctx context.Context, fields *models.HeroShopItem) (*models.HeroShopItem, error)
	GetBy(ctx context.Context, fields *models.HeroShopItem) ([]models.HeroShopItem, error)
}

type AnalyticEvent interface {
	Create(ctx context.Context, event models.AnalyticEvent) error
}

type Repository struct {
	Chat
	Hero
	Fight
	Enemy
	Item
	HeroItem
	HeroShop
	HeroShopItem
	AnalyticEvent
}

func NewRepository(db *gorm.DB) *Repository {
	baseRepo := pgsql.NewBaseRepo(db)

	return &Repository{
		Chat:          pgsql.NewChatRepository(*baseRepo),
		Hero:          pgsql.NewHeroRepository(*baseRepo),
		Fight:         pgsql.NewFightRepository(*baseRepo),
		Enemy:         pgsql.NewEnemyRepository(*baseRepo),
		Item:          pgsql.NewItemRepository(*baseRepo),
		HeroItem:      pgsql.NewHeroItemRepository(*baseRepo),
		HeroShop:      pgsql.NewHeroShopRepository(*baseRepo),
		HeroShopItem:  pgsql.NewHeroShopItemRepository(*baseRepo),
		AnalyticEvent: pgsql.NewAnalyticEventRepository(*baseRepo),
	}
}

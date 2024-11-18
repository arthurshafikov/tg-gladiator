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
	// UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Chat, error)
}

type Fight interface {
	Create(ctx context.Context, fight models.Fight) (*models.Fight, error)
	Find(ctx context.Context, id int64) (*models.Fight, error)
	FindBy(ctx context.Context, fields *models.Fight) (*models.Fight, error)
	Update(ctx context.Context, id int64, fields *models.Fight) (*models.Fight, error)
	UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Fight, error)
}

type Enemy interface {
	Find(ctx context.Context, id int64) (*models.Enemy, error)
	FindBy(ctx context.Context, fields *models.Enemy) (*models.Enemy, error)
	GetBy(ctx context.Context, fields *models.Enemy) (*[]models.Enemy, error)
	// Update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error)
	// UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Chat, error)
}

type Repository struct {
	Chat
	Hero
	Fight
	Enemy
}

func NewRepository(db *gorm.DB) *Repository {
	baseRepo := pgsql.NewBaseRepo(db)

	return &Repository{
		Chat:  pgsql.NewChatRepository(*baseRepo),
		Hero:  pgsql.NewHeroRepository(*baseRepo),
		Fight: pgsql.NewFightRepository(*baseRepo),
		Enemy: pgsql.NewEnemyRepository(*baseRepo),
	}
}

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
	// Find(ctx context.Context, id int64) (*models.Chat, error)
	// FindBy(ctx context.Context, fields *models.Chat) (*models.Chat, error)
	GetBy(ctx context.Context, fields *models.Hero) (*[]models.Hero, error)
	// Create(ctx context.Context, chat models.Chat) (*models.Chat, error)
	// Update(ctx context.Context, id int64, fields *models.Chat) (*models.Chat, error)
	// UpdateMap(ctx context.Context, id int64, fields map[string]interface{}) (*models.Chat, error)
}

type Repository struct {
	Chat
	Hero
}

func NewRepository(db *gorm.DB) *Repository {
	baseRepo := pgsql.NewBaseRepo(db)

	return &Repository{
		Chat: pgsql.NewChatRepository(*baseRepo),
		Hero: pgsql.NewHeroRepository(*baseRepo),
	}
}

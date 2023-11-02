package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, category *entity.Category) (*entity.Category, error)
	FindAll(ctx context.Context) ([]*entity.Category, error)
	FindByID(ctx context.Context, categoryID string) (*entity.Category, error)
	FindByName(ctx context.Context, name string) ([]*entity.Category, error)
	FindCompleteByID(ctx context.Context, categoryID string) (*entity.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category *entity.Category) (*entity.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, categoryID string) error
}

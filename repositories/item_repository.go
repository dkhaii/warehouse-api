package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type ItemRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, item *entity.Item) (*entity.Item, error)
	FindAll(ctx context.Context) ([]*entity.Item, error)
	FindByID(ctx context.Context, itmID uuid.UUID) (*entity.Item, error)
	FindByName(ctx context.Context, name string) ([]*entity.Item, error)
	FindCompleteByID(ctx context.Context, itmID uuid.UUID) (*entity.Item, error)
	FindByCategoryName(ctx context.Context, ctgName string) ([]*entity.Item, error)
	Update(ctx context.Context, tx *sql.Tx, item *entity.Item) error
	Delete(ctx context.Context, tx *sql.Tx, itmID uuid.UUID) error
}

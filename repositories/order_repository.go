package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
)

type OrderRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, order *entity.Order) (*entity.Order, error)
	FindAll(ctx context.Context) ([]*entity.Order, error)
	FindCompleteByID(ctx context.Context, orderID string) (*entity.Order, error)
	Delete(ctx context.Context, tx *sql.Tx, orderID string)
}

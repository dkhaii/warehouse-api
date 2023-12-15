package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type OrderRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, order *entity.Order) (*entity.Order, error)
	FindAll(ctx context.Context) ([]*entity.Order, error)
	FindCompleteByID(ctx context.Context, orderID uuid.UUID) (*entity.Order, error)
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*entity.Order, error)
}

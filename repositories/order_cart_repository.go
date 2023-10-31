package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
)

type OrderCartRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, orderCart *entity.OrderCart) (*entity.OrderCart, error)
}
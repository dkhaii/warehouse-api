package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
)

type orderCartRepositoryImpl struct {
	database *sql.DB
}

func NewOrderCartRepository(database *sql.DB) OrderCartRepository {
	return &orderCartRepositoryImpl{
		database: database,
	}
}

func (repository *orderCartRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, ordCart *entity.OrderCart) (*entity.OrderCart, error) {
	query := `
		INSERT INTO order_carts
		(id, order_id, item_id, quantity)
		VALUES
		(?, ?, ?, ?)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		ordCart.ID,
		ordCart.OrderID,
		ordCart.ItemID,
		ordCart.Quantity,
	)
	if err != nil {
		return nil, err
	}

	return ordCart, nil
}

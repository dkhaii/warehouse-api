package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/google/uuid"
)

type transferOrderRepositoryImpl struct {
	database *sql.DB
}

func NewTransferOrderRepository(database *sql.DB) TransferOrderRepository {
	return &transferOrderRepositoryImpl{
		database: database,
	}
}

func (repository *transferOrderRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, trfOrd *entity.TransferOrder) (*entity.TransferOrder, error) {
	query := `
		INSERT INTO transfer_orders
		(id, order_id, user_id, status, fulfilled_date, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		trfOrd.ID,
		trfOrd.OrderID,
		trfOrd.UserID,
		trfOrd.Status,
		trfOrd.FulfilledDate,
		trfOrd.CreatedAt,
		trfOrd.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return trfOrd, nil
}

func (repository *transferOrderRepositoryImpl) FindAll(ctx context.Context) ([]*entity.TransferOrder, error) {
	query := `SELECT * FROM transfer_orders`

	rows, err := repository.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var listOfTransferOrders []*entity.TransferOrder

	for rows.Next() {
		var to entity.TransferOrder

		err := rows.Scan(
			&to.ID,
			&to.OrderID,
			&to.UserID,
			&to.Status,
			&to.FulfilledDate,
			&to.CreatedAt,
			&to.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfTransferOrders = append(listOfTransferOrders, &to)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfTransferOrders, nil
}

func (repository *transferOrderRepositoryImpl) FindByID(ctx context.Context, trfOrdID uuid.UUID) (*entity.TransferOrder, error) {
	var to entity.TransferOrder

	query := "SELECT * FROM transfer_orders WHERE id = ?"

	err := repository.database.QueryRowContext(ctx, query, trfOrdID).Scan(
		&to.ID,
		&to.OrderID,
		&to.UserID,
		&to.Status,
		&to.FulfilledDate,
		&to.CreatedAt,
		&to.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrTransferOrderNotFound
		}
		return nil, err
	}

	return &to, nil
}

func (repository *transferOrderRepositoryImpl) FindCompleteByID(ctx context.Context, trfOrdID uuid.UUID) (*entity.TransferOrder, error) {
	var to entity.TransferOrder
	var order entity.Order

	query := `
		SELECT to.*, o.*
		FROM transfer_orders to
		LEFT JOIN orders o ON o.id = to.order_id
		WHERE to.id = ?
	`

	err := repository.database.QueryRowContext(ctx, query, trfOrdID).Scan(
		&to.ID,
		&to.OrderID,
		&to.UserID,
		&to.Status,
		&to.FulfilledDate,
		&to.CreatedAt,
		&to.UpdatedAt,
		&order.ID,
		&order.UserID,
		&order.Notes,
		&order.RequestTransferDate,
		&order.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrTransferOrderNotFound
		}
		return nil, err
	}

	var items []entity.ItemFiltered

	query2 := `
		SELECT i.id, i.name, i.description, i.availability, i.category_id
		FROM items i
		LEFT JOIN order_carts oc ON i.id = oc.item_id
		LEFT JOIN order o ON o.id = oc.order_id
		WHERE o.id = ?
	`

	rows, err := repository.database.QueryContext(ctx, query2, order.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.ItemFiltered

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Availability,
			&item.CategoryID,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	to.Order = &order
	to.Order.Items = items

	return &to, nil
}

func (repository *transferOrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, trfOrd *entity.TransferOrder) (*entity.TransferOrder, error) {
	query := `
		UPDATE transfer_orders
		SET user_id = ?, status = ?, filled_date = ?, update_at = ?
		WHERE id = ?
	`

	_, err := tx.ExecContext(
		ctx, 
		query,
		trfOrd.UserID,
		trfOrd.Status,
		trfOrd.FulfilledDate,
		trfOrd.UpdatedAt,
		trfOrd.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return trfOrd, nil
}
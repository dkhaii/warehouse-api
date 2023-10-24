package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/google/uuid"
)

type orderRepositoryImpl struct {
	database *sql.DB
}

func NewOrderRepository(database *sql.DB) OrderRepository {
	return &orderRepositoryImpl{
		database: database,
	}
}

func (repository *orderRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, ord *entity.Order) (*entity.Order, error) {
	query := `
	INSER INTO orders
	(id, item_id, user_id, quantity, notes, request_transfer_date, created_at)
	VALUES
	(?, ?, ?, ?, ?, ?, ?)
	`

	_, err := repository.database.ExecContext(
		ctx,
		query,
		ord.ID,
		ord.ItemID,
		ord.UserID,
		ord.Quantity,
		ord.Notes,
		ord.RequestTransferDate,
		ord.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return ord, nil
}

func (repository *orderRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Order, error) {
	query := "SELECT * FROM orders"

	rows, err := repository.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listOfOrders []*entity.Order

	for rows.Next() {
		var order entity.Order

		err := rows.Scan(
			&order.ID,
			&order.ItemID,
			&order.UserID,
			&order.Quantity,
			&order.Notes,
			&order.RequestTransferDate,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfOrders = append(listOfOrders, &order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfOrders, nil
}

func (repository *orderRepositoryImpl) FindCompleteByID(ctx context.Context, ordID uuid.UUID) (*entity.Order, error) {
	var order entity.Order
	var user entity.UserFiltered	
	
	query := `
	SELECT o.*, u.id, u.username, u.contact
	FROM orders o
	LEFT JOIN users u
	ON u.id = o.user_id
	WHERE o.id = ? 
	`
	
	err := repository.database.QueryRowContext(ctx, query, ordID).Scan(
		&order.ID,
		&order.ItemID,
		&order.UserID,
		&order.Quantity,
		&order.Notes,
		&order.RequestTransferDate,
		&order.CreatedAt,
		&user.ID,
		&user.Username,
		&user.Contact,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrOrderNotFound
		}
		return nil, err
	}

	order.User = &user

	var items []entity.ItemFiltered
	
	query2 := `
	SELECT o.id, o.item_id, i.id, i.name, i.description, i.availability, i.categoy_id
	FROM orders o
	LEFT JOIN items i
	ON i.id = o.item_id
	WHERE o.id = ?
	` 

	rows, err := repository.database.QueryContext(ctx, query2, ordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.ItemFiltered

		err := rows.Scan(
			&order.ID,
			&order.ItemID,
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

	order.Item = items

	return &order, nil
}
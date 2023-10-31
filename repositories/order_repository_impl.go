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
	INSERT INTO orders
	(id, user_id, notes, request_transfer_date, created_at)
	VALUES
	(?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		ord.ID,
		ord.UserID,
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
			&order.UserID,
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
	SELECT o.*, u.username, u.contact
	FROM orders o
	LEFT JOIN users u
	ON u.id = o.user_id
	WHERE o.id = ? 
	`

	err := repository.database.QueryRowContext(ctx, query, ordID).Scan(
		&order.ID,
		&order.UserID,
		&order.Notes,
		&order.RequestTransferDate,
		&order.CreatedAt,
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

	query2 := `
		SELECT i.id, i.name, i.description, i.availability, i.category_id
		FROM items i
		LEFT JOIN order_carts oc ON i.id = oc.item_id
		LEFT JOIN orders o ON o.id = oc.order_id
		WHERE o.id = ?
	`

	rows, err := repository.database.QueryContext(ctx, query2, ordID)
	if err != nil {
		return nil, err
	}

	var listOfItems []entity.ItemFiltered

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

		listOfItems = append(listOfItems, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	order.Items = listOfItems

	return &order, nil
}

package repositories

// import (
// 	"context"
// 	"database/sql"

// 	"github.com/dkhaii/warehouse-api/entity"
// )

// type orderRepositoryImpl struct {
// 	database *sql.DB
// }

// func NewOrderRepository(database *sql.DB) OrderRepository {
// 	return &orderRepositoryImpl{
// 		database: database,
// 	}
// }

// func (repository *orderRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, ord *entity.Order) (*entity.Order, error) {
// 	query := `
// 	INSER INTO orders
// 	(id, item_id, user_id, quantity, notes, request_transfer_date, created_at)
// 	VALUES
// 	(?, ?, ?, ?, ?, ?, ?)
// 	`

// 	_, err := repository.database.ExecContext(
// 		ctx,
// 		query,
// 		ord.ID,
// 		ord.ItemID,
// 		ord.UserID,
// 		ord.Quantity,
// 		ord.Notes,
// 		ord.RequestTransferDate,
// 		ord.CreatedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return ord, nil
// }

// func (repository *orderRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Order, error) {
// 	query := "SELECT * FROM orders"

// 	rows, err := repository.database.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var listOfOrders []*entity.Order

// 	for rows.Next() {
// 		var order entity.Order

// 		err := rows.Scan(
// 			&order.ID,
// 			&order.ItemID,
// 			&order.UserID,
// 			&order.Quantity,
// 			&order.Notes,
// 			&order.RequestTransferDate,
// 			&order.CreatedAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}

// 		listOfOrders = append(listOfOrders, &order)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return listOfOrders, nil
// }

// func (repository *orderRepositoryImpl) FindCompleteByID(ctx context.Context, ordID string) (*entity.Order, error) {
// 	var order entity.Order
// 	var items []entity.Item
	
// 	query := `
// 	SELECT o.*, i.id, i.name, i.availability, i.location_id, i.category_id
// 	FROM orders o
// 	LEFT JOIN items i
// 	ON o.id = i.id
// 	WHERE o.id = ?
// 	`
// }
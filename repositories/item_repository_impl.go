package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/google/uuid"
)

type itemRepositoryImpl struct {
	database *sql.DB
}

func NewItemRepository(database *sql.DB) ItemRepository {
	return &itemRepositoryImpl{
		database: database,
	}
}

func (repository *itemRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, itm *entity.Item) (*entity.Item, error) {
	query := `
		INSERT INTO items 
		(id, name, description, quantity, availability, category_id, user_id, created_at, updated_at) 
		VALUES 
		(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		itm.ID,
		itm.Name,
		itm.Description,
		itm.Quantity,
		itm.Availability,
		itm.CategoryID,
		itm.UserID,
		itm.CreatedAt,
		itm.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return itm, nil
}

func (repository *itemRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Item, error) {
	query := "SELECT * FROM items"

	rows, err := repository.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listOfItems []*entity.Item

	for rows.Next() {
		var item entity.Item

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Quantity,
			&item.Availability,
			&item.CategoryID,
			&item.UserID,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfItems = append(listOfItems, &item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfItems, nil
}

func (repository *itemRepositoryImpl) FindByID(ctx context.Context, itmID uuid.UUID) (*entity.Item, error) {
	var item entity.Item

	query := "SELECT * FROM items WHERE id = ?"

	err := repository.database.QueryRowContext(ctx, query, itmID).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.Quantity,
		&item.Availability,
		&item.CategoryID,
		&item.UserID,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrItemNotFound
		}

		return nil, err
	}

	return &item, nil
}

func (repository *itemRepositoryImpl) FindByName(ctx context.Context, name string) ([]*entity.Item, error) {
	query := "SELECT * FROM items WHERE name = ?"

	rows, err := repository.database.QueryContext(ctx, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrItemNotFound
		}

		return nil, err
	}
	defer rows.Close()

	var listOfItems []*entity.Item

	for rows.Next() {
		var item entity.Item

		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Quantity,
			&item.Availability,
			&item.CategoryID,
			&item.UserID,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfItems = append(listOfItems, &item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfItems, nil
}

func (repository *itemRepositoryImpl) FindCompleteByID(ctx context.Context, itmID uuid.UUID) (*entity.Item, error) {
	var item entity.Item
	var category entity.CategoryFiltered
	var user entity.UserFiltered

	query := `
		SELECT i.*, c.id, c.name, c.description, u.username, u.contact
		FROM items i
		LEFT JOIN categories c ON c.id = i.category_id
		LEFT JOIN users u ON u.id = i.user_id
		WHERE i.id = ?
	`

	err := repository.database.QueryRowContext(ctx, query, itmID).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.Quantity,
		&item.Availability,
		&item.CategoryID,
		&item.UserID,
		&item.CreatedAt,
		&item.UpdatedAt,
		&category.ID,
		&category.Name,
		&category.Description,
		&user.Username,
		&user.Contact,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrItemNotFound
		}
		return nil, err
	}

	item.Category = &category
	item.User = &user

	var location entity.LocationFiltered

	query2 := `
		SELECT l.id, l.description
		FROM categories c
		LEFT JOIN locations l
		ON l.id = c.location_id
		WHERE c.id = ?
	`

	err = repository.database.QueryRowContext(ctx, query2, category.ID).Scan(
		&location.ID,
		&location.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.ErrLocationNotFound
		}
		return nil, err
	}

	item.Location = &location

	return &item, nil
}

func (repository *itemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, itm *entity.Item) error {
	query := "UPDATE items SET name = ?, description = ?, quantity = ?, availability = ?, category_id = ?, user_id = ?, updated_at = ? WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		query,
		itm.Name,
		itm.Description,
		itm.Quantity,
		itm.Availability,
		itm.CategoryID,
		itm.UserID,
		itm.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.ErrItemNotFound
		}

		return err
	}
	defer repository.database.Close()

	return nil
}

func (repository *itemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, itmID uuid.UUID) error {
	query := "DELETE FROM items WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, itmID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return nil
	}

	defer repository.database.Close()

	return nil
}

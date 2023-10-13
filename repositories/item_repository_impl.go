package repositories

import (
	"database/sql"
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

var ErrItemNotFound = errors.New("item not found")

type itemRepositoryImpl struct {
	database *sql.DB
}

func NewItemRepository(database *sql.DB) ItemRepository {
	return &itemRepositoryImpl{
		database: database,
	}
}

func (repository *itemRepositoryImpl) Insert(itm *entity.Item) (*entity.Item, error) {
	query := "INSERT INTO items (id, name, description, quantity, availability, location_id, category_id, user_id, created_at, updated_at), VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := repository.database.Exec(
		query,
		itm.ID,
		itm.Name,
		itm.Description,
		itm.Quantity,
		itm.Availability,
		itm.LocationID,
		itm.CategoryID,
		itm.UserID,
		itm.CreatedAt,
		itm.UpdatedAt,
	)
	if err != nil {
		return &entity.Item{}, err
	}

	return itm, nil
}

func (repository *itemRepositoryImpl) FindAll() ([]*entity.Item, error) {
	query := "SELECT * FROM items"

	items, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer items.Close()

	var listOfItems []*entity.Item

	for items.Next() {
		var i entity.Item

		err := items.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Quantity,
			&i.Availability,
			&i.LocationID,
			&i.CategoryID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfItems = append(listOfItems, &i)
	}

	err = items.Err()
	if err != nil {
		return nil, err
	}

	return listOfItems, nil
}

func (repository *itemRepositoryImpl) FindByID(itmID uuid.UUID) (*entity.Item, error) {
	query := "SELECT * FROM items WHERE id = ?"

	sqlResult := repository.database.QueryRow(query, itmID)

	var item entity.Item
	err := sqlResult.Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.Quantity,
		&item.Availability,
		&item.LocationID,
		&item.CategoryID,
		&item.UserID,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &item, ErrItemNotFound
		}

		return &item, err
	}

	return &item, nil
}

func (repository *itemRepositoryImpl) FindByName(name string) ([]*entity.Item, error) {
	query := "SELECT * FROM items WHERE name = ?"

	items, err := repository.database.Query(query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrItemNotFound
		}

		return nil, err
	}
	defer items.Close()

	var listOfItems []*entity.Item

	for items.Next() {
		var i entity.Item

		err := items.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Quantity,
			&i.Availability,
			&i.LocationID,
			&i.CategoryID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfItems = append(listOfItems, &i)
	}

	err = items.Err()
	if err != nil {
		return nil, err
	}

	return listOfItems, nil
}

func (repository *itemRepositoryImpl) Update(itm *entity.Item) error {
	query := "UPDATE items SET name = ?, description = ?, quantity = ?, availability = ?, location_id = ?, category_id = ?, user_id = ?, updated_at = ? WHERE id = ?"

	_, err := repository.database.Exec(
		query,
		itm.Name,
		itm.Description,
		itm.Quantity,
		itm.Availability,
		itm.LocationID,
		itm.CategoryID,
		itm.UserID,
		itm.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrItemNotFound
		}

		return err
	}
	defer repository.database.Close()

	return nil
}

func (repository *itemRepositoryImpl) Delete(itmID uuid.UUID) error {
	query := "DELETE FROM users WHERE id = ?"

	_, err := repository.database.Exec(query, itmID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return nil
	}

	defer repository.database.Close()

	return nil
}

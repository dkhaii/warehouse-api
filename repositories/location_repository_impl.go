package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
)

type locationRepositoryImpl struct {
	database *sql.DB
}

func NewLocationRepository(database *sql.DB) LocationRepository {
	return &locationRepositoryImpl{
		database: database,
	}
}

func (repository *locationRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, loc *entity.Location) (*entity.Location, error) {
	query := `
		INSERT INTO locations 
		(id, description, created_at, updated_at) 
		VALUES 
		(?, ?, ?, ?)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		loc.ID,
		loc.Description,
		loc.CreatedAt,
		loc.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (repository *locationRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Location, error) {
	query := "SELECT * FROM locations"

	rows, err := repository.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listOfLocations []*entity.Location

	for rows.Next() {
		var location entity.Location

		err := rows.Scan(
			&location.ID,
			&location.Description,
			&location.CreatedAt,
			&location.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfLocations = append(listOfLocations, &location)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfLocations, nil
}

func (repository *locationRepositoryImpl) FindByID(ctx context.Context, locID string) (*entity.Location, error) {
	var location entity.Location

	query := "SELECT * FROM locations WHERE id = ?"

	err := repository.database.QueryRowContext(ctx, query, locID).Scan(
		&location.ID,
		&location.Description,
		&location.CreatedAt,
		&location.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrLocationNotFound
		}
		return nil, err
	}

	return &location, nil
}

func (repository *locationRepositoryImpl) FindCompleteByID(ctx context.Context, locID string) (*entity.Location, error) {
	var location entity.Location
	var categories []entity.CategoryFiltered

	query := `
		SELECT loc.*, ctg.id, ctg.name, ctg.description
		FROM locations loc 
		LEFT JOIN categories ctg
		ON loc.id  = ctg.location_id
		WHERE loc.id = ?
	`

	rows, err := repository.database.QueryContext(ctx, query, locID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hasNonEmptyCategory := false

	for rows.Next() {
		var category entity.CategoryFiltered
		var categoryID, categoryName, categoryDescription sql.NullString

		err := rows.Scan(
			&location.ID,
			&location.Description,
			&location.CreatedAt,
			&location.UpdatedAt,
			&categoryID,
			&categoryName,
			&categoryDescription,
		)
		if err != nil {
			return nil, err
		}

		category.ID = helpers.NullStringToString(categoryID)
		category.Name = helpers.NullStringToString(categoryName)
		category.Description = helpers.NullStringToString(categoryDescription)

		if category.ID != "" || category.Name != "" || category.Description != "" {
			hasNonEmptyCategory = true
		}	

		categories = append(categories, category)
	}

	if hasNonEmptyCategory {
		location.Category = categories
	}

	return &location, nil
}

func (repository *locationRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, loc *entity.Location) (*entity.Location, error) {
	query := `
		UPDATE locations 
		SET description = ?, updated_at = ? 
		WHERE id = ?
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		loc.Description,
		loc.UpdatedAt,
		loc.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrLocationNotFound
		}
		return nil, err
	}

	return loc, nil
}

func (repository *locationRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, locID string) error {
	query := "DELETE FROM locations WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, locID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrLocationNotFound
		}
		return err
	}

	return nil
}

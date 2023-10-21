package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
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

	_, err := repository.database.ExecContext(
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
	var categories []entity.Category

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

	for rows.Next() {
		var category entity.Category

		err := rows.Scan(
			&location.ID,
			&location.Description,
			&location.CreatedAt,
			&location.UpdatedAt,
			&category.ID,
			&category.Name,
			&category.Description,
		)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	location.Category = categories

	return &location, nil
}

func (repository *locationRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, loc *entity.Location) (*entity.Location, error) {
	query := `
	UPDATE locations 
	SET category_id = ?, description = ?, updated_at = ? 
	WHERE id = ?
	`

	_, err := repository.database.Exec(
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

	_, err := repository.database.Exec(query, locID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrLocationNotFound
		}
		return err
	}

	return nil
}

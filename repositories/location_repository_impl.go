package repositories

import (
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

func (repository *locationRepositoryImpl) Insert(loc *entity.Location) (*entity.Location, error) {
	query := `
	INSERT INTO locations 
	(id, category_id, description, created_at, updated_at) 
	VALUES 
	(?, ?, ?, ?, ?)
	`

	_, err := repository.database.Exec(
		query,
		loc.ID,
		loc.CategoryID,
		loc.Description,
		loc.CreatedAt,
		loc.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (repository *locationRepositoryImpl) FindAll() ([]*entity.Location, error) {
	query := "SELECT * FROM locations"

	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listOfLocations []*entity.Location

	for rows.Next() {
		var location entity.Location

		err := rows.Scan(
			&location.ID,
			&location.CategoryID,
			&location.Description,
			&location.CreatedAt,
			&location.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfLocations = append(listOfLocations, &location)
	}
	// rerr := rows.Close()
	// if rerr != nil {
	// 	return nil, err
	// }

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfLocations, nil
}

func (repository *locationRepositoryImpl) FindByID(locID string) (*entity.Location, error) {
	query := "SELECT * FROM locations WHERE id = ?"

	row := repository.database.QueryRow(query, locID)

	var location entity.Location
	err := row.Scan(
		&location.ID,
		&location.CategoryID,
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

func (repository *locationRepositoryImpl) FindCompleteByIDWithJoin(locID string) (*entity.Location, error) {
	var location entity.Location
	var category entity.Category

	query := `
	SELECT loc.*, ctg.* FROM locations loc 
	LEFT JOIN categories ctg 
	on ctg.id = loc.category_id 
	WHERE loc.id = ?
	`

	err := repository.database.QueryRow(query, locID).Scan(
		&location.ID,
		&location.CategoryID,
		&location.Description,
		&location.CreatedAt,
		&location.CreatedAt,
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrLocationNotFound
		}
		return nil, err
	}

	location.Category = &category

	return &location, nil
}

func (repository *locationRepositoryImpl) Update(loc *entity.Location) error {
	query := `
	UPDATE locations 
	SET category_id = ?, description = ?, updated_at = ? 
	WHERE id = ?
	`

	_, err := repository.database.Exec(
		query,
		loc.CategoryID,
		loc.Description,
		loc.UpdatedAt,
		loc.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrLocationNotFound
		}
		return err
	}

	return nil
}

func (repository *locationRepositoryImpl) Delete(locID string) error {
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

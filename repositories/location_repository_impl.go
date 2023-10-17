package repositories

import (
	"database/sql"
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
)

var ErrLocationNotFound = errors.New("location not found")

type locationRepositoryImpl struct {
	database *sql.DB
}

func NewLocationRepository(database *sql.DB) LocationRepository {
	return &locationRepositoryImpl{
		database: database,
	}
}

func (repository *locationRepositoryImpl) Insert(loc *entity.Location) (*entity.Location, error) {
	query := "INSERT INTO locations (id, category_id, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	_, err := repository.database.Exec(
		query,
		loc.ID,
		loc.CategoryID,
		loc.Description,
		loc.CreatedAt,
		loc.UpdatedAt,
	)
	if err != nil {
		return &entity.Location{}, err
	}

	return loc, nil
}

func (repository *locationRepositoryImpl) FindAll() ([]*entity.Location, error) {
	query := "SELECT * FROM locations"

	locations, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}

	var listOfLocations []*entity.Location

	for locations.Next() {
		var l entity.Location

		err := locations.Scan(
			&l.ID,
			&l.CategoryID,
			&l.Description,
			&l.CreatedAt,
			&l.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfLocations = append(listOfLocations, &l)
	}

	err = locations.Err()
	if err != nil {
		return nil, err
	}

	return listOfLocations, nil
}

func (repository *locationRepositoryImpl) FindByID(locID string) (*entity.Location, error) {
	query := "SELECT * FROM locations WHERE id = ?"

	sqlResult := repository.database.QueryRow(query, locID)

	var location entity.Location
	err := sqlResult.Scan(
		&location.ID,
		&location.CategoryID,
		&location.Description,
		&location.CreatedAt,
		&location.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &entity.Location{}, ErrLocationNotFound
		}
		return &entity.Location{}, err
	}

	return &location, nil
}

func (repository *locationRepositoryImpl) Update(loc *entity.Location) error {
	query := "UPDATE locations SET category_id = ?, description = ?, updated_at = ? WHERE id = ?"

	_, err := repository.database.Exec(
		query,
		loc.CategoryID,
		loc.Description,
		loc.UpdatedAt,
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

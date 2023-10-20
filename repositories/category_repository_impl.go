package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
)

type categoryRepositoryImpl struct {
	database *sql.DB
}

func NewCategoryRepository(database *sql.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		database: database,
	}
}

func (repository *categoryRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, ctg *entity.Category) (*entity.Category, error) {
	query := `
	INSERT INTO categories 
	(id, name, description, created_at, updated_at) 
	VALUES 
	(?, ?, ?, ?, ?)
	`

	_, err := repository.database.ExecContext(
		ctx,
		query,
		ctg.ID,
		ctg.Name,
		ctg.Description,
		ctg.CreatedAt,
		ctg.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return ctg, nil
}

func (respository *categoryRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Category, error) {
	query := "SELECT * FROM categories"

	rows, err := respository.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listOfCategories []*entity.Category

	for rows.Next() {
		var category entity.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfCategories = append(listOfCategories, &category)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfCategories, nil
}

func (repository *categoryRepositoryImpl) FindByID(ctx context.Context, ctgID string) (*entity.Category, error) {
	var category entity.Category

	query := "SELECT * FROM categories WHERE id = ?"

	err := repository.database.QueryRowContext(ctx, query, ctgID).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return &category, nil
}

func (repository *categoryRepositoryImpl) FindByName(ctx context.Context, name string) ([]*entity.Category, error) {
	query := "SELECT * FROM categories WHERE name LIKE ?"
	name = name + "%"

	rows, err := repository.database.QueryContext(ctx, query, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var listOfCategories []*entity.Category

	for rows.Next() {
		var category entity.Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfCategories = append(listOfCategories, &category)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return listOfCategories, nil
}

func (repository *categoryRepositoryImpl) FindByIDWithJoin(ctx context.Context, ctgID string) (*entity.Category, error) {
	var category entity.Category
	var location entity.Location

	query := `
	SELECT ctg.*, loc.* 
	FROM categories ctg
	LEFT JOIN locations loc
	ON loc.id = ctg.location_id
	WHERE ctg.id = ?
	`

	err := repository.database.QueryRowContext(ctx, query, ctgID).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.LocationID,
		&category.CreatedAt,
		&category.UpdatedAt,
		&location.ID,
		&location.Description,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	category.Location = &location

	return &category, nil
}

func (repository *categoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ctg *entity.Category) (*entity.Category, error) {
	query := `
	UPDATE categories 
	SET name = ?, description = ?, updated_at = ? 
	WHERE id = ?
	`

	_, err := repository.database.ExecContext(
		ctx,
		query,
		ctg.Name,
		ctg.Description,
		ctg.UpdatedAt,
		ctg.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}

	return ctg, nil
}

func (repository *categoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ctgID string) error {
	query := "DELETE FROM categories WHERE id = ?"

	_, err := repository.database.ExecContext(ctx, query, ctgID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrCategoryNotFound
		}
		return err
	}

	return nil
}

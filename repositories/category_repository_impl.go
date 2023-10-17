package repositories

import (
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

func (repository *categoryRepositoryImpl) Insert(ctg *entity.Category) (*entity.Category, error) {
	query := "INSERT INTO categories (id, name, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	_, err := repository.database.Exec(
		query,
		ctg.ID,
		ctg.Name,
		ctg.Description,
		ctg.CreatedAt,
		ctg.UpdatedAt,
	)
	if err != nil {
		return &entity.Category{}, err
	}

	return ctg, nil
}

func (respository *categoryRepositoryImpl) FindAll() ([]*entity.Category, error) {
	query := "SELECT * FROM categories"

	categories, err := respository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer categories.Close()

	var listOfCategories []*entity.Category

	for categories.Next() {
		var c entity.Category

		err := categories.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfCategories = append(listOfCategories, &c)
	}

	err = categories.Err()
	if err != nil {
		return nil, err
	}

	return listOfCategories, nil
}

func (repository *categoryRepositoryImpl) FindByID(ctgID string) (*entity.Category, error) {
	query := "SELECT * FROM categories WHERE id = ?"

	sqlResult := repository.database.QueryRow(query, ctgID)

	var category entity.Category
	err := sqlResult.Scan(
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

func (repository *categoryRepositoryImpl) FindByName(name string) ([]*entity.Category, error) {
	query := "SELECT * FROM categories WHERE name LIKE ?"
	name = name + "%"

	categories, err := repository.database.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer categories.Close()

	var listOfCategories []*entity.Category

	for categories.Next() {
		var c entity.Category

		err := categories.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		listOfCategories = append(listOfCategories, &c)
	}

	err = categories.Err()
	if err != nil {
		return nil, err
	}

	return listOfCategories, err
}

func (repository *categoryRepositoryImpl) Update(ctg *entity.Category) error {
	query := "UPDATE categories SET name = ?, description = ?, updated_at = ? WHERE id = ?"

	_, err := repository.database.Exec(
		query,
		ctg.Name,
		ctg.Description,
		ctg.UpdatedAt,
		ctg.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrCategoryNotFound
		}
		return err
	}

	return nil
}

func (repository *categoryRepositoryImpl) Delete(ctgID string) error {
	query := "DELETE FROM categories WHERE id = ?"

	_, err := repository.database.Exec(query, ctgID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrCategoryNotFound
		}
		return err
	}

	return nil
}

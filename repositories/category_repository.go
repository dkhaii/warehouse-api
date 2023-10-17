package repositories

import (
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepository interface {
	Insert(category *entity.Category) (*entity.Category, error)
	FindAll() ([]*entity.Category, error)
	FindByID(categoryID string) (*entity.Category, error)
	FindByName(name string) ([]*entity.Category, error)
	Update(category *entity.Category) error
	Delete(categoryID string) error
}
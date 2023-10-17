package repositories

import (
	"github.com/dkhaii/warehouse-api/entity"
)

type CategoryRepository interface {
	Insert(category *entity.Category) (*entity.Category, error)
	FindAll() ([]*entity.Category, error)
	FindByID(categoryID string) (*entity.Category, error)
	FindByName(name string) ([]*entity.Category, error)
	Update(category *entity.Category) error
	Delete(categoryID string) error
}

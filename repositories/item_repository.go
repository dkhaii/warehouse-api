package repositories

import (
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type ItemRepository interface {
	Insert(item *entity.Item) (*entity.Item, error)
	FindAll() ([]*entity.Item, error)
	FindByID(id uuid.UUID) (*entity.Item, error)
	FindByName(name string) ([]*entity.Item, error)
	Update(item *entity.Item) error
	Delete(id uuid.UUID) error
}

package repositories

import (
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

var ErrItemNotFound = errors.New("item not found")

type ItemRepository interface {
	Insert(item *entity.Item) (*entity.Item, error)
	FindAll() ([]*entity.Item, error)
	FindByID(itmID uuid.UUID) (*entity.Item, error)
	FindByName(name string) ([]*entity.Item, error)
	FindCompleteByIDWithJoin(itmID uuid.UUID) (*entity.Item, error)
	Update(item *entity.Item) error
	Delete(itmID uuid.UUID) error
}

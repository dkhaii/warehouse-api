package repositories

import (
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
)

var ErrLocationNotFound = errors.New("location not found")

type LocationRepository interface {
	Insert(loc *entity.Location) (*entity.Location, error)
	FindAll() ([]*entity.Location, error)
	FindByID(locID string) (*entity.Location, error)
	FindCompleteByIDWithJoin(locID string) (*entity.Location, error)
	Update(loc *entity.Location) error
	Delete(locID string) error
}

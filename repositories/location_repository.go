package repositories

import "github.com/dkhaii/warehouse-api/entity"

type LocationRepository interface {
	Insert(loc *entity.Location) (*entity.Location, error)
	FindAll() ([]*entity.Location, error)
	FindByID(locID string) (*entity.Location, error)
	Update(loc *entity.Location) error
	Delete(locID string) error
}

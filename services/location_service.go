package services

import "github.com/dkhaii/warehouse-api/models"

type LocationService interface {
	Create(request models.CreateLocationRequest) (models.CreateLocationResponse, error)
	GetAll() ([]models.GetLocationResponse, error)
	GetCompleteByID(locID string) (models.GetCompleteLocationByIDResponse, error)
	Update(request models.UpdateLocationRequest) error
	Delete(locID string) error
}

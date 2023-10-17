package services

import "github.com/dkhaii/warehouse-api/models"

type LocationService interface {
	Create(request models.CreateLocationRequest) (models.CreateLocationResponse, error)
	FindAll() ([]models.GetLocationResponse, error)
	FindCompleteByID(locID string) (models.GetCompleteLocationByIDResponse, error)
	Update(request models.UpdateLocationRequest) error
	Delete(locID string) error
}

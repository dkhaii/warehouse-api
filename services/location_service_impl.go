package services

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/internal/validationutil"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
)

type locationServiceImpl struct {
	locationRepository repositories.LocationRepository
}

func NewLocationService(locationRepository repositories.LocationRepository) LocationService {
	return &locationServiceImpl{
		locationRepository: locationRepository,
	}
}

func (service *locationServiceImpl) Create(request models.CreateLocationRequest) (models.CreateLocationResponse, error) {
	err := validationutil.ValidateRequest(request)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	createdAt := time.Now()

	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	location := entity.Location{
		ID:          request.ID,
		CategoryID:  request.CategoryID,
		Description: request.Description,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		Category:    nil,
	}

	createdLocation, err := service.locationRepository.Insert(&location)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	response := models.CreateLocationResponse{
		ID:          createdLocation.ID,
		CategoryID:  createdLocation.CategoryID,
		Description: createdLocation.Description,
		CreatedAt:   createdLocation.CreatedAt,
		UpdatedAt:   createdLocation.UpdatedAt,
	}

	return response, nil
}

func (service *locationServiceImpl) GetAll() ([]models.GetLocationResponse, error) {
	rows, err := service.locationRepository.FindAll()
	if err != nil {
		return nil, err
	}

	locations := make([]models.GetLocationResponse, len(rows))

	for key, location := range rows {
		locations[key] = models.GetLocationResponse{
			ID:          location.ID,
			CategoryID:  location.CategoryID,
			Description: location.Description,
			CreatedAt:   location.CreatedAt,
			UpdatedAt:   location.UpdatedAt,
		}
	}

	return locations, nil
}

func (service *locationServiceImpl) GetCompleteByID(locID string) (models.GetCompleteLocationByIDResponse, error) {
	location, err := service.locationRepository.FindCompleteByIDWithJoin(locID)
	if err != nil {
		return models.GetCompleteLocationByIDResponse{}, err
	}

	response := models.GetCompleteLocationByIDResponse{
		ID:          location.ID,
		CategoryID:  location.CategoryID,
		Description: location.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
		Category:    location.Category,
	}

	return response, nil
}

func (service *locationServiceImpl) Update(request models.UpdateLocationRequest) error {
	err := validationutil.ValidateRequest(request)
	if err != nil {
		return err
	}

	location, err := service.locationRepository.FindByID(request.ID)
	if err != nil {
		return err
	}

	request.UpdatedAt = time.Now()

	updatedLocation := entity.Location{
		ID:          location.ID,
		CategoryID:  request.CategoryID,
		Description: request.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		Category:    nil,
	}

	err = service.locationRepository.Update(&updatedLocation)
	if err != nil {
		return err
	}

	return nil
}

func (service *locationServiceImpl) Delete(locID string) error {
	location, err := service.locationRepository.FindByID(locID)
	if err != nil {
		return err
	}

	err = service.locationRepository.Delete(location.ID)
	if err != nil {
		return err
	}

	return nil
}

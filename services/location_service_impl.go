package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
)

type locationServiceImpl struct {
	locationRepository repositories.LocationRepository
	database           *sql.DB
}

func NewLocationService(locationRepository repositories.LocationRepository, database *sql.DB) LocationService {
	return &locationServiceImpl{
		locationRepository: locationRepository,
		database:           database,
	}
}

func (service *locationServiceImpl) Create(ctx context.Context, request models.CreateLocationRequest) (models.CreateLocationResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateLocationResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	createdAt := time.Now()
	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	location := entity.Location{
		ID:          request.ID,
		Description: request.Description,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		Category:    nil,
	}

	createdLocation, err := service.locationRepository.Insert(ctx, tx, &location)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	response := models.CreateLocationResponse{
		ID:          createdLocation.ID,
		Description: createdLocation.Description,
		CreatedAt:   createdLocation.CreatedAt,
		UpdatedAt:   createdLocation.UpdatedAt,
	}

	return response, nil
}

func (service *locationServiceImpl) GetAll(ctx context.Context) ([]models.GetLocationResponse, error) {
	rows, err := service.locationRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	locations := make([]models.GetLocationResponse, len(rows))

	for key, location := range rows {
		locations[key] = models.GetLocationResponse{
			ID:          location.ID,
			Description: location.Description,
			CreatedAt:   location.CreatedAt,
			UpdatedAt:   location.UpdatedAt,
		}
	}

	return locations, nil
}

func (service *locationServiceImpl) GetCompleteByID(ctx context.Context, locID string) (models.GetCompleteLocationResponse, error) {
	location, err := service.locationRepository.FindCompleteByID(ctx, locID)
	if err != nil {
		return models.GetCompleteLocationResponse{}, err
	}

	response := models.GetCompleteLocationResponse{
		ID:          location.ID,
		Description: location.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   location.UpdatedAt,
		Category:    location.Category,
	}

	return response, nil
}

func (service *locationServiceImpl) Update(ctx context.Context, request models.UpdateLocationRequest) (models.CreateLocationResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	location, err := service.locationRepository.FindByID(ctx, request.ID)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateLocationResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	request.UpdatedAt = time.Now()

	updatedLocation := entity.Location{
		ID:          location.ID,
		Description: request.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
		Category:    nil,
	}

	locationData, err := service.locationRepository.Update(ctx, tx, &updatedLocation)
	if err != nil {
		return models.CreateLocationResponse{}, err
	}

	response := models.CreateLocationResponse{
		ID:          locationData.ID,
		Description: locationData.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   locationData.UpdatedAt,
	}

	return response, nil
}

func (service *locationServiceImpl) Delete(ctx context.Context, locID string) error {
	location, err := service.locationRepository.FindByID(ctx, locID)
	if err != nil {
		return err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return err
	}
	defer helpers.CommitOrRollBack(tx)

	err = service.locationRepository.Delete(ctx, tx, location.ID)
	if err != nil {
		return err
	}

	return nil
}

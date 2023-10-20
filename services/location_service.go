package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
)

type LocationService interface {
	Create(ctx context.Context, request models.CreateLocationRequest) (models.CreateLocationResponse, error)
	GetAll(ctx context.Context) ([]models.GetLocationResponse, error)
	GetCompleteByID(ctx context.Context, locID string) (models.GetCompleteLocationByIDResponse, error)
	Update(ctx context.Context, request models.UpdateLocationRequest) (models.CreateLocationResponse, error)
	Delete(ctx context.Context, locID string) error
}

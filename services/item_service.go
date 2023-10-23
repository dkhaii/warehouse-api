package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type ItemService interface {
	Create(ctx context.Context, request models.CreateItemRequest) (models.CreateItemResponse, error)
	GetAll(ctx context.Context) ([]models.GetItemResponse, error)
	GetByID(ctx context.Context, itmID uuid.UUID) (models.GetItemResponse, error)
	GetByName(ctx context.Context, name string) ([]models.GetItemResponse, error)
	GetCompleteByID(ctx context.Context, itmID uuid.UUID) (models.GetCompleteItemResponse, error)
	Update(ctx context.Context, request models.UpdateItemRequest) (models.CreateItemResponse, error)
	Delete(ctx context.Context, itmID uuid.UUID) error
}

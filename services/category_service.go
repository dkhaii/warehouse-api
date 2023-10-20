package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
)

type CategoryService interface {
	Create(ctx context.Context, request models.CreateCategoryRequest) (models.CreateCategoryResponse, error)
	GetAll(ctx context.Context) ([]models.GetCategoryResponse, error)
	GetByID(ctx context.Context, ctgID string) (models.GetCategoryResponse, error)
	GetByName(ctx context.Context, name string) ([]models.GetCategoryResponse, error)
	Update(ctx context.Context, request models.UpdateCategoryRequest) (models.CreateCategoryResponse, error)
	Delete(ctx context.Context, ctgID string) error
}

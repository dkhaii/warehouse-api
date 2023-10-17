package services

import (
	"github.com/dkhaii/warehouse-api/models"
)

type CategoryService interface {
	Create(request models.CreateCategoryRequest) (models.CreateCategoryResponse, error)
	GetAll() ([]models.GetCategoryResponse, error)
	GetByID(ctgID string) (models.GetCategoryResponse, error)
	GetByName(name string) ([]models.GetCategoryResponse, error)
	Update(request models.UpdateCategoryRequest) error
	Delete(ctgID string) error
}

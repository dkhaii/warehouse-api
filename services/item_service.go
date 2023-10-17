package services

import (
	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type ItemService interface {
	Create(request models.CreateItemRequest) (models.CreateItemResponse, error)
	GetAll() ([]models.GetItemResponse, error)
	GetByID(itmID uuid.UUID) (models.GetItemResponse, error)
	GetByName(name string) ([]models.GetItemResponse, error)
	Update(request models.CreateItemRequest) (models.CreateItemResponse, error)
	Delete(itmID uuid.UUID) error
}
package service

import (
	"github.com/dkhaii/warehouse-api/model"
	"github.com/google/uuid"
)

type ItemService interface {
	Create(request model.CreateItemRequest) (model.CreateItemResponse, error)
	GetAll() ([]model.GetItemResponse, error)
	GetByID(itmID uuid.UUID) (model.GetItemResponse, error)
	GetByName(name string) ([]model.GetItemResponse, error)
	Update(request model.CreateItemRequest) (model.CreateItemResponse, error)
	Delete(itmID uuid.UUID) error
}

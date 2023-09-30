package service

import (
	"github.com/dkhaii/warehouse-api/model"
	"github.com/google/uuid"
)

type UserService interface {
	Create(request model.CreateUserRequest) (model.CreateUserResponse, error)
	GetByID(id uuid.UUID) (model.GetUserResponse, error)
	GetByName(name string) ([]model.GetUserResponse, error)
	Update(request model.CreateUserRequest) error
	Delete(id uuid.UUID) error
}

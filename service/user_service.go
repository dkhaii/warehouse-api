package service

import (
	"github.com/dkhaii/warehouse-api/model"
	"github.com/google/uuid"
)

type UserService interface {
	Create(request model.CreateUserRequest) (model.CreateUserResponse, error)
	GetAll() ([]model.GetUserResponse, error)
	GetByID(usrID uuid.UUID) (model.GetUserResponse, error)
	GetByUsername(name string) ([]model.GetUserResponse, error)
	Update(request model.CreateUserRequest) error
	Delete(usrID uuid.UUID) error
	// GetUsers(request model.GetUserRequest) ([]model.GetUserResponse, error)
	// Login(request model.LoginUserRequest) (model.LoginUserResponse, error)
}

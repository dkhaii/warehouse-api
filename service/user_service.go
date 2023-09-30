package service

import (
	"github.com/dkhaii/warehouse-api/model"
	"github.com/google/uuid"
)

type UserService interface {
	Create(request model.CreateUserRequest) (model.CreateUserResponse, error)
	GetByID(usrID uuid.UUID) (model.GetUserResponse, error)
	GetByUsername(name string) ([]model.GetUserResponse, error)
	Update(request model.CreateUserRequest) error
	Delete(usrID uuid.UUID) error
	// Login(request model.LoginUserRequest) (model.LoginUserResponse, error)
}

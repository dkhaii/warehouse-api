package services

import (
	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type UserService interface {
	Create(request models.CreateUserRequest) (models.CreateUserResponse, error)
	GetAll() ([]models.GetUserResponse, error)
	GetByID(usrID uuid.UUID) (models.GetUserResponse, error)
	GetByUsername(name string) ([]models.GetUserResponse, error)
	Update(request models.UpdateUserRequest) error
	Delete(usrID uuid.UUID) error
	// GetUsers(request models.GetUserRequest) ([]models.GetUserResponse, error)
	Login(request models.LoginUserRequest) (models.TokenResponse, error)
}

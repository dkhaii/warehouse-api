package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, request models.CreateUserRequest) (models.CreateUserResponse, error)
	GetAll(ctx context.Context) ([]models.GetUserResponse, error)
	GetCompleteByID(ctx context.Context, usrID uuid.UUID) (models.GetCompleteUserResponse, error)
	GetByUsername(ctx context.Context, name string) ([]models.GetUserResponse, error)
	Update(ctx context.Context, request models.UpdateUserRequest) (models.CreateUserResponse, error)
	Delete(ctx context.Context, usrID uuid.UUID) error
	Login(ctx context.Context, request models.LoginUserRequest) (models.LoginUserResponse, error)
}

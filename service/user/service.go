package user

import (
	"github.com/dkhaii/warehouse-api/domain/user"
	"github.com/google/uuid"
)

type UserService interface {
	Create(user *user.UserEntity) error
	FindByID(id uuid.UUID) (*user.UserEntity, error)
	FindByName(name string) (*user.UserEntity, error)
	Update(user *user.UserEntity) error
	Delete(id uuid.UUID) error
}

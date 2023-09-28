package repository

import (
	"github.com/dkhaii/warehouse-api/domain/user"
	"github.com/google/uuid"
)

type UserRepository interface {
	Insert(user *user.UserEntity) (*user.UserEntity, error)
	FindAll() ([]*user.UserEntity, error)
	FindByID(id uuid.UUID) (*user.UserEntity, error)
	FindByName(name string) ([]*user.UserEntity, error)
	Update(user *user.UserEntity) error
	Delete(id uuid.UUID) error
}

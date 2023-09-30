package repository

import (
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Insert(user *entity.User) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	FindByID(id uuid.UUID) (*entity.User, error)
	FindByUsername(name string) ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id uuid.UUID) error
}

package repositories

import (
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Insert(user *entity.User) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	FindByID(usrID uuid.UUID) (*entity.User, error)
	GetByUsername(name string) ([]*entity.User, error)
	FindByUsername(name string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(usrID uuid.UUID) error
}

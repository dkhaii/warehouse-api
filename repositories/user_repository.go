package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, usrID uuid.UUID) (*entity.User, error)
	FindCompleteByID(ctx context.Context, usrID uuid.UUID) (*entity.User, error)
	GetByUsername(ctx context.Context, name string) ([]*entity.User, error)
	FindByUsername(ctx context.Context, name string) (*entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, tx *sql.Tx,usrID uuid.UUID) error
}

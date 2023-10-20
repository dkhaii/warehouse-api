package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dkhaii/warehouse-api/entity"
)

var ErrLocationNotFound = errors.New("location not found")

type LocationRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, loc *entity.Location) (*entity.Location, error)
	FindAll(ctx context.Context) ([]*entity.Location, error)
	FindByID(ctx context.Context, locID string) (*entity.Location, error)
	FindCompleteByID(ctx context.Context, locID string) (*entity.Location, error)
	Update(ctx context.Context, tx *sql.Tx, loc *entity.Location) (*entity.Location, error)
	Delete(ctx context.Context, tx *sql.Tx, locID string) error
}

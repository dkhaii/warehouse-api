package repositories

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type TransferOrderRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, trfOrd *entity.TransferOrder) (*entity.TransferOrder, error)
	FindAll(ctx context.Context) ([]*entity.TransferOrder, error)
	FindByID(ctx context.Context, trfOrdID uuid.UUID) (*entity.TransferOrder, error)
	FindCompleteByOrderID(ctx context.Context, ordID uuid.UUID) (*entity.TransferOrder, error)
	Update(ctx context.Context, tx *sql.Tx, trfOrd *entity.TransferOrder) (*entity.TransferOrder, error)
}
package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransferOrder struct {
	ID            uuid.UUID
	OrderID       uuid.UUID
	UserID        uuid.UUID
	Status        string
	FulfilledDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Order         *Order
}

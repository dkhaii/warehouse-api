package transfer_order

import (
	"time"

	"github.com/google/uuid"
)

type TransferOrder struct {
	ID            uuid.UUID `json:"id"`
	OrderID       uuid.UUID `json:"order_id"`
	Status        int       `json:"status"`
	UserID        uuid.UUID `json:"user_id"`
	FulfilledDate time.Time `json:"fulfilled_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

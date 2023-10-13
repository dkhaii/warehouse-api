package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	ID                  uuid.UUID `json:"id"`
	ItemID              uuid.UUID `json:"item_id"`
	UserID              uuid.UUID `json:"user_id"`
	Quantity            int       `json:"quantity"`
	RequestTransferDate time.Time `json:"request_transfer_date"`
	Notes               string    `json:"notes"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

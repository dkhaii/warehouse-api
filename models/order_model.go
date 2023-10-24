package models

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	ID                  uuid.UUID `json:"id"`
	ItemID              uuid.UUID `json:"item_id" validate:"required"`
	UserID              uuid.UUID `json:"user_id" validate:"required"`
	Quantity            int       `json:"quantity" validate:"required"`
	RequestTransferDate time.Time `json:"request_transfer_date" validate:"required"`
	Notes               string    `json:"notes" validate:"required"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type CreateOrderResponse struct {
	ID                  uuid.UUID `json:"id"`
	ItemID              uuid.UUID `json:"item_id"`
	UserID              uuid.UUID `json:"user_id"`
	Quantity            int       `json:"quantity"`
	RequestTransferDate time.Time `json:"request_transfer_date"`
	Notes               string    `json:"notes"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type GetOrderResponse struct {
	ID                  uuid.UUID `json:"id"`
	ItemID              uuid.UUID `json:"item_id"`
	UserID              uuid.UUID `json:"user_id"`
	Quantity            int       `json:"quantity"`
	RequestTransferDate time.Time `json:"request_transfer_date"`
	Notes               string    `json:"notes"`
	CreatedAt           time.Time `json:"created_at"`
}

type GetOrderByIDQueryRequest struct {
	ID uuid.UUID `query:"id"`
}

type GetCompleteOrderResponse struct {
	ID                  uuid.UUID             `json:"id"`
	ItemID              uuid.UUID             `json:"item_id"`
	UserID              uuid.UUID             `json:"user_id"`
	Quantity            int                   `json:"quantity"`
	RequestTransferDate time.Time             `json:"request_transfer_date"`
	Notes               string                `json:"notes"`
	CreatedAt           time.Time             `json:"created_at"`
	User                *entity.UserFiltered  `json:"user"`
	Item                []entity.ItemFiltered `json:"item"`
}

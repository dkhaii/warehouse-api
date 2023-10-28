package models

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type CreateTransferOrderRequest struct {
	ID            uuid.UUID `json:"id"`
	OrderID       uuid.UUID `json:"order_id" validate:"required"`
	UserID        uuid.UUID `json:"user_id" validate:"required"`
	Status        string       `json:"status"`
	FulfilledDate time.Time `json:"fulfilled_date" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateTransferOrderResponse struct {
	ID            uuid.UUID `json:"id"`
	OrderID       uuid.UUID `json:"order_id"`
	UserID        uuid.UUID `json:"user_id"`
	Status        string       `json:"status"`
	FulfilledDate time.Time `json:"fulfilled_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetTransferOrderByIDParamRequest struct {
	ID uuid.UUID `param:"id"`
}

type GetTransferOrderResponse struct {
	ID            uuid.UUID `json:"id"`
	OrderID       uuid.UUID `json:"order_id"`
	UserID        uuid.UUID `json:"user_id"`
	Status        string       `json:"status"`
	FulfilledDate time.Time `json:"fulfilled_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetCompleteTransferOrderResponse struct {
	ID            uuid.UUID     `json:"id"`
	OrderID       uuid.UUID     `json:"order_id"`
	UserID        uuid.UUID     `json:"user_id"`
	Status        string           `json:"status"`
	FulfilledDate time.Time     `json:"fulfilled_date"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Order         *entity.Order `json:"order"`
}

type UpdateTransferOrderRequest struct {
	ID            uuid.UUID `param:"id"`
	OrderID       uuid.UUID `json:"order_id"`
	UserID        uuid.UUID `json:"user_id"`
	Status        string       `json:"status"`
	FulfilledDate time.Time `json:"fulfilled_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

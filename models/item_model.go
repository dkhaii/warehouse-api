package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateItemRequest struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Quantity     int       `json:"quantity" validate:"required"`
	Availability bool      `json:"availability" validate:"required"`
	LocationID   string    `json:"location_id" validate:"required"`
	CategoryID   string    `json:"category_id" validate:"required"`
	UserID       uuid.UUID `json:"category" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateItemResponse struct {
	ID           uuid.UUID `param:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	LocationID   string    `json:"location_id"`
	CategoryID   string    `json:"category_id"`
	UserID       uuid.UUID `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetItemRequest struct {
	ID   uuid.UUID `query:"id"`
	Name string    `query:"name"`
}

type GetItemResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	LocationID   string    `json:"location_id"`
	CategoryID   string    `json:"category_id"`
	UserID       uuid.UUID `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

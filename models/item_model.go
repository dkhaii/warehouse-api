package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateItemRequest struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	LocationID   uuid.UUID `json:"location_id"`
	CategoryID   int       `json:"category_id"`
	UserID       uuid.UUID `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateItemResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	LocationID   uuid.UUID `json:"location_id"`
	CategoryID   int       `json:"category_id"`
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
	LocationID   uuid.UUID `json:"location_id"`
	CategoryID   int       `json:"category_id"`
	UserID       uuid.UUID `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

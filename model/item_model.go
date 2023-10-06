package model

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
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

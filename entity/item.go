package entity

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Quantity     int
	Availability bool
	LocationID   uuid.UUID
	CategoryID   int
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

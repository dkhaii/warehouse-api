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
	LocationID   string
	CategoryID   string
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

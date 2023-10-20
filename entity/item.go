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
	Location     *Location // has one relation
	Category     *Category // has one relation
	User         *User // has one relation
}

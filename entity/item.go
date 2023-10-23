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
	CategoryID   string
	UserID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Category     *CategoryFiltered // has one relation
	User         *UserFiltered     // has one relation
	Location     *LocationFiltered // has one relation
}

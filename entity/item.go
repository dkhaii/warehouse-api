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

type ItemFiltered struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Availability bool      `json:"availability"`
	CategoryID   string    `json:"category_id"`
}

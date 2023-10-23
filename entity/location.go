package entity

import (
	"time"
)

type Location struct {
	ID          string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    []CategoryFiltered
}

type LocationFiltered struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

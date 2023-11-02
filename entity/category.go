package entity

import (
	// "database/sql"
	"time"
)

type Category struct {
	ID          string
	Name        string
	Description string
	LocationID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Location    *Location
}

type CategoryFiltered struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

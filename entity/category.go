package entity

import (
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

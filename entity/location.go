package entity

import (
	"time"
)

type Location struct {
	ID          string
	CategoryID  string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

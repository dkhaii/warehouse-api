package entity

import (
	"time"
)

type Location struct {
	ID          string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Category    []Category
}

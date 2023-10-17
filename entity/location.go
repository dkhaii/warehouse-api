package entity

import (
	"time"
)

type Location struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

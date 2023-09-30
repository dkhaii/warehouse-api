package entity

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

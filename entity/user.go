package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Contact   string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
}

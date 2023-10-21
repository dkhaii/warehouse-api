package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Contact   string
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      *RoleFiltered
}

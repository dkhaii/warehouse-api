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

type UserFiltered struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Contact  string    `json:"contact"`
}

type UserClaim struct {
	ID       uuid.UUID
	Username string
	Contact  string
	RoleID   int
}

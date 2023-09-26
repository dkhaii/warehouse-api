package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Contact string    `json:"contact"`
	Role    int       `json:"role"`
}

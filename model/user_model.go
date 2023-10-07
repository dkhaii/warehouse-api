package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Contact   string    `json:"contact"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Contact   string    `json:"contact"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Contact   string    `json:"contact"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type GetUserIDRequest struct {
	ID uuid.UUID `param:"id"`
}

type GetUserRequest struct {
	ID       uuid.UUID `query:"id"`
	Username string    `query:"username"`
}

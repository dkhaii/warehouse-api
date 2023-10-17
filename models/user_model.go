package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Contact   string    `json:"contact" validate:"required"`
	Role      int       `json:"role" validate:"required"`
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
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type GetUserIDRequest struct {
	ID uuid.UUID `param:"id"`
}

type GetUserRequest struct {
	ID       uuid.UUID `query:"id"`
	Username string    `query:"username"`
}

type UpdateUserRequest struct {
	ID        uuid.UUID `param:"id"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Contact   string    `json:"contact" validate:"required"`
	Role      int       `json:"role" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
}

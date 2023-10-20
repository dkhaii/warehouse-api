package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Contact  string    `json:"contact"`
	RoleID   int       `json:"role_id"`
	jwt.RegisteredClaims
}

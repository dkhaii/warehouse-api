package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Contact  string    `json:"contact"`
	Role     int       `json:"role"`
	jwt.RegisteredClaims
}

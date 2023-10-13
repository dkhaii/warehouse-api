package tokenutil

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *entity.User, jwtSecret string, expire int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire))

	claims := &models.JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		Contact:  user.Contact,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := rawToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

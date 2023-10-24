package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateAccessToken(user *entity.User, jwtSecret string, expire int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire))

	claims := &models.JWTClaims{
		ID:       user.ID,
		Username: user.Username,
		Contact:  user.Contact,
		RoleID:   user.RoleID,
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

func GetSplitedToken(app echo.Context) (string, error) {
	currentUserToken := app.Request().Header.Get("Authorization")

	if currentUserToken == "" {
		return "", errors.New("uthorization header is missing")
	}

	splitToken := strings.Split(currentUserToken, " ")

	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", errors.New("nvalid token format")
	}

	splitedCurrentUserToken := splitToken[1]

	return splitedCurrentUserToken, nil
}

func GetUserClaimsFromToken(tokenStr string, jwtSecret string) (*entity.UserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTClaims{}, func(tkn *jwt.Token) (interface{}, error) {
		_, ok := tkn.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if ok && token.Valid {
		user := &entity.UserClaim{
			ID:       claims.ID,
			Username: claims.Username,
			Contact:  claims.Contact,
			RoleID:   claims.RoleID,
		}
		return user, nil
	}
	return nil, errors.New("invalid token or claims")
}

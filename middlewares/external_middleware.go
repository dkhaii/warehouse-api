package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/labstack/echo/v4"
)

func ExternalMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(app echo.Context) error {
			tokenString := app.Request().Header.Get("Authorization")

			if tokenString == "" {
				return helpers.CreateResponseError(
					app,
					http.StatusUnauthorized,
					errors.New("JWT AUTH IS NOT VALID, Authorization header is missing"),
				)
			}

			splitToken := strings.Split(tokenString, " ")
			if len(splitToken) != 2 || splitToken[0] != "Bearer" {
				return helpers.CreateResponseError(
					app,
					http.StatusUnauthorized,
					errors.New("JWT AUTH IS NOT VALID, Invalid token format"),
				)
			}

			tokenString = splitToken[1]

			config, err := config.Init()
			if err != nil {
				return err
			}

			jwtSecret := config.GetString("JWT_SECRET")

			currentUser, err := helpers.GetUserClaimsFromToken(tokenString, jwtSecret)
			if err != nil {
				return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
			}

			if currentUser.RoleID == 1 || currentUser.RoleID == 3 {
				return next(app)
			}

			return helpers.CreateResponseError(app, http.StatusForbidden, errors.New("insufficient role access"))
		}
	}
}

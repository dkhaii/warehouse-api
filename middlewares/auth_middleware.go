package middlewares

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(app echo.Context) error {
			username, password, ok := app.Request().BasicAuth()

			if !ok || subtle.ConstantTimeCompare([]byte(username), []byte("testing")) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte("1234567890")) == 1 {
				return echo.ErrUnauthorized
			}

			return next(app)
		}
	}
}

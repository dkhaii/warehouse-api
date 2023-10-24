package middlewares

import (
	"net/http"
	"strings"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(app echo.Context) error {
			tokenString := app.Request().Header.Get("Authorization")

			if tokenString == "" {
				return app.JSON(http.StatusUnauthorized, models.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "JWT AUTH IS NOT VALID",
					Data:   "Authorization header is missing",
				})
			}

			splitToken := strings.Split(tokenString, " ")
			if len(splitToken) != 2 || splitToken[0] != "Bearer" {
				return app.JSON(http.StatusUnauthorized, models.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "JWT AUTH IS NOT VALID",
					Data:   "Invalid token format",
				})
			}

			tokenString = splitToken[1]

			config, err := config.Init()
			if err != nil {
				return err
			}

			jwtSecret := config.GetString("JWT_SECRET")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})
			if err != nil || !token.Valid {
				return app.JSON(http.StatusUnauthorized, models.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "JWT AUTH IS NOT VALID",
					Data:   err.Error(),
				})
			}

			return next(app)
		}
	}
}

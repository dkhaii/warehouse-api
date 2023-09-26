package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes(app *echo.Echo) {
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "warehouse api")
	})
}

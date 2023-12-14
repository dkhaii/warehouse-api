package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/labstack/echo/v4"
)

func WelcomeMessage(app *echo.Echo, controller controller.WelcomeController) {
	app.GET("/", controller.WelcomeMessage)
}

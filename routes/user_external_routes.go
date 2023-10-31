package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProtectedUserExternalRoutes(app *echo.Echo, controller controller.UserExternalController) {
	routes := app.Group("/api/v1/auth/order")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())

	routes.POST("/create", controller.CreateOrder)
	// routes.GET("/find", controller.GetOrder)
}

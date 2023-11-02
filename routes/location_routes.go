package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProtectedLocationRoutes(app *echo.Echo, controller controller.LocationController) {
	routes := app.Group("/api/v1/auth/location")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	routes.Use(middlewares.AdminMiddleware())

	routes.POST("/create", controller.Create)
	routes.GET("/find", controller.GetLocation)
	routes.PUT("/update/:id", controller.Update)
	routes.DELETE("/delete/:id", controller.Delete)
}

package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ExternalUserOrderRoutes(app *echo.Echo, controller controller.UserExternalController) {
	routes := app.Group("/api/v1/auth/order")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	// routes.Use(middlewares.ExternalMiddleware())
	routes.Use(middlewares.AuthMiddleware(1, 2, 3))

	routes.POST("/create", controller.CreateOrder)
	routes.GET("/find", controller.GetAllOrderByUser)
}

func StaffUserOrderRoutes(app *echo.Echo, controller controller.OrderController) {
	routes := app.Group("/api/v1/staff/order")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	// routes.Use(middlewares.ExternalMiddleware())
	routes.Use(middlewares.AuthMiddleware(1, 2))

	routes.POST("/find", controller.GetOrder)
}
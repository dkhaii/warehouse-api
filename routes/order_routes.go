package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProtectedOrderRoutes(app *echo.Echo, ctrlOrder controller.OrderController, ctrlUserExternal controller.UserExternalController) {
	routes := app.Group("/api/v1/auth/order")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())

	routes.POST("/create", ctrlUserExternal.CreateOrder)
	routes.GET("/find", ctrlOrder.GetOrder)
}

package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ProtectedTransferOrderRoutes(app *echo.Echo, controller controller.TransferOrderController) {
	routes := app.Group("/api/v1/auth/transfer-order")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())

	routes.GET("/find", controller.GetTransferOrder)
	routes.PUT("/update/:id", controller.Update)
}
package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StaffUserTransferOrderRoutes(app *echo.Echo, controller controller.TransferOrderController) {
	routes := app.Group("/api/v1/staff/transfer-order")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	// routes.Use(middlewares.StaffMiddleware())
	routes.Use(middlewares.AuthMiddleware(1, 2))

	routes.GET("/find", controller.GetTransferOrder)
	routes.PUT("/update/:id", controller.Update)
}

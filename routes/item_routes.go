package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ExternalUserItemRoutes(app *echo.Echo, controller controller.UserExternalController) {
	routes := app.Group("api/v1/auth/item")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	routes.Use(middlewares.ExternalMiddleware())

	routes.GET("/find", controller.GetItem)
}

func StaffUserItemRoutes(app *echo.Echo, controller controller.ItemController) {
	routes := app.Group("api/v1/staff/item")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	routes.Use(middlewares.StaffMiddleware())

	routes.POST("/create", controller.Create)
	routes.GET("/find", controller.GetItem)
	routes.GET("/find/:id", controller.GetCompleteByID)
	routes.PUT("/update/:id", controller.Update)
	routes.DELETE("/delete/:id", controller.Delete)
}

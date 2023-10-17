package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func PublicUserRoutes(app *echo.Echo, controller *controller.UserController) {
	routes := app.Group("api/v1/users")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())

	routes.POST("/login", controller.Login)
	routes.POST("/register", controller.Create)
}

func ProtectedUserRoutes(app *echo.Echo, controller *controller.UserController) {
	routes := app.Group("api/v1/auth/users")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())

	routes.GET("/find", controller.GetUser)
	routes.PUT("/update/:id", controller.Update)
	routes.DELETE("/delete/:id", controller.Delete)
}

package routes

import (
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AdminUserCategoryRoutes(app *echo.Echo, controller controller.CategoryController) {
	routes := app.Group("api/v1/admin/category")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	// routes.Use(middlewares.AdminMiddleware())
	routes.Use(middlewares.AuthMiddleware(1))

	routes.POST("/create", controller.Create)
	routes.GET("/find", controller.GetCategory)
	routes.PUT("/update/:id", controller.Update)
	routes.DELETE("/delete/:id", controller.Delete)
}

func StaffUserCategoryRoutes(app *echo.Echo, controller controller.CategoryController) {
	routes := app.Group("api/v1/staff/category")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middlewares.JWTMiddleware())
	// routes.Use(middlewares.StaffMiddleware())
	routes.Use(middlewares.AuthMiddleware(2))

	routes.GET("/find", controller.GetCategory)
}

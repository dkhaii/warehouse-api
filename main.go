package main

import (
	"fmt"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/dkhaii/warehouse-api/routes"
	"github.com/dkhaii/warehouse-api/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	configuration, err := config.Init()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	database, err := config.NewMySQLDatabase(configuration)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// repository dependency injection
	userRepository := repositories.NewUserRepository(database)
	itemRepository := repositories.NewItemRepository(database)
	categoryRepository := repositories.NewCategoryRepository(database)
	locationRepository := repositories.NewLocationRepository(database)
	orderRepository := repositories.NewOrderRepository(database)

	// service dependency injection
	userService := services.NewUserService(userRepository, database)
	itemService := services.NewItemService(itemRepository, database)
	categoryService := services.NewCategoryService(categoryRepository, database)
	locationService := services.NewLocationService(locationRepository, database)
	orderService := services.NewOrderService(orderRepository, database)

	// controller dependency injection
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)
	categoryController := controller.NewCategoryController(categoryService)
	locationController := controller.NewLocationController(locationService)
	orderController := controller.NewOrderController(orderService)

	app := echo.New()

	// router
	routes.PublicUserRoutes(app, &userController)
	routes.ProtectedUserRoutes(app, &userController)
	routes.ProtectedItemRoutes(app, &itemController)
	routes.ProtectedCategoryRoutes(app, &categoryController)
	routes.ProtectedLocationRoutes(app, &locationController)
	routes.ProtectedOrderRoutes(app, &orderController)

	app.Logger.Fatal(app.Start(":8080"))
}

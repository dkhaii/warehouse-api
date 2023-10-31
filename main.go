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
	orderCartRepository := repositories.NewOrderCartRepository(database)
	transferOrderRepository := repositories.NewTransferOrderRepository(database)

	// service dependency injection
	userService := services.NewUserService(userRepository, database)
	itemService := services.NewItemService(itemRepository, database)
	categoryService := services.NewCategoryService(categoryRepository, database)
	locationService := services.NewLocationService(locationRepository, database)
	transferOrderService := services.NewTransferOrderService(transferOrderRepository, database)
	orderService := services.NewOrderService(orderRepository, orderCartRepository, transferOrderRepository, database)
	orderCartService := services.NewOrderCartService(orderCartRepository, database)
	userExternalService := services.NewUserExternalService(orderService, orderCartService, transferOrderService, database)

	// controller dependency injection
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)
	categoryController := controller.NewCategoryController(categoryService)
	locationController := controller.NewLocationController(locationService)
	orderController := controller.NewOrderController(orderService)
	userExternalController := controller.NewUserExternalController(userExternalService)

	app := echo.New()

	// router
	routes.PublicUserRoutes(app, userController)
	routes.ProtectedUserRoutes(app, userController)
	routes.ProtectedItemRoutes(app, itemController)
	routes.ProtectedCategoryRoutes(app, categoryController)
	routes.ProtectedLocationRoutes(app, locationController)
	routes.ProtectedOrderRoutes(app, orderController, userExternalController)

	app.Logger.Fatal(app.Start(":8080"))
}

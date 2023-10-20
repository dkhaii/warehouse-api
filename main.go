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
	configuration, err := config.New()
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

	// service dependency injection
	userService := services.NewUserService(userRepository)
	itemService := services.NewItemService(itemRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	locationService := services.NewLocationService(locationRepository, database)

	// controller dependency injection
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)
	categoryController := controller.NewCategoryController(categoryService)
	locationController := controller.NewLocationController(locationService)

	app := echo.New()

	// router
	routes.PublicUserRoutes(app, &userController)
	routes.ProtectedUserRoutes(app, &userController)
	itemController.Routes(app)
	routes.ProtectedCategoryRoutes(app, &categoryController)
	routes.ProtectedLocationRoutes(app, &locationController)

	app.Logger.Fatal(app.Start(":8080"))
}

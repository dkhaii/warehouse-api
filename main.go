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

	// service dependency injection
	userService := services.NewUserService(userRepository)
	itemService := services.NewItemService(itemRepository)

	// controller dependency injection
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)

	app := echo.New()

	// router
	routes.PublicUserRoutes(app, &userController)
	routes.AuthenticatedUserRoutes(app, &userController)
	itemController.Routes(app)

	app.Logger.Fatal(app.Start(":8080"))
}

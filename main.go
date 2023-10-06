package main

import (
	"fmt"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/repository"
	"github.com/dkhaii/warehouse-api/service"
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
	userRepository := repository.NewUserRepository(database)
	itemRepository := repository.NewItemRepository(database)

	// service dependency injection
	userService := service.NewUserService(userRepository)
	itemService := service.NewItemService(itemRepository)

	// controller dependency injection
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)

	app := echo.New()

	userController.Routes(app)
	itemController.Routes(app)

	app.Logger.Fatal(app.Start(":8080"))
}

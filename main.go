package main

import (
	"fmt"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/repository"
	"github.com/dkhaii/warehouse-api/service"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
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

	// repository implementation
	userRepository := repository.NewUserRepository(database)

	// service implementation
	userService := service.NewUserService(userRepository)

	// controller implementation
	userController := controller.NewUserController(userService)

	app := echo.New()

	userController.Routes(app)

	app.Logger.Fatal(app.Start(":8080"))
}

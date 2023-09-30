package main

import (
	"fmt"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/repository"
	"github.com/dkhaii/warehouse-api/service"
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
	}

	userRepository := repository.NewUserRepository(database)

	userService := service.NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	app := echo.New()

	userController.Routes(app)

	app.Logger.Fatal(app.Start(":8080"))
}

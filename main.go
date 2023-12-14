package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/controller"
	"github.com/dkhaii/warehouse-api/database/seeder"
	_ "github.com/dkhaii/warehouse-api/docs"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/dkhaii/warehouse-api/routes"
	"github.com/dkhaii/warehouse-api/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Cozy Warehouse API
// @version 1.0
// @description Inventory Management System.
// @termsOfService http://swagger.io/terms/

// @contact.name Mordekhai Gerin
// @contact.email mordekhaigerinlumangkun@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// configuration initialization
	cfg, err := config.Init()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	appENV := cfg.GetString("APP_ENV")
	var database *sql.DB
	var port string

	switch appENV {
	case "development":
		database, err = config.NewMySQLDatabase(cfg)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		port = cfg.GetString("APP_PORT")
	case "production":
		database, err = config.NewMySQLCloudDatabase(cfg)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		port = os.Getenv("PORT")
	default:
		log.Fatalf("unknown environment")
	}

	// database seeder
	err = seeder.RolesSeed(database)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = seeder.AdminUserSeed(database)
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
	userExternalService := services.NewUserExternalService(orderService, orderCartService, transferOrderService, itemService)

	// controller dependency injection
	welcomeController := controller.NewWelcomeController()
	userController := controller.NewUserController(userService)
	itemController := controller.NewItemController(itemService)
	categoryController := controller.NewCategoryController(categoryService)
	locationController := controller.NewLocationController(locationService)
	orderController := controller.NewOrderController(orderService)
	userExternalController := controller.NewUserExternalController(userExternalService)
	transferOrderController := controller.NewTransferOrderController(transferOrderService)

	app := echo.New()

	// swagger
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	// router
	routes.WelcomeMessage(app, welcomeController)
	routes.PublicUserRoutes(app, userController)

	routes.AdminUserRoutes(app, userController)
	routes.AdminUserLocationRoutes(app, locationController)
	routes.AdminUserCategoryRoutes(app, categoryController)

	routes.StaffUserItemRoutes(app, itemController)
	routes.StaffUserCategoryRoutes(app, categoryController)
	routes.StaffUserLocationRoutes(app, locationController)
	routes.StaffUserOrderRoutes(app, orderController)
	routes.StaffUserTransferOrderRoutes(app, transferOrderController)

	routes.ExternalUserItemRoutes(app, userExternalController)
	routes.ExternalUserOrderRoutes(app, userExternalController)

	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	err = app.Start(":" + port)
	if err != nil {
		app.Logger.Fatalf("application failed to start: %w", err)
	}
}

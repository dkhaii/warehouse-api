package main

import (
	"github.com/dkhaii/warehouse-api/infrastructure/http"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	http.Routes(app)

	app.Logger.Fatal(app.Start(":8080"))
}

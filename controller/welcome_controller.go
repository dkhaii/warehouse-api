package controller

import "github.com/labstack/echo/v4"

type WelcomeController interface {
	WelcomeMessage(app echo.Context) error
}

package controller

import "github.com/labstack/echo/v4"

type UserExternalController interface {
	CreateOrder(app echo.Context) error
	GetAllOrder(app echo.Context) error
	GetItem(app echo.Context) error
}

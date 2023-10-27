package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	Create(app echo.Context) error
	GetOrder(app echo.Context) error
}
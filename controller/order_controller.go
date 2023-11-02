package controller

import "github.com/labstack/echo/v4"

type OrderController interface {
	GetOrder(app echo.Context) error
}
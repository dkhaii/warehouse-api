package controller

import "github.com/labstack/echo/v4"

type TransferOrderController interface {
	Create(app echo.Context) error
	FindTransferOrder(app echo.Context) error
	Update(app echo.Context) error
}
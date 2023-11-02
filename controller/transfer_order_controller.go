package controller

import "github.com/labstack/echo/v4"

type TransferOrderController interface {
	GetTransferOrder(app echo.Context) error
	Update(app echo.Context) error
}
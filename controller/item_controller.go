package controller

import "github.com/labstack/echo/v4"

type ItemController interface {
	Create(app echo.Context) error
	GetItem(app echo.Context) error
	GetCompleteByID(app echo.Context) error
	Update(app echo.Context) error
	Delete(app echo.Context) error
}
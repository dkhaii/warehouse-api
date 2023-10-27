package controller

import "github.com/labstack/echo/v4"

type CategoryController interface {
	Create(app echo.Context) error
	GetCategory(app echo.Context) error
	Update(app echo.Context) error
	Delete (app echo.Context) error
}
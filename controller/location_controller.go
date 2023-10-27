package controller

import "github.com/labstack/echo/v4"

type LocationController interface {
	Create(app echo.Context) error
	GetLocation(app echo.Context) error
	Update(app echo.Context) error
	Delete(app echo.Context) error
}
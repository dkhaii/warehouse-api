package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Create(app echo.Context) error
	GetUser(app echo.Context) error
	Update(app echo.Context) error
	Delete(app echo.Context) error
	Login(app echo.Context) error
}
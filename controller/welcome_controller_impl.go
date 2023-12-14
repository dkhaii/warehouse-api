package controller

import (
	"net/http"

	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/labstack/echo/v4"
)

type welcomeControllerImpl struct{}

func NewWelcomeController() WelcomeController {
	return &welcomeControllerImpl{}
}

func (controller *welcomeControllerImpl) WelcomeMessage(app echo.Context) error {
	response := models.GetWelcomeMesssage{
		AppName:   "Cozy Warehouse RESTful API",
		Developer: "SI TAMPAN DAN PERKASA",
		Message:   "Welcome! untuk info lebih lanjut silahkan hubungi developer",
	}

	return helpers.CreateResponse(app, http.StatusOK, response)
}

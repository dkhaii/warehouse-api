package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/services"
	"github.com/labstack/echo/v4"
)

type locationControllerImpl struct {
	locationService services.LocationService
}

func NewLocationController(locationService services.LocationService) LocationController {
	return &locationControllerImpl{
		locationService: locationService,
	}
}

func (controller *locationControllerImpl) Create(app echo.Context) error {
	var request models.CreateLocationRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.locationService.Create(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}
	return helpers.CreateResponse(app, http.StatusCreated, response)
}

func (controller *locationControllerImpl) GetLocation(app echo.Context) error {
	var queryParam models.GetLocationByIDQueryRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancel()

	if queryParam.ID == "" {
		response, err := controller.locationService.GetAll(ctx)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusOK, response)
	}

	response, err := controller.locationService.GetCompleteByID(ctx, queryParam.ID)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *locationControllerImpl) Update(app echo.Context) error {
	var request models.UpdateLocationRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.locationService.Update(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *locationControllerImpl) Delete(app echo.Context) error {
	var urlParam models.GetLocationByIDParamRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	err = controller.locationService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		return helpers.CreateResponse(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, nil)
}

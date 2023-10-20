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

type LocationController struct {
	LocationService services.LocationService
}

func NewLocationController(locationService services.LocationService) LocationController {
	return LocationController{
		LocationService: locationService,
	}
}

func (controller *LocationController) Create(app echo.Context) error {
	var request models.CreateLocationRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.LocationService.Create(app.Request().Context(), request)
	if err != nil {
		return app.JSON(http.StatusNotFound, models.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	return app.JSON(http.StatusCreated, models.WebResponse{
		Code:   http.StatusCreated,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *LocationController) GetLocation(app echo.Context) error {
	var queryParam models.GetLocationByIDQueryRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancel()

	if queryParam.ID == "" {
		response, err := controller.LocationService.GetAll(ctx)
		if err != nil {
			if err == context.DeadlineExceeded {
				return app.JSON(http.StatusRequestTimeout, models.WebResponse{
					Code:   http.StatusRequestTimeout,
					Status: "FAIL",
					Data:   helpers.ErrRequestTimedOut,
				})
			}
			return app.JSON(http.StatusNotFound, models.WebResponse{
				Code:   http.StatusNotFound,
				Status: "FAIL",
				Data:   err.Error(),
			})
		}
		return app.JSON(http.StatusFound, models.WebResponse{
			Code:   http.StatusFound,
			Status: "SUCCESS",
			Data:   response,
		})
	}

	response, err := controller.LocationService.GetCompleteByID(ctx, queryParam.ID)
	if err != nil {
		if err == context.DeadlineExceeded {
			return app.JSON(http.StatusRequestTimeout, models.WebResponse{
				Code:   http.StatusRequestTimeout,
				Status: "FAIL",
				Data:   helpers.ErrRequestTimedOut,
			})
		}
		return app.JSON(http.StatusNotFound, models.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	return app.JSON(http.StatusFound, models.WebResponse{
		Code:   http.StatusFound,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *LocationController) Update(app echo.Context) error {
	var request models.UpdateLocationRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.LocationService.Update(app.Request().Context(), request)
	if err != nil {
		return app.JSON(http.StatusNotFound, models.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	return app.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *LocationController) Delete(app echo.Context) error {
	var urlParam models.GetLocationByIDParamRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	err = controller.LocationService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   nil,
	})
}

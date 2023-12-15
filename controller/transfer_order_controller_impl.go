package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type transferOrderControllerImpl struct {
	transferOrderService services.TransferOrderService
}

func NewTransferOrderController(transferOrderService services.TransferOrderService) TransferOrderController {
	return &transferOrderControllerImpl{
		transferOrderService: transferOrderService,
	}
}

func (controller *transferOrderControllerImpl) GetTransferOrder(app echo.Context) error {
	var queryParam models.GetTransferOrderQueryRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancle()

	if queryParam.ID != uuid.Nil {
		response, err := controller.transferOrderService.GetByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
		}
		return helpers.CreateResponse(app, http.StatusOK, response)
	}

	if queryParam.OrderID != uuid.Nil {
		response, err := controller.transferOrderService.GetCompleteByOrderID(ctx, queryParam.OrderID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
		}
		return helpers.CreateResponse(app, http.StatusOK, response)
	}

	response, err := controller.transferOrderService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}

	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *transferOrderControllerImpl) Update(app echo.Context) error {
	var request models.UpdateTransferOrderRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	currentUserToken, err := helpers.GetSplitedToken(app)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	response, err := controller.transferOrderService.Update(app.Request().Context(), request, currentUserToken)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}

	return helpers.CreateResponse(app, http.StatusOK, response)
}

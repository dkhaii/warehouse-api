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

type OrderController struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return OrderController{
		OrderService: orderService,
	}
}

func (controller *OrderController) Create(app echo.Context) error {
	var request models.CreateOrderRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	currentUserToken, err := helpers.GetSplitedToken(app)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	response, err := controller.OrderService.Create(app.Request().Context(), request, currentUserToken)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}
	return helpers.CreateResponse(app, http.StatusCreated, response)
}

func (controller *OrderController) GetOrder(app echo.Context) error {
	var urlParam models.GetOrderByIDQueryRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancle()

	if urlParam.ID != uuid.Nil {
		response, err := controller.OrderService.GetCompleteByID(ctx, urlParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	response, err := controller.OrderService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

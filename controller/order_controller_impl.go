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

type orderControllerImpl struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &orderControllerImpl{
		orderService: orderService,
	}
}

// func (controller *orderControllerImpl) Create(app echo.Context) error {
// 	var requestOrder models.CreateOrderRequest
// 	err := app.Bind(&requestOrder)
// 	if err != nil {
// 		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
// 	}

// 	currentUserToken, err := helpers.GetSplitedToken(app)
// 	if err != nil {
// 		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
// 	}

// 	response, err := controller.orderService.Create(app.Request().Context(), requestOrder, currentUserToken)
// 	if err != nil {
// 		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
// 	}

// 	return helpers.CreateResponse(app, http.StatusCreated, response)
// }

func (controller *orderControllerImpl) GetOrder(app echo.Context) error {
	var urlParam models.GetOrderByIDQueryRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancle()

	if urlParam.ID != uuid.Nil {
		response, err := controller.orderService.GetCompleteByID(ctx, urlParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	response, err := controller.orderService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

package controller

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userExternalControllerImpl struct {
	userExternalService services.UserExternalService
}

func NewUserExternalController(userExternalService services.UserExternalService) UserExternalController {
	return &userExternalControllerImpl{
		userExternalService: userExternalService,
	}
}

func (controller *userExternalControllerImpl) CreateOrder(app echo.Context) error {
	var requestOrder models.CreateOrderRequest
	err := app.Bind(&requestOrder)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	requestOrderCart := models.CreateOrderCartRequest{
		OrderID:            requestOrder.ID,
		ItemIDWithQuantity: make(map[uuid.UUID]int),
	}

	if len(requestOrder.ItemID) != len(requestOrder.Quantity) {
		errReqItemIDNotEqualToReqQuantity := errors.New("item id must be equal length to quantity")
		return helpers.CreateResponseError(app, http.StatusBadRequest, errReqItemIDNotEqualToReqQuantity)
	}

	for index, itemID := range requestOrder.ItemID {
		requestOrderCart.ItemIDWithQuantity[itemID] = requestOrder.Quantity[index]
	}

	requestTransferOrder := models.CreateTransferOrderRequest{
		OrderID: requestOrder.ID,
	}

	currentUser, err := helpers.GetSplitedToken(app)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	response, err := controller.userExternalService.CreateOrder(app.Request().Context(), requestOrder, requestOrderCart, requestTransferOrder, currentUser)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}

	return helpers.CreateResponse(app, http.StatusCreated, response)
}

// func (controller *userExternalControllerImpl) GetAllOrder(app echo.Context) error {
// 	ctx, cancle := context.WithTimeout(app.Request().Context(), 30*time.Second)
// 	defer cancle()

// 	currentUser, err := helpers.GetSplitedToken(app)
// 	if err != nil {
// 		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
// 	}

// 	responses, err := controller.userExternalService.GetAllOrder(ctx, currentUser)
// 	if err != nil {
// 		if err == context.DeadlineExceeded {
// 			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
// 		}
// 		return helpers.CreateResponseError(app, http.StatusNotFound, err)
// 	}
// 	return helpers.CreateResponse(app, http.StatusFound, responses)
// }

func (controller *userExternalControllerImpl) GetAllOrderByUser(app echo.Context) error {
	ctx, cancle := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancle()

	currentUser, err := helpers.GetSplitedToken(app)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	responses, err := controller.userExternalService.GetAllOrderCompleteByUserID(ctx, currentUser)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}

	return helpers.CreateResponse(app, http.StatusFound, responses)
}

func (controller *userExternalControllerImpl) GetItem(app echo.Context) error {
	var queryParam models.GetItemByNameCategoryRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancle()

	if queryParam.Name != "" {
		responses, err := controller.userExternalService.FindItemByName(ctx, queryParam.Name)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, responses)
	}

	if queryParam.Category != "" {
		responses, err := controller.userExternalService.FindItemByCategory(ctx, queryParam.Category)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, responses)
	}

	responses, err := controller.userExternalService.GetAllItem(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, responses)
}

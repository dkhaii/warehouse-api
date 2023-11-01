package controller

import (
	"errors"
	"net/http"

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
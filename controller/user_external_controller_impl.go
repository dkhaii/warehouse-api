package controller

import (
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

	var requestOrderCart models.CreateOrderCartRequest
	requestOrderCart.OrderID = requestOrder.ID

	requestOrderCart.ItemID = make([]uuid.UUID, len(requestOrder.ItemID))

	requestOrderCart.ItemID = append(requestOrderCart.ItemID, requestOrder.ItemID...)

	var requestTransferOrder models.CreateTransferOrderRequest
	requestTransferOrder.OrderID = requestOrder.ID

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

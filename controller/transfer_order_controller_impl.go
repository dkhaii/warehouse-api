package controller

// import (
// 	"net/http"

// 	"github.com/dkhaii/warehouse-api/helpers"
// 	"github.com/dkhaii/warehouse-api/models"
// 	"github.com/dkhaii/warehouse-api/services"
// 	"github.com/labstack/echo/v4"
// )

// type transferOrderControllerImpl struct {
// 	transferOrderService services.TransferOrderService
// 	orderService services.OrderService
// }

// func NewTransferOrderController(transferOrderService services.TransferOrderService, orderService services.OrderService) TransferOrderController {
// 	return &transferOrderControllerImpl{
// 		transferOrderService: transferOrderService,
// 		orderService: orderService,
// 	}
// }

// // func (controller *transferOrderControllerImpl) Create(app echo.Context) error {
// // 	var request models.CreateTransferOrderRequest
// // 	err := app.Bind(&request)
// // 	if err != nil {
// // 		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
// // 	}

// // 	currentOrder, err := controller.orderService.GetCompleteByID()
// // }
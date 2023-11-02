package services

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type userExternalServiceImpl struct {
	orderService         OrderService
	orderCartService     OrderCartService
	transferOrderService TransferOrderService
	database             *sql.DB
}

func NewUserExternalService(orderService OrderService, orderCartService OrderCartService, transferOrderService TransferOrderService, database *sql.DB) UserExternalService {
	return &userExternalServiceImpl{
		orderService:         orderService,
		orderCartService:     orderCartService,
		transferOrderService: transferOrderService,
		database:             database,
	}
}

func (service *userExternalServiceImpl) CreateOrder(ctx context.Context, requestOrder models.CreateOrderRequest, requestOrderCart models.CreateOrderCartRequest, requestTransferOrder models.CreateTransferOrderRequest, currentUserToken string) (models.CreateOrderResponse, error) {
	order, err := service.orderService.Create(ctx, requestOrder, currentUserToken)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	ocItemID := requestOrder.ItemID
	for index, itemID := range ocItemID {
		requestOrderCart.OrderID = order.ID
		requestOrderCart.ItemIDWithQuantity = map[uuid.UUID]int{itemID: requestOrder.Quantity[index]}

		err := service.orderCartService.Create(ctx, requestOrderCart)
		if err != nil {
			return models.CreateOrderResponse{}, err
		}
	}

	requestTransferOrder.OrderID = order.ID
	_, err = service.transferOrderService.Create(ctx, requestTransferOrder)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	return order, nil
}
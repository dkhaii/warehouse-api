package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/google/uuid"
)

type orderServiceImpl struct {
	orderRepository         repositories.OrderRepository
	orderCartRepository     repositories.OrderCartRepository
	transferOrderRepository repositories.TransferOrderRepository
	database                *sql.DB
}

func NewOrderService(orderRepository repositories.OrderRepository, orderCartRepository repositories.OrderCartRepository, transferOrderRepository repositories.TransferOrderRepository, database *sql.DB) OrderService {
	return &orderServiceImpl{
		orderRepository:         orderRepository,
		orderCartRepository:     orderCartRepository,
		transferOrderRepository: transferOrderRepository,
		database:                database,
	}
}

func (service *orderServiceImpl) Create(ctx context.Context, request models.CreateOrderRequest, currentUserToken string) (models.CreateOrderResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateOrderResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	config, err := config.Init()
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	currentUser, err := helpers.GetUserClaimsFromToken(currentUserToken, config.GetString("JWT_SECRET"))
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	uID := uuid.New()
	userID := currentUser.ID
	createdAt := time.Now()
	request.ID = uID
	request.UserID = userID
	request.CreatedAt = createdAt

	order := entity.Order{
		ID:                  request.ID,
		UserID:              request.UserID,
		Notes:               request.Notes,
		RequestTransferDate: request.RequestTransferDate,
		CreatedAt:           request.CreatedAt,
		User:                nil,
		Items:               nil,
	}

	createdOrder, err := service.orderRepository.Insert(ctx, tx, &order)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	// orderCart := make([]entity.OrderCart, len(request.ItemID))

	// for index, itemID := range request.ItemID {
	// 	orderCart[index] = entity.OrderCart{
	// 		ID:       uID,
	// 		OrderID:  createdOrder.ID,
	// 		ItemID:   itemID,
	// 		Quantity: request.Quantity,
	// 	}

	// 	_, err = service.orderCartRepository.Insert(ctx, tx, &orderCart[index])

	// 	if err != nil {
	// 		return models.CreateOrderResponse{}, nil
	// 	}
	// }

	// transferOrder := entity.TransferOrder{
	// 	ID:            uID,
	// 	OrderID:       createdOrder.ID,
	// 	UserID:        currentUser.ID,
	// 	Status:        "Pending",
	// 	FulfilledDate: time.Time{},
	// 	CreatedAt:     createdAt,
	// 	UpdatedAt:     createdAt,
	// 	Order:         nil,
	// }

	// _, err = service.transferOrderRepository.Insert(ctx, tx, &transferOrder)
	// if err != nil {
	// 	return models.CreateOrderResponse{}, err
	// }

	response := models.CreateOrderResponse{
		ID:                  createdOrder.ID,
		UserID:              createdOrder.UserID,
		Notes:               createdOrder.Notes,
		RequestTransferDate: createdOrder.RequestTransferDate,
		CreatedAt:           createdOrder.CreatedAt,
	}

	return response, nil
}

func (service *orderServiceImpl) GetAll(ctx context.Context) ([]models.GetOrderResponse, error) {
	rows, err := service.orderRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	orders := make([]models.GetOrderResponse, len(rows))

	for key, order := range rows {
		orders[key] = models.GetOrderResponse{
			ID:                  order.ID,
			UserID:              order.UserID,
			Notes:               order.Notes,
			RequestTransferDate: order.RequestTransferDate,
			CreatedAt:           order.CreatedAt,
		}
	}

	return orders, nil
}

func (service *orderServiceImpl) GetCompleteByID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteOrderResponse, error) {
	order, err := service.orderRepository.FindCompleteByID(ctx, ordID)
	if err != nil {
		return models.GetCompleteOrderResponse{}, err
	}

	response := models.GetCompleteOrderResponse{
		ID:                  order.ID,
		UserID:              order.UserID,
		Notes:               order.Notes,
		RequestTransferDate: order.RequestTransferDate,
		CreatedAt:           order.CreatedAt,
		User:                order.User,
		Items:               order.Items,
	}

	return response, nil
}

package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/google/uuid"
)

type orderServiceImpl struct {
	orderRepository repositories.OrderRepository
	database        *sql.DB
}

func NewOrderService(orderRepository repositories.OrderRepository, database *sql.DB) OrderService {
	return &orderServiceImpl{
		orderRepository: orderRepository,
		database:        database,
	}
}

func (service *orderServiceImpl) Create(ctx context.Context, request models.CreateOrderRequest) (models.CreateOrderResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateOrderResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	orderID := uuid.New()
	createdAt := time.Now()
	request.ID = orderID
	request.CreatedAt = createdAt

	order := entity.Order{
		ID:                  request.ID,
		ItemID:              request.ItemID,
		UserID:              request.UserID,
		Quantity:            request.Quantity,
		Notes:               request.Notes,
		RequestTransferDate: request.RequestTransferDate,
		CreatedAt:           request.CreatedAt,
		User:                nil,
		Item:                nil,
	}

	createdOrder, err := service.orderRepository.Insert(ctx, tx, &order)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	response := models.CreateOrderResponse{
		ID:                  createdOrder.ID,
		ItemID:              createdOrder.ItemID,
		UserID:              createdOrder.UserID,
		Quantity:            createdOrder.Quantity,
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
			ItemID:              order.ItemID,
			UserID:              order.UserID,
			Quantity:            order.Quantity,
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
		ID: order.ID,
		ItemID: order.ItemID,
		UserID: order.UserID,
		Quantity: order.Quantity,
		Notes: order.Notes,
		RequestTransferDate: order.RequestTransferDate,
		CreatedAt: order.CreatedAt,
		User: order.User,
		Item: order.Item,
	}

	return response, nil
}
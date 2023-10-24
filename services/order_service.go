package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type OrderService interface {
	Create(ctx context.Context, request models.CreateOrderRequest, currentUserToken string) (models.CreateOrderResponse, error)
	GetAll(ctx context.Context) ([]models.GetOrderResponse, error)
	GetCompleteByID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteOrderResponse, error)
}

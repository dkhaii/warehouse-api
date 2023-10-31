package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
)

type UserExternalService interface {
	CreateOrder(ctx context.Context, requestOrder models.CreateOrderRequest, requestOrderCart models.CreateOrderCartRequest, request models.CreateTransferOrderRequest, currentUserToken string) (models.CreateOrderResponse, error)
}
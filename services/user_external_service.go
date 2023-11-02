package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
)

type UserExternalService interface {
	CreateOrder(ctx context.Context, requestOrder models.CreateOrderRequest, requestOrderCart models.CreateOrderCartRequest, request models.CreateTransferOrderRequest, currentUserToken string) (models.CreateOrderResponse, error)
	GetAllOrder(ctx context.Context, currentUserToken string) ([]models.GetOrderResponse, error)
	GetAllItem(ctx context.Context) ([]models.GetItemFilteredResponse, error)
	FindItemByName(ctx context.Context, itmName string) ([]models.GetItemFilteredResponse, error)
	FindItemByCategory(ctx context.Context, ctgName string) ([]models.GetItemFilteredResponse, error)
}

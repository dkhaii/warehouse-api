package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type TransferOrderService interface {
	Create(ctx context.Context, requestTrfOrd models.CreateTransferOrderRequest, requestOrder models.GetOrderByIDQueryRequest) (models.CreateTransferOrderResponse, error)
	FindByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetTransferOrderResponse, error)
	FindCompleteByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetCompleteTransferOrderResponse, error)
	Update(ctx context.Context, request models.UpdateTransferOrderRequest, currentUserToken string) (models.CreateTransferOrderResponse, error)
}
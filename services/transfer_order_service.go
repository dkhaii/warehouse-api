package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type TransferOrderService interface {
	Create(ctx context.Context, request models.CreateTransferOrderRequest) (models.CreateTransferOrderResponse, error)
	FindAll(ctx context.Context) ([]models.GetTransferOrderResponse, error)
	FindByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetTransferOrderResponse, error)
	FindCompleteByOrderID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteTransferOrderResponse, error)
	Update(ctx context.Context, request models.UpdateTransferOrderRequest, currentUserToken string) (models.CreateTransferOrderResponse, error)
}
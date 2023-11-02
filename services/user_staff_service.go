package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type UserStaffService interface {
	CreateItem(ctx context.Context, request models.CreateItemRequest, currentUserToken string) (models.CreateItemResponse, error)
	GetAllItem(ctx context.Context) ([]models.GetItemResponse, error)
	GetItemByID(ctx context.Context, itmID uuid.UUID) (models.GetItemResponse, error)
	GetItemByName(ctx context.Context, itmName string) ([]models.GetItemResponse, error)
	GetItemByCategoryName(ctx context.Context, ctgName string) ([]models.GetItemResponse, error)
	GetCompleteItemByID(ctx context.Context, itmID uuid.UUID) (models.GetCompleteItemResponse, error)
	UpdateItem(ctx context.Context, request models.UpdateItemRequest, currentUserToken string) (models.CreateItemResponse, error)
	DeleteItem(ctx context.Context, itmID uuid.UUID) error
	GetAllLocation(ctx context.Context) ([]models.GetLocationResponse, error)
	GetCompleteLocationByID(ctx context.Context, locID string) (models.GetCompleteLocationResponse, error)
	GetAllCategory(ctx context.Context) ([]models.GetCategoryResponse, error)
	GetCategoryByID(ctx context.Context, ctgID string) (models.GetCategoryResponse, error)
	GetCategoryByName(ctx context.Context, ctgName string) ([]models.GetCategoryResponse, error)
	GetAllTransferOrder(ctx context.Context) ([]models.GetTransferOrderResponse, error)
	GetTransferOrderByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetTransferOrderResponse, error)
	GetCompleteTransferOrderByOrderID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteTransferOrderResponse, error)
	UpdateTransferOrder(ctx context.Context, request models.UpdateTransferOrderRequest, currentUserToken string) (models.CreateTransferOrderResponse, error)
}

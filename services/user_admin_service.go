package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type UserAdminService interface {
	CreateCategory(ctx context.Context, request models.CreateCategoryRequest) (models.CreateCategoryResponse, error)
	GetAllCategory(ctx context.Context) ([]models.GetCategoryResponse, error)
	GetCategoryByID(ctx context.Context, ctgID string) (models.GetCategoryResponse, error)
	GetCategoryByName(ctx context.Context, ctgName string) ([]models.GetCategoryResponse, error)
	UpdateCategory(ctx context.Context, request models.UpdateCategoryRequest) (models.CreateCategoryResponse, error)
	DeleteCategory(ctx context.Context, ctgID string) error
	CreateLocation(ctx context.Context, request models.CreateLocationRequest) (models.CreateLocationResponse, error)
	GetAllLocation(ctx context.Context) ([]models.GetLocationResponse, error)
	GetCompleteLocationByID(ctx context.Context, locID string) (models.GetCompleteLocationResponse, error)
	UpdateLocation(ctx context.Context, request models.UpdateLocationRequest) (models.CreateLocationResponse, error)
	DeleteLocation(ctx context.Context, locID string) error
	CreateItem(ctx context.Context, request models.CreateItemRequest, currentUserToken string) (models.CreateItemResponse, error)
	GetAllItem(ctx context.Context) ([]models.GetItemResponse, error)
	GetItemByID(ctx context.Context, itmID uuid.UUID) (models.GetItemResponse, error)
	GetItemByName(ctx context.Context, name string) ([]models.GetItemResponse, error)
	GetItemByCategoryName(ctx context.Context, ctgName string) ([]models.GetItemResponse, error)
	GetCompleteItemByID(ctx context.Context, itmID uuid.UUID) (models.GetCompleteItemResponse, error)
	UpdateItem(ctx context.Context, request models.UpdateItemRequest, currentUserToken string) (models.CreateItemResponse, error)
	DeleteItem(ctx context.Context, itmID uuid.UUID) error
	GetAllOrder(ctx context.Context) ([]models.GetOrderResponse, error)
	GetCompleteOrderByID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteOrderResponse, error)
	GetAllTransferOrder(ctx context.Context) ([]models.GetTransferOrderResponse, error)
	GetByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetTransferOrderResponse, error)
	GetCompleteByOrderID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteTransferOrderResponse, error)

}

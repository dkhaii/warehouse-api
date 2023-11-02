package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type userStaffServiceImpl struct {
	itemService          ItemService
	locationService      LocationService
	categoryService      CategoryService
	transferOrderService TransferOrderService
}

func NewUserStaffService(itemService ItemService, locationService LocationService, categoryService CategoryService, transferOrderService TransferOrderService) UserStaffService {
	return &userStaffServiceImpl{
		itemService:          itemService,
		locationService:      locationService,
		categoryService:      categoryService,
		transferOrderService: transferOrderService,
	}
}

func (service *userStaffServiceImpl) CreateItem(ctx context.Context, request models.CreateItemRequest, currentUserToken string) (models.CreateItemResponse, error) {
	item, err := service.itemService.Create(ctx, request, currentUserToken)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	return item, nil
}

func (service *userStaffServiceImpl) GetAllItem(ctx context.Context) ([]models.GetItemResponse, error) {
	rows, err := service.itemService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetItemByID(ctx context.Context, itmID uuid.UUID) (models.GetItemResponse, error) {
	item, err := service.itemService.GetByID(ctx, itmID)
	if err != nil {
		return models.GetItemResponse{}, err
	}

	return item, nil
}

func (service *userStaffServiceImpl) GetItemByName(ctx context.Context, itmName string) ([]models.GetItemResponse, error) {
	rows, err := service.itemService.GetByName(ctx, itmName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetItemByCategoryName(ctx context.Context, ctgName string) ([]models.GetItemResponse, error) {
	rows, err := service.itemService.GetByCategoryName(ctx, ctgName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetCompleteItemByID(ctx context.Context, itmID uuid.UUID) (models.GetCompleteItemResponse, error) {
	item, err := service.itemService.GetCompleteByID(ctx, itmID)
	if err != nil {
		return models.GetCompleteItemResponse{}, err
	}

	return item, nil
}

func (service *userStaffServiceImpl) UpdateItem(ctx context.Context, request models.UpdateItemRequest, currentUserToken string) (models.CreateItemResponse, error) {
	item, err := service.itemService.Update(ctx, request, currentUserToken)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	return item, nil
}

func (service *userStaffServiceImpl) DeleteItem(ctx context.Context, itmID uuid.UUID) error {
	err := service.itemService.Delete(ctx, itmID)
	if err != nil {
		return err
	}

	return nil
}

func (service *userStaffServiceImpl) GetAllLocation(ctx context.Context) ([]models.GetLocationResponse, error) {
	rows, err := service.locationService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetCompleteLocationByID(ctx context.Context, locID string) (models.GetCompleteLocationResponse, error) {
	location, err := service.locationService.GetCompleteByID(ctx, locID)
	if err != nil {
		return models.GetCompleteLocationResponse{}, err
	}

	return location, nil
}

func (service *userStaffServiceImpl) GetAllCategory(ctx context.Context) ([]models.GetCategoryResponse, error) {
	rows, err := service.categoryService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetCategoryByID(ctx context.Context, ctgID string) (models.GetCategoryResponse, error) {
	category, err := service.categoryService.GetByID(ctx, ctgID)
	if err != nil {
		return models.GetCategoryResponse{}, err
	}

	return category, nil
}

func (service *userStaffServiceImpl) GetCategoryByName(ctx context.Context, ctgName string) ([]models.GetCategoryResponse, error) {
	rows, err := service.categoryService.GetByName(ctx, ctgName)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetAllTransferOrder(ctx context.Context) ([]models.GetTransferOrderResponse, error) {
	rows, err := service.transferOrderService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userStaffServiceImpl) GetTransferOrderByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetTransferOrderResponse, error) {
	transferOrder, err := service.transferOrderService.GetByID(ctx, trfOrdID)
	if err != nil {
		return models.GetTransferOrderResponse{}, err
	}

	return transferOrder, nil
}

func (service *userStaffServiceImpl) GetCompleteTransferOrderByOrderID(ctx context.Context, ordID uuid.UUID) (models.GetCompleteTransferOrderResponse, error) {
	transferOrder, err := service.transferOrderService.GetCompleteByOrderID(ctx, ordID)
	if err != nil {
		return models.GetCompleteTransferOrderResponse{}, err
	}

	return transferOrder, nil
}

func (service *userStaffServiceImpl) UpdateTransferOrder(ctx context.Context, request models.UpdateTransferOrderRequest, currentUserToken string) (models.CreateTransferOrderResponse, error) {
	transferOrder, err := service.transferOrderService.Update(ctx, request, currentUserToken)
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	return transferOrder, err
}

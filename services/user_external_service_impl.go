package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
	"github.com/google/uuid"
)

type userExternalServiceImpl struct {
	orderService         OrderService
	orderCartService     OrderCartService
	transferOrderService TransferOrderService
	itemService          ItemService
}

func NewUserExternalService(orderService OrderService, orderCartService OrderCartService, transferOrderService TransferOrderService, itemService ItemService) UserExternalService {
	return &userExternalServiceImpl{
		orderService:         orderService,
		orderCartService:     orderCartService,
		transferOrderService: transferOrderService,
		itemService:          itemService,
	}
}

func (service *userExternalServiceImpl) CreateOrder(ctx context.Context, requestOrder models.CreateOrderRequest, requestOrderCart models.CreateOrderCartRequest, requestTransferOrder models.CreateTransferOrderRequest, currentUserToken string) (models.CreateOrderResponse, error) {
	order, err := service.orderService.Create(ctx, requestOrder, currentUserToken)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	ocItemID := requestOrder.ItemID
	for index, itemID := range ocItemID {
		requestOrderCart.OrderID = order.ID
		requestOrderCart.ItemIDWithQuantity = map[uuid.UUID]int{itemID: requestOrder.Quantity[index]}

		err := service.orderCartService.Create(ctx, requestOrderCart)
		if err != nil {
			return models.CreateOrderResponse{}, err
		}
	}

	requestTransferOrder.OrderID = order.ID
	_, err = service.transferOrderService.Create(ctx, requestTransferOrder)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	return order, nil
}

func (service *userExternalServiceImpl) GetAllOrder(ctx context.Context, currentUserToken string) ([]models.GetOrderResponse, error) {
	rows, err := service.orderService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (service *userExternalServiceImpl) GetAllItem(ctx context.Context) ([]models.GetItemFilteredResponse, error) {
	rows, err := service.itemService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemFilteredResponse, len(rows))

	for index, item := range rows {
		responses[index] = models.GetItemFilteredResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			CategoryID:   item.CategoryID,
		}
	}

	return responses, nil
}

func (service *userExternalServiceImpl) FindItemByName(ctx context.Context, itmName string) ([]models.GetItemFilteredResponse, error) {
	rows, err := service.itemService.GetByName(ctx, itmName)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemFilteredResponse, len(rows))

	for index, item := range rows {
		responses[index] = models.GetItemFilteredResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			CategoryID:   item.CategoryID,
		}
	}

	return responses, nil
}

func (service *userExternalServiceImpl) FindItemByCategory(ctx context.Context, ctgName string) ([]models.GetItemFilteredResponse, error) {
	rows, err := service.itemService.GetByCategoryName(ctx, ctgName)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemFilteredResponse, len(rows))

	for index, item := range rows {
		responses[index] = models.GetItemFilteredResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			CategoryID:   item.CategoryID,
		}
	}

	return responses, nil
}

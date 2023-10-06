package service

import (
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/model"
	"github.com/dkhaii/warehouse-api/repository"
	"github.com/google/uuid"
)

type itemServiceImpl struct {
	itemRepository repository.ItemRepository
}

func NewItemService(itemRepository repository.ItemRepository) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (service *itemServiceImpl) Create(request model.CreateItemRequest) (model.CreateItemResponse, error) {
	item := entity.Item{
		ID:           request.ID,
		Name:         request.Name,
		Description:  request.Description,
		Quantity:     request.Quantity,
		Availability: request.Availability,
		LocationID:   request.LocationID,
		CategoryID:   request.CategoryID,
		UserID:       request.UserID,
		CreatedAt:    request.CreatedAt,
		UpdatedAt:    request.UpdatedAt,
	}

	_, err := service.itemRepository.Insert(&item)
	if err != nil {
		return model.CreateItemResponse{}, err
	}

	response := model.CreateItemResponse{
		ID:           item.ID,
		Name:         item.Name,
		Description:  item.Description,
		Quantity:     item.Quantity,
		Availability: item.Availability,
		LocationID:   item.LocationID,
		CategoryID:   item.CategoryID,
		UserID:       item.UserID,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	return response, nil
}

func (service *itemServiceImpl) GetAll() ([]model.GetItemResponse, error) {
	items, err := service.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]model.GetItemResponse, len(items))

	for key, item := range items {
		responses[key] = model.GetItemResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			LocationID:   item.LocationID,
			CategoryID:   item.CategoryID,
			UserID:       item.UserID,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *itemServiceImpl) GetByID(itmID uuid.UUID) (model.GetItemResponse, error) {
	item, err := service.itemRepository.FindByID(itmID)
	if err != nil {
		return model.GetItemResponse{}, err
	}

	response := model.GetItemResponse{
		ID:           item.ID,
		Name:         item.Name,
		Description:  item.Description,
		Quantity:     item.Quantity,
		Availability: item.Availability,
		LocationID:   item.LocationID,
		CategoryID:   item.CategoryID,
		UserID:       item.UserID,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	return response, nil
}

func (service *itemServiceImpl) GetByName(name string) ([]model.GetItemResponse, error) {
	items, err := service.itemRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	responses := make([]model.GetItemResponse, len(items))

	for key, item := range items {
		responses[key] = model.GetItemResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			LocationID:   item.LocationID,
			CategoryID:   item.CategoryID,
			UserID:       item.UserID,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *itemServiceImpl) Update(request model.CreateItemRequest) (model.CreateItemResponse, error) {
	isItem, err := service.itemRepository.FindByID(request.ID)
	if err != nil {
		return model.CreateItemResponse{}, err
	}

	updatedItem := entity.Item{
		ID:           isItem.ID,
		Name:         request.Name,
		Description:  request.Description,
		Quantity:     request.Quantity,
		Availability: request.Availability,
		LocationID:   request.LocationID,
		CategoryID:   request.CategoryID,
		UserID:       request.UserID,
		CreatedAt:    isItem.CreatedAt,
		UpdatedAt:    request.UpdatedAt,
	}

	err = service.itemRepository.Update(&updatedItem)
	if err != nil {
		return model.CreateItemResponse{}, err
	}

	response := model.CreateItemResponse{
		ID:           updatedItem.ID,
		Name:         updatedItem.Name,
		Description:  updatedItem.Description,
		Quantity:     updatedItem.Quantity,
		Availability: updatedItem.Availability,
		LocationID:   updatedItem.LocationID,
		CategoryID:   updatedItem.CategoryID,
		UserID:       updatedItem.UserID,
		CreatedAt:    updatedItem.CreatedAt,
		UpdatedAt:    updatedItem.UpdatedAt,
	}

	return response, nil
}

func (service *itemServiceImpl) Delete(itmID uuid.UUID) error {
	item, err := service.itemRepository.FindByID(itmID)
	if err != nil {
		return err
	}

	err = service.itemRepository.Delete(item.ID)
	if err != nil {
		return err
	}

	return nil
}

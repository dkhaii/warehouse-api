package services

import (
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/google/uuid"
)

type itemServiceImpl struct {
	itemRepository repositories.ItemRepository
}

func NewItemService(itemRepository repositories.ItemRepository) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
	}
}

func (service *itemServiceImpl) Create(request models.CreateItemRequest) (models.CreateItemResponse, error) {
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
		return models.CreateItemResponse{}, err
	}

	response := models.CreateItemResponse{
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

func (service *itemServiceImpl) GetAll() ([]models.GetItemResponse, error) {
	items, err := service.itemRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemResponse, len(items))

	for key, item := range items {
		responses[key] = models.GetItemResponse{
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

func (service *itemServiceImpl) GetByID(itmID uuid.UUID) (models.GetItemResponse, error) {
	item, err := service.itemRepository.FindByID(itmID)
	if err != nil {
		return models.GetItemResponse{}, err
	}

	response := models.GetItemResponse{
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

func (service *itemServiceImpl) GetByName(name string) ([]models.GetItemResponse, error) {
	items, err := service.itemRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemResponse, len(items))

	for key, item := range items {
		responses[key] = models.GetItemResponse{
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

func (service *itemServiceImpl) Update(request models.CreateItemRequest) (models.CreateItemResponse, error) {
	isItem, err := service.itemRepository.FindByID(request.ID)
	if err != nil {
		return models.CreateItemResponse{}, err
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
		return models.CreateItemResponse{}, err
	}

	response := models.CreateItemResponse{
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

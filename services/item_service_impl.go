package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/google/uuid"
)

type itemServiceImpl struct {
	itemRepository repositories.ItemRepository
	database       *sql.DB
}

func NewItemService(itemRepository repositories.ItemRepository, database *sql.DB) ItemService {
	return &itemServiceImpl{
		itemRepository: itemRepository,
		database:       database,
	}
}

func (service *itemServiceImpl) Create(ctx context.Context, request models.CreateItemRequest, currentUserToken string) (models.CreateItemResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateItemResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	config, err := config.Init()
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	currentUser, err := helpers.GetUserClaimsFromToken(currentUserToken, config.GetString("JWT_SECRET"))
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	itemID := uuid.New()
	createdAt := time.Now()
	userID := currentUser.ID
	request.ID = itemID
	request.UserID = userID
	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	item := entity.Item{
		ID:           request.ID,
		Name:         request.Name,
		Description:  request.Description,
		Quantity:     request.Quantity,
		Availability: request.Availability,
		CategoryID:   request.CategoryID,
		UserID:       request.UserID,
		CreatedAt:    request.CreatedAt,
		UpdatedAt:    request.UpdatedAt,
		Category:     nil,
		User:         nil,
		Location:     nil,
	}

	createdItem, err := service.itemRepository.Insert(ctx, tx, &item)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	response := models.CreateItemResponse{
		ID:           createdItem.ID,
		Name:         createdItem.Name,
		Description:  createdItem.Description,
		Quantity:     createdItem.Quantity,
		Availability: createdItem.Availability,
		CategoryID:   createdItem.CategoryID,
		UserID:       createdItem.UserID,
		CreatedAt:    createdItem.CreatedAt,
		UpdatedAt:    createdItem.UpdatedAt,
	}

	return response, nil
}

func (service *itemServiceImpl) GetAll(ctx context.Context) ([]models.GetItemResponse, error) {
	rows, err := service.itemRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemResponse, len(rows))

	for key, item := range rows {
		responses[key] = models.GetItemResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			CategoryID:   item.CategoryID,
			UserID:       item.UserID,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *itemServiceImpl) GetByID(ctx context.Context, itmID uuid.UUID) (models.GetItemResponse, error) {
	item, err := service.itemRepository.FindByID(ctx, itmID)
	if err != nil {
		return models.GetItemResponse{}, err
	}

	response := models.GetItemResponse{
		ID:           item.ID,
		Name:         item.Name,
		Description:  item.Description,
		Quantity:     item.Quantity,
		Availability: item.Availability,
		CategoryID:   item.CategoryID,
		UserID:       item.UserID,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	return response, nil
}

func (service *itemServiceImpl) GetByName(ctx context.Context, name string) ([]models.GetItemResponse, error) {
	rows, err := service.itemRepository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemResponse, len(rows))

	for index, item := range rows {
		responses[index] = models.GetItemResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			CategoryID:   item.CategoryID,
			UserID:       item.UserID,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *itemServiceImpl) GetCompleteByID(ctx context.Context, itmID uuid.UUID) (models.GetCompleteItemResponse, error) {
	item, err := service.itemRepository.FindCompleteByID(ctx, itmID)
	if err != nil {
		return models.GetCompleteItemResponse{}, err
	}

	response := models.GetCompleteItemResponse{
		ID:           item.ID,
		Name:         item.Name,
		Description:  item.Description,
		Quantity:     item.Quantity,
		Availability: item.Availability,
		CategoryID:   item.CategoryID,
		UserID:       item.UserID,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
		Category:     item.Category,
		User:         item.User,
		Location:     item.Location,
	}

	return response, nil
}

func (service *itemServiceImpl) GetByCategoryName(ctx context.Context, ctgName string) ([]models.GetItemResponse, error) {
	rows, err := service.itemRepository.FindByCategoryName(ctx, ctgName)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetItemResponse, len(rows))

	for index, item := range rows {
		responses[index] = models.GetItemResponse{
			ID:           item.ID,
			Name:         item.Name,
			Description:  item.Description,
			Quantity:     item.Quantity,
			Availability: item.Availability,
			CategoryID:   item.CategoryID,
			UserID:       item.UserID,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *itemServiceImpl) Update(ctx context.Context, request models.UpdateItemRequest, currentUserToken string) (models.CreateItemResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	item, err := service.itemRepository.FindByID(ctx, request.ID)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateItemResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	config, err := config.Init()
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	currentUser, err := helpers.GetUserClaimsFromToken(currentUserToken, config.GetString("JWT_SECRET"))
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	updatedAt := time.Now()
	request.UpdatedAt = updatedAt
	request.UserID = currentUser.ID

	updatedItem := entity.Item{
		ID:           item.ID,
		Name:         request.Name,
		Description:  request.Description,
		Quantity:     request.Quantity,
		Availability: request.Availability,
		CategoryID:   request.CategoryID,
		UserID:       request.UserID,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    request.UpdatedAt,
		Category:     nil,
		User:         nil,
		Location:     nil,
	}

	err = service.itemRepository.Update(ctx, tx, &updatedItem)
	if err != nil {
		return models.CreateItemResponse{}, err
	}

	response := models.CreateItemResponse{
		ID:           updatedItem.ID,
		Name:         updatedItem.Name,
		Description:  updatedItem.Description,
		Quantity:     updatedItem.Quantity,
		Availability: updatedItem.Availability,
		CategoryID:   updatedItem.CategoryID,
		UserID:       updatedItem.UserID,
		CreatedAt:    updatedItem.CreatedAt,
		UpdatedAt:    updatedItem.UpdatedAt,
	}

	return response, nil
}

func (service *itemServiceImpl) Delete(ctx context.Context, itmID uuid.UUID) error {
	item, err := service.itemRepository.FindByID(ctx, itmID)
	if err != nil {
		return err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return err
	}
	defer helpers.CommitOrRollBack(tx)

	err = service.itemRepository.Delete(ctx, tx, item.ID)
	if err != nil {
		return err
	}

	return nil
}

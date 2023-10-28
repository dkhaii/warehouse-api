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

type transferOrderServiceImpl struct {
	transferOrderRepository repositories.TransferOrderRepository
	database                *sql.DB
}

func NewTransferOrderService(transferOrderRepository repositories.TransferOrderRepository, database *sql.DB) TransferOrderService {
	return &transferOrderServiceImpl{
		transferOrderRepository: transferOrderRepository,
		database:                database,
	}
}

func (service *transferOrderServiceImpl) Create(ctx context.Context, requestTrfOrd models.CreateTransferOrderRequest, requestOrder models.GetOrderByIDQueryRequest) (models.CreateTransferOrderResponse, error) {
	err := helpers.ValidateRequest(requestTrfOrd)
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	toID := uuid.New()
	userID := uuid.Nil
	status := "Pending"
	CreatedAt := time.Now()
	requestTrfOrd.ID = toID
	requestTrfOrd.UserID = userID
	requestTrfOrd.Status = status
	requestTrfOrd.CreatedAt = CreatedAt
	requestTrfOrd.UpdatedAt = requestTrfOrd.CreatedAt

	transferOrder := entity.TransferOrder{
		ID:            requestTrfOrd.ID,
		OrderID:       requestOrder.ID,
		UserID:        requestTrfOrd.UserID,
		Status:        requestTrfOrd.Status,
		FulfilledDate: requestTrfOrd.FulfilledDate,
		CreatedAt:     requestTrfOrd.CreatedAt,
		UpdatedAt:     requestTrfOrd.UpdatedAt,
	}

	createdTrfOrd, err := service.transferOrderRepository.Insert(ctx, tx, &transferOrder)
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	response := models.CreateTransferOrderResponse{
		ID:            createdTrfOrd.ID,
		OrderID:       createdTrfOrd.ID,
		UserID:        createdTrfOrd.UserID,
		Status:        createdTrfOrd.Status,
		FulfilledDate: createdTrfOrd.FulfilledDate,
		CreatedAt:     createdTrfOrd.CreatedAt,
		UpdatedAt:     createdTrfOrd.UpdatedAt,
	}

	return response, nil
}

func (service *transferOrderServiceImpl) FindAll(ctx context.Context) ([]models.GetTransferOrderResponse, error) {
	rows, err := service.transferOrderRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetTransferOrderResponse, len(rows))

	for key, to := range rows {
		responses[key] = models.GetTransferOrderResponse{
			ID:            to.ID,
			OrderID:       to.OrderID,
			UserID:        to.UserID,
			Status:        to.Status,
			FulfilledDate: to.FulfilledDate,
			CreatedAt:     to.CreatedAt,
			UpdatedAt:     to.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *transferOrderServiceImpl) FindByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetTransferOrderResponse, error) {
	to, err := service.transferOrderRepository.FindByID(ctx, trfOrdID)
	if err != nil {
		return models.GetTransferOrderResponse{}, err
	}

	response := models.GetTransferOrderResponse{
		ID:            to.ID,
		OrderID:       to.OrderID,
		UserID:        to.UserID,
		Status:        to.Status,
		FulfilledDate: to.FulfilledDate,
		CreatedAt:     to.CreatedAt,
		UpdatedAt:     to.UpdatedAt,
	}

	return response, nil
}

func (service *transferOrderServiceImpl) FindCompleteByID(ctx context.Context, trfOrdID uuid.UUID) (models.GetCompleteTransferOrderResponse, error) {
	to, err := service.transferOrderRepository.FindCompleteByID(ctx, trfOrdID)
	if err != nil {
		return models.GetCompleteTransferOrderResponse{}, nil
	}

	response := models.GetCompleteTransferOrderResponse{
		ID:            to.ID,
		OrderID:       to.OrderID,
		UserID:        to.UserID,
		Status:        to.Status,
		FulfilledDate: to.FulfilledDate,
		CreatedAt:     to.CreatedAt,
		UpdatedAt:     to.UpdatedAt,
		Order:         to.Order,
	}

	return response, nil
}

func (service *transferOrderServiceImpl) Update(ctx context.Context, request models.UpdateTransferOrderRequest, currentUserToken string) (models.CreateTransferOrderResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	to, err := service.transferOrderRepository.FindByID(ctx, request.ID)
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	config, err := config.Init()
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	currentUser, err := helpers.GetUserClaimsFromToken(currentUserToken, config.GetString("JWT_SECRET"))
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	updatedAt := time.Now()
	request.ID = to.ID
	request.OrderID = to.OrderID
	request.UserID = currentUser.ID
	request.FulfilledDate = time.Time{}
	request.CreatedAt = to.CreatedAt
	request.UpdatedAt = updatedAt

	if request.Status == "Finished" {
		request.FulfilledDate = time.Now()
	}

	updatedTrfOrd := entity.TransferOrder{
		ID:            request.ID,
		OrderID:       request.OrderID,
		UserID:        request.UserID,
		Status:        request.Status,
		FulfilledDate: request.FulfilledDate,
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     request.UpdatedAt,
		Order:         nil,
	}

	createdTrfOrd, err := service.transferOrderRepository.Update(ctx, tx, &updatedTrfOrd)
	if err != nil {
		return models.CreateTransferOrderResponse{}, err
	}

	response := models.CreateTransferOrderResponse{
		ID:            createdTrfOrd.ID,
		OrderID:       createdTrfOrd.OrderID,
		UserID:        createdTrfOrd.UserID,
		Status:        createdTrfOrd.Status,
		FulfilledDate: createdTrfOrd.FulfilledDate,
		CreatedAt:     createdTrfOrd.CreatedAt,
		UpdatedAt:     createdTrfOrd.UpdatedAt,
	}

	return response, nil
}

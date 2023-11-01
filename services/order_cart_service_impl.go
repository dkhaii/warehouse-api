package services

import (
	"context"
	"database/sql"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/google/uuid"
)

type orderCartServiceImpl struct {
	orderCartRepository repositories.OrderCartRepository
	database            *sql.DB
}

func NewOrderCartService(orderCartRepository repositories.OrderCartRepository, database *sql.DB) OrderCartService {
	return &orderCartServiceImpl{
		orderCartRepository: orderCartRepository,
		database:            database,
	}
}

func (service *orderCartServiceImpl) Create(ctx context.Context, request models.CreateOrderCartRequest) error {
	tx, err := service.database.Begin()
	if err != nil {
		return err
	}
	defer helpers.CommitOrRollBack(tx)

	orderCartID := uuid.New()
	request.ID = orderCartID

	var listOfOrderCart []entity.OrderCart

	for itemID, quantity := range request.ItemIDWithQuantity {
		orderCart := entity.OrderCart{
			ID:       request.ID,
			OrderID:  request.OrderID,
			ItemID:   itemID,
			Quantity: quantity,
		}

		listOfOrderCart = append(listOfOrderCart, orderCart)
	}

	for _, orderCart := range listOfOrderCart {
		_, err := service.orderCartRepository.Insert(ctx, tx, &orderCart)
		if err != nil {
			return err
		}
	}

	return nil
}
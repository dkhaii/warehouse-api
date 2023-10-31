package services

import (
	"context"

	"github.com/dkhaii/warehouse-api/models"
)

type OrderCartService interface {
	Create(ctx context.Context, request models.CreateOrderCartRequest) error
}
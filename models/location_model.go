package models

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
)

type CreateLocationRequest struct {
	ID          string    `json:"id" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateLocationResponse struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetLocationResponse struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetLocationByIDQueryRequest struct {
	ID string `query:"id"`
}

type GetLocationByIDParamRequest struct {
	ID string `param:"id"`
}

type GetCompleteLocationResponse struct {
	ID          string                    `json:"id"`
	Description string                    `json:"description"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	Category    []entity.CategoryFiltered `json:"category"`
}

type UpdateLocationRequest struct {
	ID          string    `param:"id"`
	Description string    `json:"description" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
}

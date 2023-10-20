package models

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
)

type CreateCategoryRequest struct {
	ID          string    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	LocationID  string    `json:"location_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCategoryResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LocationID  string    `json:"location_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetCategoryIDRequest struct {
	ID string `param:"id"`
}

type GetCategoryRequest struct {
	ID   string `query:"id"`
	Name string `query:"name"`
}

type GetCategoryResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LocationID  string    `json:"location_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetCompleteCategoryByIDResponse struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	LocationID  string           `json:"location_id"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Location    *entity.Location `json:"location"`
}

type UpdateCategoryRequest struct {
	ID          string    `param:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	LocationID  string    `json:"location_id" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at"`
}

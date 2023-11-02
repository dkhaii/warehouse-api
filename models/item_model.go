package models

import (
	"time"

	"github.com/dkhaii/warehouse-api/entity"
	"github.com/google/uuid"
)

type CreateItemRequest struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Quantity     int       `json:"quantity" validate:"required"`
	Availability bool      `json:"availability" validate:"required"`
	CategoryID   string    `json:"category_id" validate:"required"`
	UserID       uuid.UUID `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateItemResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	CategoryID   string    `json:"category_id"`
	UserID       uuid.UUID `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetItemRequest struct {
	ID   uuid.UUID `query:"id"`
	Name string    `query:"name"`
}

type GetItemByIDParamRequest struct {
	ID uuid.UUID `param:"id"`
}

type GetItemResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	CategoryID   string    `json:"category_id"`
	UserID       uuid.UUID `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetItemFilteredResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Quantity     int       `json:"quantity"`
	Availability bool      `json:"availability"`
	CategoryID   string    `json:"category_id"`
}

type GetCompleteItemResponse struct {
	ID           uuid.UUID                `json:"id"`
	Name         string                   `json:"name"`
	Description  string                   `json:"description"`
	Quantity     int                      `json:"quantity"`
	Availability bool                     `json:"availability"`
	CategoryID   string                   `json:"category_id"`
	UserID       uuid.UUID                `json:"user_id"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	Category     *entity.CategoryFiltered `json:"category"`
	User         *entity.UserFiltered     `json:"user"`
	Location     *entity.LocationFiltered `json:"location"`
}

type UpdateItemRequest struct {
	ID           uuid.UUID `param:"id"`
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Quantity     int       `json:"quantity" validate:"required"`
	Availability bool      `json:"availability" validate:"required"`
	CategoryID   string    `json:"category_id" validate:"required"`
	UserID       uuid.UUID `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

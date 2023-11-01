package models

import "github.com/google/uuid"

type CreateOrderCartRequest struct {
	ID                 uuid.UUID         `json:"id"`
	OrderID            uuid.UUID         `json:"order_id"`
	ItemIDWithQuantity map[uuid.UUID]int `json:"itemid_quantity"`
}

type CreateOrderCartResponse struct {
	OrderID  uuid.UUID `json:"order_id"`
	ItemID   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
}

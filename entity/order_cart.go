package entity

import "github.com/google/uuid"

type OrderCart struct {
	ID       uuid.UUID
	OrderID  uuid.UUID
	ItemID   uuid.UUID
	Quantity int
}

type OrderCartFiltered struct {
	OrderID uuid.UUID `json:"order_id"`
	ItemID  uuid.UUID `json:"item_id"`
}

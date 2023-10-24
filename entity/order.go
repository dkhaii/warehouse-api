package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                  uuid.UUID
	ItemID              uuid.UUID
	UserID              uuid.UUID
	Quantity            int
	Notes               string
	RequestTransferDate time.Time
	CreatedAt           time.Time
	User                *UserFiltered
	Item                []ItemFiltered
}

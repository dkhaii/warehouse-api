package entity

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                  uuid.UUID
	UserID              uuid.UUID
	Notes               string
	RequestTransferDate time.Time
	CreatedAt           time.Time
	User                *UserFiltered
	Items               []ItemFiltered
}

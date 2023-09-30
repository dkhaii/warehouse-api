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
	RequestTransferDate time.Time 
	Notes               string    
	CreatedAt           time.Time 
	UpdatedAt           time.Time 
}

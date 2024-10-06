package entities

import (
	"github.com/google/uuid"
)

type Buddy struct {
	BaseModel
	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid"`
	BuddyID uuid.UUID `json:"buddy_id" gorm:"type:uuid"`
	Status  string    `json:"status"`
}

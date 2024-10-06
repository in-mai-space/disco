package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	BaseModel
	SenderID   uuid.UUID `json:"sender_id" gorm:"type:uuid"`
	ReceiverID uuid.UUID `json:"receiver_id" gorm:"type:uuid"`
	Content    string    `json:"content"`
	Timestamp  time.Time `json:"timestamp"`
	IsRead     bool      `json:"is_read"`
}

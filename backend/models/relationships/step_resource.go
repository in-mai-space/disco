package models

import "github.com/google/uuid"

type StepResource struct {
	StepID     uuid.UUID `gorm:"primaryKey" json:"step_id"`
	ResourceID uuid.UUID `gorm:"primaryKey" json:"resource_id"`
}

package models

import "github.com/google/uuid"

type PlanResource struct {
	PlanID     uuid.UUID `gorm:"primaryKey" json:"plan_id"`
	ResourceID uuid.UUID `gorm:"primaryKey" json:"resource_id"`
}

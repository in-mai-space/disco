package models

import "github.com/google/uuid"

type GoalResource struct {
	GoalID     uuid.UUID `gorm:"primaryKey" json:"goal_id"`
	ResourceID uuid.UUID `gorm:"primaryKey" json:"resource_id"`
}

package models

import "github.com/google/uuid"

type PlanGoal struct {
	PlanID uuid.UUID `gorm:"primaryKey" json:"plan_id"`
	GoalID uuid.UUID `gorm:"primaryKey" json:"goal_id"`
}

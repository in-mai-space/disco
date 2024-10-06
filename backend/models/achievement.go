package models

import (
	"time"

	"github.com/google/uuid"
)

type Achievement struct {
	BaseModel
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	GoalID     *uuid.UUID `json:"goal_id" gorm:"type:uuid"`
	PlanID     *uuid.UUID `json:"plan_id" gorm:"type:uuid"`
	StepID     *uuid.UUID `json:"step_id" gorm:"type:uuid"`
	BadgeId    *uuid.UUID `json:"badge_id" gorm:"type:uuid"`
	DateEarned time.Time  `json:"date_earned"`
}

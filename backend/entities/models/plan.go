package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/in-mai-space/disco/types"
)

type Plan struct {
	BaseModel
	UserID         uuid.UUID    `json:"user_id" gorm:"type:uuid"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Status         types.Status `json:"status"`
	StartDate      time.Time    `json:"start_date"`
	EndDate        time.Time    `json:"end_date"`
	Goals          []Goal       `gorm:"many2many:plan_goals;" json:"goals"`
	Steps          []Step       `json:"steps" gorm:"foreignKey:PlanId"`
	Resources      []Resource   `gorm:"many2many:plan_resources;" json:"resources"`
	StudyFrequency Frequency    `json:"study_frequency"`
}

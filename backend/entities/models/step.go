package models

import (
	"github.com/google/uuid"
	"github.com/in-mai-space/disco/types"
)

type Step struct {
	BaseModel
	PlanId      uuid.UUID    `json:"plan_id" gorm:"type:uuid"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      types.Status `json:"status"`
	Entries     []Entry      `json:"entries" gorm:"foreignKey:StepId"`
	Resources   []Resource   `gorm:"many2many:step_resources;" json:"resources"`
}

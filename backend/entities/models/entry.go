package entities

import (
	"github.com/google/uuid"
)

type Entry struct {
	BaseModel
	StepId     uuid.UUID `json:"step_id"`
	Duration   string    `json:"duration"`
	Reflection string    `json:"reflection"`
}

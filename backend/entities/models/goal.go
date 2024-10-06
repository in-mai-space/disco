package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/in-mai-space/disco/types"
)

type Goal struct {
	BaseModel
	UserID      uuid.UUID    `json:"user_id" gorm:"type:uuid"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      types.Status `json:"status"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Plans       []Plan       `gorm:"many2many:plan_goals;" json:"plans"`
	Resources   []Resource   `gorm:"many2many:goal_resources;" json:"resources"`
}

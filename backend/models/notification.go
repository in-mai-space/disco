package models

import "github.com/google/uuid"

type Notification struct {
	BaseModel
	UserID        uuid.UUID        `json:"user_id"`
	Message       string           `json:"message"`
	Read          bool             `json:"read"`
	Type          NotificationType `json:"type"`
	PlanID        *uuid.UUID       `json:"plan_id,omitempty"`
	AchievementID *uuid.UUID       `json:"achievement_id,omitempty"`
	BuddyID       *uuid.UUID       `json:"buddy_id,omitempty"`
}

type NotificationType string

const (
	StudyFrequencyNotification NotificationType = "study_frequency"
	AchievementNotification    NotificationType = "achievement"
	BuddyMessageNotification   NotificationType = "buddy_message"
)

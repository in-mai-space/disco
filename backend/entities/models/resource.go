package entities

type Resource struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Goals       []Goal `gorm:"many2many:goal_resources;" json:"goals"`
	Plans       []Plan `gorm:"many2many:plan_resources;" json:"plans"`
	Steps       []Step `gorm:"many2many:step_resources;" json:"steps"`
}

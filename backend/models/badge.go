package models

type Badge struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
}

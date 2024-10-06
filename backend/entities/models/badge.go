package entities

type Badge struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
}

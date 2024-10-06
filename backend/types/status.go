package types

type Status string

const (
	NotStarted Status = "Not Started"
	InProgress Status = "In Progress"
	Completed  Status = "Completed"
	OnHold     Status = "On Hold"
	Canceld    Status = "Canceld"
)

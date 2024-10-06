package models

type FrequencyUnit string

const (
	Day   FrequencyUnit = "day"
	Week  FrequencyUnit = "week"
	Month FrequencyUnit = "month"
)

type Frequency struct {
	Interval     int           `json:"interval"`
	Unit         FrequencyUnit `json:"unit"`
	SpecificDays []string      `json:"specific_days"`
}

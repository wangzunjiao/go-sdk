package model

type PeriodRule struct {
	PeriodType  string `json:"periodType,omitempty"`
	PeriodCount int    `json:"periodCount,omitempty"`
}

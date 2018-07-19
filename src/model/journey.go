package model

import "time"

const (
	JourneyStatusFree int = iota
	JourneyStatusOngoing
	JourneyStatusFinish
)

type Journey struct {
	ID           int       `json:"id"`
	From         Location  `json:"from"`
	To           Location  `json:"to"`
	Steps        int       `json:"steps"`
	CurrentSteps int       `json:"current_steps"`
	StartTime    time.Time `json:"start_time,string"`
	EndTime      time.Time `json:"end_time,string"`
	UpdateTime   time.Time `json:"update_time,string"`
	Status       int       `json:"status"` // 0:abandon;1:ongoing;2:finish
	Members      []User    `json:"members"`
}

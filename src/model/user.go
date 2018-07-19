package model

import "time"

type User struct {
	ID           int       `json:"id"`
	WXID         string    `json:"wxid"`
	WXName       string    `json:"wxname"`
	Ride         Transport `json:"ride"`
	Steps        int       `json:"steps"`
	JourneySteps int       `json:"journey_steps"`
	UpdateTime   time.Time `json:"update_time,string"`
	Journey      Journey   `json:"journey"`
}

type UserTransport struct {
	ID   int `json:"id"`
	list []Transport
}

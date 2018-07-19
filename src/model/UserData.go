package model

type UserData struct {
	UID       int        `json:"uid"`
	Score     int        `json:"score"`
	PostCards []PostCard `json:"postCards"`
}

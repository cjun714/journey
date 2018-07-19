package model

type Province struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	TotalCards int    `json:"total_cards"`
}

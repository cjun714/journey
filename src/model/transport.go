package model

type Transport struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Rate float32 `json:"rate"`
}

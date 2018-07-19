package model

type PostCard struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImgURL      string   `json:"img_url"`
	Province    Province `json:"province"`
}

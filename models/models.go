package models

type Comment struct {
	Country string `json:"country"`
	Comment string `json:"comment"`
}

type BodyData struct {
	Country string `json:"country"`
	Comment string `json:"comment"`
}
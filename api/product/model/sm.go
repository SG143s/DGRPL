package model

type Sm struct {
	ID    string `json:"prID"`
	Own   string `json:"owID"`
	Name  string `json:"prName"`
	Price int    `json:"prPr"`
	Image string `json:"prImage"`
}

package model

type Pr struct {
	ID    string `json:"prID"`
	Name  string `json:"prName"`
	Own   string `json:"owID"`
	Price int    `json:"prPr"`
	Stock int    `json:"prSt"`
}

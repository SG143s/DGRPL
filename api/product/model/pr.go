package model

type Pr struct {
	ID       string   `json:"prID"`
	Own      string   `json:"owID"`
	Name     string   `json:"prName"`
	TimeUsed int      `json:"prUsed"`
	Desc     string   `json:"prDesc"`
	Stock    int      `json:"prSt"`
	Price    int      `json:"prPr"`
	Image    []string `json:"imgs"`
}

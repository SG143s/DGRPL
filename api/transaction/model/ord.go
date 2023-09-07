package model

type Ord struct {
	ID     string `json:"ordID"`
	OwID   string `json:"ownID"`
	PrName string `json:"prName"`
	PrImg  string `json:"prImg"`
	UPrice int    `json:"uPrice"`
	Stock  int    `json:"prStock"`
}

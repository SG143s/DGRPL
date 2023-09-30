package model

type Ins struct {
	ID     string `json:"ordID"`
	BuyID  string `json:"buyID"`
	SellID string `json:"sellID"`
	PrID   string `json:"prID"`
	Quant  int    `json:"quant"`
}

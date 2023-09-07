package model

type Tuser struct {
	ID       string `json:"uid"`
	Name     string `json:"ufname"`
	Pic      string `json:"upic"`
	Uname    string `json:"uname"`
	Password string `json:"upass"`
	Email    string `json:"umail"`
	Balance  int    `json:"ubal"`
}

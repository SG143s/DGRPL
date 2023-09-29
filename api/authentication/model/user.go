package model

type User struct {
	ID       string `json:"uid"`
	Uname    string `json:"uname"`
	Email    string `json:"umail"`
	Password string `json:"upass"`
	Name     string `json:"ufname"`
	Pic      string `json:"upic"`
}

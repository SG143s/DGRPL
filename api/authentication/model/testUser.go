package model

type Tuser struct {
	ID       string `json:"uid"`
	Name     string `json:"uname"`
	Email    string `json:"umail"`
	Password string `json:"upass"`
}

package user

type Loginf struct {
	Name string `json:"uname"`
	Pass string `json:"upass"`
}

type Reginf struct {
	Uname    string `json:"uname"`
	Email    string `json:"umail"`
	Password string `json:"upass"`
	Name     string `json:"ufname"`
}

type Small struct {
	Id   string `json:"uid"`
	Name string `json:"uname"`
	PPic string `json:"upic"`
}

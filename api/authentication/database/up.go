package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/authentication/model"
)

func Reg(data md.Tuser) bool {
	_, err := db.Query("INSERT INTO tusers values (?, ?, ?, ?, ?, ?, ?)", data.ID, data.Name, data.Pic, data.Uname, data.Password, data.Email, data.Balance)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/authentication/model"
)

func Reg(data md.User) bool {
	_, err := db.Query("CALL register (?, ?, ?, ?, ?, ?)", data.ID, data.Uname, data.Email, data.Password, data.Name, data.Pic)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

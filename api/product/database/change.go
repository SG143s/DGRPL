package database

import "fmt"

func Delete(id string) bool {
	_, err := db.Query("DELETE FROM tproducts WHERE prodID = ?", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

package database

import "fmt"

func Updatestat(id string, stat string) bool {
	_, err := db.Query("CALL updatestat (?, ?)", id, stat)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

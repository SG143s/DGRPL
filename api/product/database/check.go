package database

import "fmt"

func ChId(id string) bool {
	var res int
	row, err := db.Query("CALL chid (?)", id)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&res)
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return res != 1
}

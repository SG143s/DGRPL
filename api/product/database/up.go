package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/product/model"
)

func Add(data md.Pr) bool {
	_, err := db.Query("CALL add (?, ?, ?, ?, ?, ?, ?)", data.ID, data.Own, data.Name, data.TimeUsed, data.Desc, data.Price, data.Stock)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = db.Query("CALL addimg (?, ?)", data.ID, "https://picsum.photos/500/500")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

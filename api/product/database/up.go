package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/product/model"
)

func Add(data md.Pr) bool {
	_, err := db.Query("INSERT INTO tproducts values (?, ?, ?, ?, ?, ?, ?)", data.ID, data.Own, data.Name, data.Image, data.Price, data.Stock)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

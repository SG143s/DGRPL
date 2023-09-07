package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/product/model"
)

func Search(val string) []md.Pr {
	var data []md.Pr
	var temp md.Pr

	row, err := db.Query("SELECT * FROM tproducts WHERE productName LIKE ?", "%"+val+"%")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.Own, &temp.Name, &temp.Image, &temp.Price, &temp.Stock)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		data = append(data, temp)
	}
	return data
}

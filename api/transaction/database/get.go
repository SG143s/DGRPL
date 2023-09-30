package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/transaction/model"
)

func SAll(am string) []md.Ord {
	var data []md.Ord
	var temp md.Ord

	row, err := db.Query("CALL getdata(?)", am)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PrID, &temp.BuyID, &temp.SellID, &temp.Quant, &temp.Payid, &temp.Status)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	return data
}

func SSell(am string) []md.Ord {
	var data []md.Ord
	var temp md.Ord

	row, err := db.Query("CALL getsell(?)", am)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer row.Close()

	for row.Next() {
		err := row.Scan(&temp.ID, &temp.PrID, &temp.BuyID, &temp.SellID, &temp.Quant, &temp.Payid, &temp.Status)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	return data
}

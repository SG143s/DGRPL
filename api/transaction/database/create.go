package database

import (
	"fmt"

	md "github.com/SG143s/DGRPL/api/transaction/model"
)

func Add(data md.Ord) bool {
	_, err := db.Query("CALL createord (?, ?, ?, ?, ?, ?)", data.ID, data.BuyID, data.SellID, data.PrID, data.Quant, data.Payid)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

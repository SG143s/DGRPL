package operation

import (
	//"github.com/gin-contrib/sessions"
	"fmt"

	db "github.com/SG143s/DGRPL/api/transaction/database"
	str "github.com/SG143s/DGRPL/api/transaction/model"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var ins str.Ins
	if err := c.ShouldBindJSON(&ins); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		fmt.Println(err)
		return
	}
	var info str.Ord
	info.ID = ins.ID
	info.BuyID = ins.BuyID
	info.SellID = ins.SellID
	info.PrID = ins.PrID
	info.Quant = ins.Quant
	info.Payid = Genpayid()
	if !db.Add(info) {
		c.JSON(400, "error : invalid data")
		return
	} else {
		c.Status(200)
		return
	}
}

func Setstat(c *gin.Context) {
	orid := c.Param("orid")
	st := c.Param("stid")
	stat := false
	if st == "0" {
		stat = db.Updatestat(orid, "unconfirmed")
	} else if st == "1" {
		stat = db.Updatestat(orid, "processing")
	} else if st == "2" {
		stat = db.Updatestat(orid, "shipping")
	} else if st == "3" {
		stat = db.Updatestat(orid, "completed")
	} else if st == "4" {
		stat = db.Updatestat(orid, "canceled")
	}
	if stat {
		c.Status(200)
		return
	} else {
		c.Status(404)
	}
}

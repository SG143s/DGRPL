package operation

import (
	//"github.com/gin-contrib/sessions"
	"fmt"

	db "github.com/SG143s/DGRPL/api/product/database"
	str "github.com/SG143s/DGRPL/api/product/model"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var ins str.Ins
	if err := c.ShouldBindJSON(&ins); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		fmt.Println(err)
		return
	}
	var info str.Pr
	info.ID = Genid()
	info.Own = ins.Own
	info.Name = ins.Name
	info.Desc = ins.Desc
	info.Price = ins.Price
	info.Stock = ins.Stock
	info.TimeUsed = ins.TimeUsed
	if !db.Add(info) {
		c.JSON(400, "error : invalid data")
		return
	} else {
		c.Status(200)
		return
	}
}

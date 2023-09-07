package operation

import (
	//"github.com/gin-contrib/sessions"
	db "github.com/SG143s/DGRPL/api/product/database"
	str "github.com/SG143s/DGRPL/api/product/model"
	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var info str.Pr
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}
	if !db.Add(info) {
		c.JSON(400, "error : invalid data")
	} else {
		c.Status(200)
	}
}

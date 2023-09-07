package operation

import (
	db "github.com/SG143s/DGRPL/api/authentication/database"
	md "github.com/SG143s/DGRPL/api/authentication/model"
	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Context) {
	var info md.Tuser
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}
	if !db.Reg(info) {
		c.JSON(400, "error : invalid data")
	} else {
		c.Status(200)
	}

}

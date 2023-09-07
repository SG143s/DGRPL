package operation

import (
	//"github.com/gin-contrib/sessions"
	str "github.com/SG143s/DGRPL/api/product/model"
	"github.com/gin-gonic/gin"
)

func Upd(c *gin.Context) {
	var info str.Pr
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}
	//update operation
	c.Status(200)
}

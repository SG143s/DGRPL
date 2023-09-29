package operation

import (
	db "github.com/SG143s/DGRPL/api/product/database"
	"github.com/gin-gonic/gin"
)

func Del(c *gin.Context) {
	id := c.Param("id")
	if !db.Delete(id) {
		c.JSON(403, "error:invalid input")
		return
	} else {
		c.Status(200)
	}
}

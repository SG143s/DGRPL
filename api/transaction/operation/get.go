package operation

import (
	db "github.com/SG143s/DGRPL/api/transaction/database"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	id := c.Param("id")

	data := db.SAll(id)

	c.JSON(200, data)
}

func SearchSell(c *gin.Context) {
	id := c.Param("id")

	data := db.SSell(id)

	c.JSON(200, data)
}

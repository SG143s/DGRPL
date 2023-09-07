package operation

import (
	db "github.com/SG143s/DGRPL/api/product/database"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	id := c.Param("id")

	data := db.Search(id)

	c.JSON(200, data)
}

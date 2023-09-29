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

func SinP(c *gin.Context) {
	id := c.Param("id")
	if db.ChId(id) {
		c.Status(404)
		return
	}
	c.JSON(200, db.Pr(id))
}

func Get3(c *gin.Context) {
	data := db.SAll(3)

	c.JSON(200, data)
}

func Get4(c *gin.Context) {
	data := db.SAll(4)

	c.JSON(200, data)
}

package operation

import "github.com/gin-gonic/gin"

func Del(c *gin.Context) {
	id := c.Param("id")
	//delete op
	c.JSON(200, id)
}

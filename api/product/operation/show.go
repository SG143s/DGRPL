package operation

import "github.com/gin-gonic/gin"

func Sh(c *gin.Context) {
	c.JSON(200, "no database currently")
}

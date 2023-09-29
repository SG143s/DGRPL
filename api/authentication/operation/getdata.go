package operation

import (
	db "github.com/SG143s/DGRPL/api/authentication/database"
	"github.com/gin-gonic/gin"
)

func Getdata(c *gin.Context) {
	idinf := c.Param("id")

	c.JSON(200, db.Getdat(idinf))
}

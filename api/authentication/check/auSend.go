package check

import (
	"github.com/SG143s/DGRPL/api/authentication/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SendAu(c *gin.Context) {
	sess := sessions.Default(c)
	var audata model.Small
	audata.Id = sess.Get("userID").(string)
	audata.Name = sess.Get("userName").(string)
	c.JSON(200, audata)
}

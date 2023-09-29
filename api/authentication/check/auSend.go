package check

import (
	"github.com/SG143s/DGRPL/api/authentication/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"fmt"
)

func SendAu(c *gin.Context) {
	sess := sessions.Default(c)
	logIn := sess.Get("loggedIn")

	if logInBool, ok := logIn.(bool); ok && logInBool {
		var audata model.Small
		audata.Id = sess.Get("uid").(string)
		audata.Name = sess.Get("uname").(string)
		c.JSON(200, audata)
	} else {
		c.JSON(401, gin.H{"error": "not logged in"})
	}
}

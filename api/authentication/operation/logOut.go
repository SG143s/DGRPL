package operation

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {
	session := sessions.Default(c)

	session.Clear()
	session.Set("loggedIn", false)
	session.Save()
	c.Status(200)
	return
}

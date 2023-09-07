package operation

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	sess := sessions.Default(c)
	//uname := c.Param("username")
	//pass := c.Param("pass")

	//login check operation

	sess.Set("loggedIn", true)
	sess.Save()

	c.Status(200)
}

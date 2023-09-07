package operation

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Set("loggedIn", true)
	sess.Save()
}

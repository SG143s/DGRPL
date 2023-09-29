package operation

import (
	db "github.com/SG143s/DGRPL/api/authentication/database"
	md "github.com/SG143s/DGRPL/api/authentication/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogIn(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Set("loggedIn", false)
	sess.Save()
	var loginf md.Loginf
	if err := c.ShouldBindJSON(&loginf); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}

	if !db.ChUn(loginf.Name) {
		c.JSON(401, gin.H{"error": "invalid username"})
		return
	} else {
		if !db.ChLog(loginf.Name, loginf.Pass) {
			c.JSON(401, gin.H{"error": "invalid password"})
			return
		} else {
			uid := db.Getidup(loginf.Name, loginf.Pass)
			sess.Set("loggedIn", true)
			sess.Set("uid", uid)
			sess.Set("uname", loginf.Name)
			sess.Save()
			c.JSON(200, gin.H{"userid": uid})
			return
		}
	}
}

package operation

import (
	//"fmt"

	db "github.com/SG143s/DGRPL/api/authentication/database"
	md "github.com/SG143s/DGRPL/api/authentication/model"
	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Context) {
	var reginf md.Reginf
	if err := c.ShouldBindJSON(&reginf); err != nil {
		c.JSON(400, gin.H{"error": "Invalid requestt"})
		return
	}

	if !db.ChEmail(reginf.Email) {
		//fmt.Println("email check fail")
		c.JSON(400, gin.H{"error": "email already registered"})
		return
	}
	//fmt.Println("username check run")
	if db.ChUn(reginf.Uname) {
		//fmt.Println("username check fail")
		c.JSON(400, gin.H{"error": "username " + reginf.Uname + " is taken"})
		return
	}
	if len(reginf.Password) < 8 {
		c.JSON(400, gin.H{"error": "minimal password length is 8 characters"})
		return
	}

	var info md.User
	info.ID = Genid()
	info.Name = reginf.Name
	info.Email = reginf.Email
	info.Password = reginf.Password
	info.Uname = reginf.Uname
	info.Pic = "https://picsum.photos/1000/1000"
	//fmt.Println("insert system run")
	if !db.Reg(info) {
		c.JSON(400, gin.H{"error": "invalid data"})
		return
	} else {
		c.JSON(200, gin.H{"userid": info.ID})
		return
	}

}

package main

import (
	ch "github.com/SG143s/DGRPL/api/authentication/check"
	db "github.com/SG143s/DGRPL/api/authentication/database"
	op "github.com/SG143s/DGRPL/api/authentication/operation"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	//config.AllowAllOrigins = true
	r.Use(cors.New(config))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	db.SqlStart()
	defer db.SqlStop()

	r.POST("/register", op.Reg)
	r.POST("/login", op.LogIn)
	r.POST("/logout", op.LogOut)
	r.GET("/getauinformation", ch.SendAu)

	r.Run(":8001")
}

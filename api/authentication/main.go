package main

import (
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
	config.AllowOrigins = []string{"http://localhost:5173"}
	r.Use(cors.New(config))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	db.SqlStart()

	r.POST("/register", op.Reg)

	r.Run(":8001")
}

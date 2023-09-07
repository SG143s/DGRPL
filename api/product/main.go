package main

import (
	op "github.com/SG143s/DGRPL/api/product/operation"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	r.Use(cors.New(config))
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/add", op.Add)

	r.DELETE("/delete/:id", op.Del)

	r.PUT("/update", op.Upd)

	r.GET("/product", op.Sh)

	r.Run(":8002")
}

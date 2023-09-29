package main

import (
	db "github.com/SG143s/DGRPL/api/product/database"
	op "github.com/SG143s/DGRPL/api/product/operation"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	r.Use(cors.New(config))

	db.SqlStart()

	r.POST("/add", op.Add)

	r.DELETE("/delete/:id", op.Del)

	r.PUT("/update", op.Upd)

	r.GET("/search/:id", op.Search)
	r.GET("/getfourprod", op.Get4)
	r.GET("/getthreeprod", op.Get3)

	r.Run(":8002")
}

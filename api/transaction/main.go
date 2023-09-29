package main

import (
	db "github.com/SG143s/DGRPL/api/transaction/database"
	op "github.com/SG143s/DGRPL/api/transaction/operation"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	r.Use(cors.New(config))

	db.SqlStart()
	defer db.SqlStop()

	r.GET("/ordid", func(c *gin.Context) {
		c.JSON(200, gin.H{"orderid": op.Genid()})
	})

	r.Run(":8003")
}

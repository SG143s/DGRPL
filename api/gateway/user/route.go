package user

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", Register)
		userGroup.POST("/login", Login)
		userGroup.POST("/logout", Logout)
		userGroup.GET("/getauinf", GetAu)
	}
}

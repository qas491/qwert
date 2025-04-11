package router

import (
	"api/handler/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/register", api.Register)
		user.POST("/login", api.Login)
	}
}

package router

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	router := r.Group("/v1")
	{
		apiRouter := router.Group("/api")
		{
			UserRouter(apiRouter)
		}
	}
}

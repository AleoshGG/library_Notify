package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	notifyRoutes := r.Group("/notify")
	{
		notifyRoutes.POST("/lend")
		notifyRoutes.POST("/return")
	}
}
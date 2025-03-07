package routes

import (
	"library-Notify/src/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	notifyRoutes := r.Group("/notify")
	{
		notifyRoutes.POST("/lend", controllers.NewNotifyOfLendController().NotifyOfLend)
		notifyRoutes.POST("/return", controllers.NewNotifyOfReturnController().NotifyOfReturn)
	}
}
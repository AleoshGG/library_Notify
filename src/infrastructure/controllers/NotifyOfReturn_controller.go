package controllers

import (
	"library-Notify/src/application/services"
	"library-Notify/src/infrastructure"

	"github.com/gin-gonic/gin"
)

type NotifyOfReturnController struct {
	service *services.NotifyOfReturnEvent
}

func NewNotifyOfReturnController() *NotifyOfReturnController {
	n := infrastructure.GetNotifier()
	service := services.NewNotifyOfReturnEvent(n)
	return &NotifyOfReturnController{service: service}
}

func (n *NotifyOfReturnController) NotifyOfReturn(c *gin.Context) {
	n.service.Run("Mensaje", "Correo Electrónico", "Nombre")
}


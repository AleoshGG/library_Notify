package controllers

import (
	"library-Notify/src/application/services"
	"library-Notify/src/infrastructure"

	"github.com/gin-gonic/gin"
)

type NotifyOfLendController struct {
	service *services.NotifyOfLendEvent
}

func NewNotifyOfLendController() *NotifyOfLendController {
	n := infrastructure.GetNotifier()
	service := services.NewNotifyOfLendEvent(n)
	return &NotifyOfLendController{service: service}
}

func (n *NotifyOfLendController) NotifyOfLend(c *gin.Context) {
	n.service.Run("Mensaje", "Correo Electrónico", "Fecha de devolución", "Nombre")
}


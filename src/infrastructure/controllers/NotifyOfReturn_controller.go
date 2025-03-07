package controllers

import (
	"library-Notify/src/application/services"
	"library-Notify/src/domain"
	"library-Notify/src/infrastructure"
	"library-Notify/src/infrastructure/controllers/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyOfReturnController struct {
	service *services.NotifyOfReturnEvent
	fetchAPI *services.FetchAPIService
}

func NewNotifyOfReturnController() *NotifyOfReturnController {
	n := infrastructure.GetNotifier()
	f := infrastructure.GetFeychAPI()
	service := services.NewNotifyOfReturnEvent(n)
	fetchAPI := services.NewFetchAPIService(f)
	return &NotifyOfReturnController{service: service, fetchAPI: fetchAPI}
}

func (n *NotifyOfReturnController) NotifyOfReturn(c *gin.Context) {
	var notify domain.Notify

	if err := c.ShouldBindJSON(&notify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	if err := validators.Checknotify(notify); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	responseReader, err := n.fetchAPI.Run(int(notify.Id_reader))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error": err.Error(),
		})
		return
	}

	n.service.Run("Gracias por devolver el libro", responseReader.Data[0].Email, responseReader.Data[0].First_name+responseReader.Data[0].Last_name)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}


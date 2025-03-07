package controllers

import (
	"library-Notify/src/application/services"
	"library-Notify/src/domain"
	"library-Notify/src/infrastructure"
	"library-Notify/src/infrastructure/controllers/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotifyOfLendController struct {
	service *services.NotifyOfLendEvent
	fetchAPI *services.FetchAPIService
}

func NewNotifyOfLendController() *NotifyOfLendController {
	n := infrastructure.GetNotifier()
	f := infrastructure.GetFeychAPI()
	service := services.NewNotifyOfLendEvent(n)
	fetchAPI := services.NewFetchAPIService(f)
	return &NotifyOfLendController{service: service, fetchAPI: fetchAPI}
}

func (n *NotifyOfLendController) NotifyOfLend(c *gin.Context) {
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
	
	n.service.Run("Tu prestamo ha sido registrado, no olvides regresar el libro antes de la fecha: ", responseReader.Data[0].Email, notify.Return_date, responseReader.Data[0].First_name+responseReader.Data[0].Last_name)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}


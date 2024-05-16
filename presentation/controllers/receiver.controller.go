package controllers

import (
	"net/http"

	applicationreceiver "github.com/RenzoReccio/project-management.worker/application/receiver/insertReceiver"
	"github.com/gin-gonic/gin"
)

type ReceiverController struct {
	receiverUseCase *applicationreceiver.InsertReceiverUseCase
}

func NewReceiverController(receiverUseCase *applicationreceiver.InsertReceiverUseCase) *ReceiverController {
	return &ReceiverController{
		receiverUseCase: receiverUseCase,
	}
}
func (u *ReceiverController) InsertReceiver(c *gin.Context) {
	req := new(applicationreceiver.InsertReceiverDto)

	if err := c.Bind(req); err != nil {
		c.IndentedJSON(http.StatusCreated, "Bad request")
		return
	}
	c.IndentedJSON(http.StatusCreated, req)
}

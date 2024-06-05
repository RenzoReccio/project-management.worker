package controllers

import (
	"net/http"

	application_createevent "github.com/RenzoReccio/project-management.worker/application/event/command/create-event"
	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

type EventController struct {
}

func NewEventController() *EventController {
	return &EventController{}
}
func (u *EventController) InsertEvent(c *gin.Context) {
	req := new(application_createevent.CreateProductCommand)

	if err := c.Bind(req); err != nil {
		c.IndentedJSON(http.StatusCreated, "Bad request")
		return
	}
	mediatr.Send[*application_createevent.CreateProductCommand, *string](c, req)

	c.IndentedJSON(http.StatusCreated, req)
}

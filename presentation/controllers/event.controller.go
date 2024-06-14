package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"

	application_getepic "github.com/RenzoReccio/project-management.worker/application/epic/query"
	application_createevent "github.com/RenzoReccio/project-management.worker/application/event/command/create-event"
	application_getworkitemtype "github.com/RenzoReccio/project-management.worker/application/workItemType/query/get-work-item-type"
	"github.com/RenzoReccio/project-management.worker/domain/model"
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
		c.IndentedJSON(http.StatusBadRequest, model_shared.NewResultFailure(model_shared.NewError("BAD_REQUEST", "Input not correct format")))
		return
	}
	ProcessEvent(c, req)

	// c.IndentedJSON(http.StatusCreated, model_shared.NewResultSuccess())
}

func ProcessEvent(c *gin.Context, req *application_createevent.CreateProductCommand) {
	resultEvent, _ := mediatr.Send[*application_createevent.CreateProductCommand, *model_shared.ResultWithValue[model.Event]](c, req)
	if !resultEvent.IsSuccess {
		fmt.Println(resultEvent.Error)
		return
	}
	event := resultEvent.Result()
	fmt.Println(event)

	getWorkItemTypeQuery := &application_getworkitemtype.GetWorkItemTypeQuery{
		ResourceURL: event.ResourceUrl,
	}

	resultgetWorkItemTypeQuery, _ := mediatr.Send[*application_getworkitemtype.GetWorkItemTypeQuery, *model_shared.ResultWithValue[model.WorkItemType]](c, getWorkItemTypeQuery)
	if !resultgetWorkItemTypeQuery.IsSuccess {
		fmt.Println(resultgetWorkItemTypeQuery.Error)
		return
	}
	fmt.Println(*resultgetWorkItemTypeQuery.Result())

	epic, _ := mediatr.Send[*application_getepic.GetEpicQuery, *model_shared.ResultWithValue[model.Epic]](c, &application_getepic.GetEpicQuery{ResourceURL: event.ResourceUrl})
	b, err := json.Marshal(epic.Result())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	c.IndentedJSON(http.StatusCreated, epic.Result())

	// if err != nil {
	// 	fmt.Println(task)
	// 	return
	// }
}

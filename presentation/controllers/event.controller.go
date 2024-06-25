package controllers

import (
	"context"
	"fmt"
	"net/http"

	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	factory_workitem "github.com/RenzoReccio/project-management.worker/presentation/factory/workItem"

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
	go ProcessEvent(req)

	c.IndentedJSON(http.StatusCreated, model_shared.NewResultSuccess())
}

func ProcessEvent(req *application_createevent.CreateProductCommand) {
	context := context.Background()

	resultEvent, _ := mediatr.Send[*application_createevent.CreateProductCommand, *model_shared.ResultWithValue[model.Event]](context, req)
	if !resultEvent.IsSuccess {
		fmt.Println(resultEvent.Error)
		return
	}
	event := resultEvent.Result()

	getWorkItemTypeQuery := &application_getworkitemtype.GetWorkItemTypeQuery{
		ResourceURL: event.ResourceUrl,
	}

	resultgetWorkItemTypeQuery, _ := mediatr.Send[*application_getworkitemtype.GetWorkItemTypeQuery, *model_shared.ResultWithValue[model.WorkItemType]](context, getWorkItemTypeQuery)
	if !resultgetWorkItemTypeQuery.IsSuccess {
		fmt.Println(resultgetWorkItemTypeQuery.Error)
		return
	}
	fmt.Println(*resultgetWorkItemTypeQuery.Result())

	ExecuteEvent(context, event, resultgetWorkItemTypeQuery.Result())
}

func ExecuteEvent(context context.Context, event *model.Event, workItemType *model.WorkItemType) {
	factory := factory_workitem.GetWorkItemFactory(*workItemType.Type)
	factory.ExecuteWorkItem(context, event.ResourceUrl)
}

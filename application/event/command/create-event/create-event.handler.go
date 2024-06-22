package application_createevent

import (
	"context"
	"strconv"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type CreateEventCommandHandler struct {
	eventRepository repository.EventRepository
}

func NewCreateEventCommandHandler(eventRepository repository.EventRepository) *CreateEventCommandHandler {
	return &CreateEventCommandHandler{eventRepository: eventRepository}
}

func (c *CreateEventCommandHandler) Handle(ctx context.Context, command *CreateProductCommand) (*model_shared.ResultWithValue[model.Event], error) {
	resourceURL := command.ResourceContainers.Project.BaseURL + command.ResourceContainers.Project.ID + "/_apis/wit/workItems/" + getWorkItemId(command)
	event := model.NewEvent(command.ID, command.SubscriptionID, command.EventType, command.CreatedDate, command.Resource.ID, resourceURL)
	resultcreatedEvent := c.eventRepository.InsertEvent(event)
	if !resultcreatedEvent.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Event](model_shared.NewError("EVENT_NOT_CREATED", "Failure creating event.")), nil
	}

	return model_shared.NewResultWithValueSuccess[model.Event](resultcreatedEvent.Result()), nil
}

func getWorkItemId(command *CreateProductCommand) string {
	if command.Resource.WorkItemId != 0 {
		return strconv.Itoa(command.Resource.WorkItemId)
	}
	return strconv.Itoa(command.Resource.ID)
}

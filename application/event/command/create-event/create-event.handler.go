package application_createevent

import (
	"context"

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

	event := model.NewEvent(command.ID, command.SubscriptionID, command.EventType, command.CreatedDate, command.Resource.ID, command.Resource.URL)
	resultcreatedEvent := c.eventRepository.InsertEvent(event)
	if !resultcreatedEvent.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Event](model_shared.NewError("EVENT_NOT_CREATED", "Failure creating event.")), nil
	}

	return model_shared.NewResultWithValueSuccess[model.Event](resultcreatedEvent.Result()), nil
}

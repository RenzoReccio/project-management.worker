package application_closeevent

import (
	"context"

	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type CloseEventCommandHandler struct {
	eventRepository repository.EventRepository
}

func NewCloseEventCommandHandler(eventRepository repository.EventRepository) *CloseEventCommandHandler {
	return &CloseEventCommandHandler{eventRepository: eventRepository}
}

func (c *CloseEventCommandHandler) Handle(ctx context.Context, command *CloseEventCommand) (*model_shared.ResultWithValue[string], error) {
	resultcreatedEvent := c.eventRepository.CloseEvent(command.Id)
	if !resultcreatedEvent.IsSuccess {
		return model_shared.NewResultWithValueFailure[string](model_shared.NewError("EVENT_NOT_CREATED", "Failure creating event.")), nil
	}

	return model_shared.NewResultWithValueSuccess[string](nil), nil
}

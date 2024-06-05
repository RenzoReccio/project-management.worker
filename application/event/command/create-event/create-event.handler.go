package application_createevent

import (
	"context"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type CreateEventCommandHandler struct {
	eventRepository repository.EventRepository
}

func NewCreateEventCommandHandler(eventRepository repository.EventRepository) *CreateEventCommandHandler {
	return &CreateEventCommandHandler{eventRepository: eventRepository}
}

func (c *CreateEventCommandHandler) Handle(ctx context.Context, command *CreateProductCommand) (*string, error) {

	event := model.NewEvent(command.ID, command.SubscriptionID, command.EventType, command.CreatedDate, command.Resource.ID, command.Resource.URL)
	createdEvent, err := c.eventRepository.InsertEvent(event)
	if err != nil {
		return nil, err
	}

	return &createdEvent.ID, nil
}

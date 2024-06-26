package application_sendepic

import (
	"context"
	"errors"

	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type SendEpicEventHandler struct {
	messageRepository repository.MessageRepository
}

func NewSendEpicEventHandler(
	messageRepository repository.MessageRepository,
) *SendEpicEventHandler {
	return &SendEpicEventHandler{
		messageRepository: messageRepository,
	}
}

func (c *SendEpicEventHandler) Handle(ctx context.Context, event *SendEpicEvent) error {
	result := c.messageRepository.SendEpic(event.Data)
	if !result.IsSuccess {
		return errors.New(result.Error.Description)
	}
	return nil
}

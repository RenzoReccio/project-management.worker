package application_sendtask

import (
	"context"
	"errors"

	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type SendUserStoryEventHandler struct {
	messageRepository repository.MessageRepository
}

func NewSendTaskEventHandler(
	messageRepository repository.MessageRepository,
) *SendUserStoryEventHandler {
	return &SendUserStoryEventHandler{
		messageRepository: messageRepository,
	}
}

func (c *SendUserStoryEventHandler) Handle(ctx context.Context, event *SendTaskEvent) error {
	result := c.messageRepository.SendTask(event.Data)
	if !result.IsSuccess {
		return errors.New(result.Error.Description)
	}
	return nil
}

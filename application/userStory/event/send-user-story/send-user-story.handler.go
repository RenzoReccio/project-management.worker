package application_senduserstory

import (
	"context"
	"errors"

	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type SendUserStoryEventHandler struct {
	messageRepository repository.MessageRepository
}

func NewSendUserStoryEventHandler(
	messageRepository repository.MessageRepository,
) *SendUserStoryEventHandler {
	return &SendUserStoryEventHandler{
		messageRepository: messageRepository,
	}
}

func (c *SendUserStoryEventHandler) Handle(ctx context.Context, event *SendUserStoryEvent) error {
	result := c.messageRepository.SendUserStory(event.Data)
	if !result.IsSuccess {
		return errors.New(result.Error.Description)
	}
	return nil
}

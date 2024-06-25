package application_sendfeature

import (
	"context"
	"errors"

	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type SendFeatureEventHandler struct {
	messageRepository repository.MessageRepository
}

func NewSendFeatureEventHandler(
	messageRepository repository.MessageRepository,
) *SendFeatureEventHandler {
	return &SendFeatureEventHandler{
		messageRepository: messageRepository,
	}
}

func (c *SendFeatureEventHandler) Handle(ctx context.Context, event *SendFeatureEvent) error {
	result := c.messageRepository.SendFeature(event.Data)
	if !result.IsSuccess {
		return errors.New(result.Error.Description)
	}
	return nil
}

package application_sendfeature

import (
	"context"

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
	c.messageRepository.SendFeature(event.Data)
	return nil
}

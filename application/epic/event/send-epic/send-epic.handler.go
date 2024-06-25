package application_sendepic

import (
	"context"

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
	//Do something with the event here !
	c.messageRepository.SendEpic(event.Data)
	return nil
}

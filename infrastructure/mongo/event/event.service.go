package mongoInfraestructure

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type EventService struct{}

func NewEventService() repository.EventRepository {
	return &EventService{}
}

func (u *EventService) InsertEvent(in *model.Event) (*model.Event, error) {
	return in, nil
}

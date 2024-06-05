package repository

import "github.com/RenzoReccio/project-management.worker/domain/model"

type EventRepository interface {
	InsertEvent(in *model.Event) (*model.Event, error)
}

package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type EventRepository interface {
	InsertEvent(in *model.Event) *model_shared.ResultWithValue[model.Event]
	CloseEvent(in string) *model_shared.Result
}

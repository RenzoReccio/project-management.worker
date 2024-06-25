package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type TaskRepository interface {
	GetTask(taskURL string) (*model_shared.ResultWithValue[model.Task], string)
}

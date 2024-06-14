package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type EpicRepository interface {
	GetEpic(url string) *model_shared.ResultWithValue[model.Epic]
}

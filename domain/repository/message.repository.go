package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type MessageRepository interface {
	SendEpic(in *model.Epic) *model_shared.ResultWithValue[string]
	SendFeature(in *model.Feature) *model_shared.ResultWithValue[string]
	SendUserStory(in *model.UserStory) *model_shared.ResultWithValue[string]
}

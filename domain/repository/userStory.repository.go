package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type UserStoryRepository interface {
	GetUserStory(url string) (*model_shared.ResultWithValue[model.UserStory], string)
}

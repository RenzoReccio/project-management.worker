package application_senduserstory

import "github.com/RenzoReccio/project-management.worker/domain/model"

type SendUserStoryEvent struct {
	Data *model.UserStory
}

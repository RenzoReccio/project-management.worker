package factory_workitem

import (
	"context"
	"fmt"

	application_senduserstory "github.com/RenzoReccio/project-management.worker/application/userStory/event/send-user-story"
	application_getuserstory "github.com/RenzoReccio/project-management.worker/application/userStory/query/get-user-story"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

type UserStoryFactory struct {
}

func NewUserStoryFactory() IWorkItem {
	return &UserStoryFactory{}
}

func (c UserStoryFactory) ExecuteWorkItem(context context.Context, resourceURL string) {
	resultUserStory, _ := mediatr.Send[*application_getuserstory.GetUserStoryQuery, *model_shared.ResultWithValue[model.UserStory]](
		context, &application_getuserstory.GetUserStoryQuery{ResourceURL: resourceURL})
	if !resultUserStory.IsSuccess {
		fmt.Print(resultUserStory.Error)
		return
	}

	sendUserStoryEvent := &application_senduserstory.SendUserStoryEvent{Data: resultUserStory.Result()}
	mediatr.Publish(context, sendUserStoryEvent)
}

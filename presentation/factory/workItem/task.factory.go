package factory_workitem

import (
	"context"
	"fmt"

	application_sendtask "github.com/RenzoReccio/project-management.worker/application/task/event/send-task"
	application_gettask "github.com/RenzoReccio/project-management.worker/application/task/query/get-task"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

type TaskFactory struct {
}

func NewTaskFactory() IWorkItem {
	return &TaskFactory{}
}

func (c TaskFactory) ExecuteWorkItem(context context.Context, resourceURL string) {
	resultTask, _ := mediatr.Send[*application_gettask.GetTaskQuery, *model_shared.ResultWithValue[model.Task]](
		context, &application_gettask.GetTaskQuery{ResourceURL: resourceURL})
	if !resultTask.IsSuccess {
		fmt.Print(resultTask.Error)
		return
	}
	sendTaskEvent := &application_sendtask.SendTaskEvent{Data: resultTask.Result()}
	mediatr.Publish(context, sendTaskEvent)
}

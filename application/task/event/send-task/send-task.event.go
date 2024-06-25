package application_sendtask

import "github.com/RenzoReccio/project-management.worker/domain/model"

type SendTaskEvent struct {
	Data *model.Task
}

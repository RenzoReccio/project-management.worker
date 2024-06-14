package application_gettask

import (
	"context"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type GetTaskQueryHandler struct {
	taskRepository repository.TaskRepository
}

func NewGetTaskQueryHandler(taskRepository repository.TaskRepository) *GetTaskQueryHandler {
	return &GetTaskQueryHandler{taskRepository: taskRepository}
}

func (c *GetTaskQueryHandler) Handle(ctx context.Context, query *GetTaskQuery) (*model.Task, error) {

	task, err := c.taskRepository.GetTask(query.ResourceURL)
	if err != nil {
		return nil, err
	}

	taskComments, err := c.taskRepository.GetTaskComments(query.ResourceURL)
	if err != nil {
		return nil, err
	}

	task.Comments = taskComments
	return task, nil
}

package repository

import "github.com/RenzoReccio/project-management.worker/domain/model"

type TaskRepository interface {
	GetTask(taskURL string) (*model.Task, error)
	GetTaskComments(taskURL string) (*[]model.Comment, error)
}

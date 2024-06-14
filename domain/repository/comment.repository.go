package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type CommentRepository interface {
	GetComments(url string) *model_shared.ResultWithValue[[]model.Comment]
}

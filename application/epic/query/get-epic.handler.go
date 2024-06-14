package application_getepic

import (
	"context"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"

	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type GetEpicQueryHandler struct {
	epicRepository    repository.EpicRepository
	commentRepository repository.CommentRepository
}

func NewGetEpicQueryHandler(epicRepository repository.EpicRepository, commentRepository repository.CommentRepository) *GetEpicQueryHandler {
	return &GetEpicQueryHandler{epicRepository: epicRepository, commentRepository: commentRepository}
}

func (c *GetEpicQueryHandler) Handle(ctx context.Context, query *GetEpicQuery) (*model_shared.ResultWithValue[model.Epic], error) {

	resultEpic := c.epicRepository.GetEpic(query.ResourceURL)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Epic](model_shared.NewError("ERROR_EPIC", "Failure getting epic.")), nil
	}

	resultComments := c.commentRepository.GetComments(query.ResourceURL)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Epic](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	epic := resultEpic.Result()
	epic.Comments = resultComments.Result()
	return model_shared.NewResultWithValueSuccess[model.Epic](resultEpic.Result()), nil
}

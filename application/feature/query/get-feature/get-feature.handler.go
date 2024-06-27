package application_getfeature

import (
	"context"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type GetFeatureQueryHandler struct {
	epicRepository    repository.EpicRepository
	commentRepository repository.CommentRepository
	featureRepository repository.FeatureRepository
}

func NewGetFeatureQueryHandler(
	epicRepository repository.EpicRepository,
	commentRepository repository.CommentRepository,
	featureRepository repository.FeatureRepository,
) *GetFeatureQueryHandler {
	return &GetFeatureQueryHandler{
		epicRepository:    epicRepository,
		commentRepository: commentRepository,
		featureRepository: featureRepository,
	}
}

func (c *GetFeatureQueryHandler) Handle(ctx context.Context, query *GetFeatureQuery) (*model_shared.ResultWithValue[model.Feature], error) {
	resultFeature, parentURL := c.featureRepository.GetFeature(query.ResourceURL)
	if !resultFeature.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("ERROR_FEATURE", "Failure getting feature.")), nil
	}
	if parentURL == "" {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("ERROR_FEATURE", "Feature doesn't have a parent.")), nil
	}
	feature := resultFeature.Result()
	resultComments := c.commentRepository.GetComments(query.ResourceURL)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	feature.Comments = resultComments.Result()

	//Epic
	resultEpic := c.epicRepository.GetEpic(parentURL)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("ERROR_EPIC", "Failure getting epic.")), nil
	}
	if resultEpic.Result().WorkItemType != model_shared.EpicType {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("ERROR_EPIC", "Parent is not epic.")), nil
	}
	resultCommentsEpic := c.commentRepository.GetComments(parentURL)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Feature](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	feature.ParentEpic = resultEpic.Result()
	feature.ParentEpic.Comments = resultCommentsEpic.Result()
	return model_shared.NewResultWithValueSuccess[model.Feature](feature), nil
}

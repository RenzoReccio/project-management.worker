package application_getuserstory

import (
	"context"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type GetUserStoryQueryHandler struct {
	epicRepository      repository.EpicRepository
	commentRepository   repository.CommentRepository
	featureRepository   repository.FeatureRepository
	userStoryRepository repository.UserStoryRepository
}

func NewGetUserStoryQueryHandler(
	epicRepository repository.EpicRepository,
	commentRepository repository.CommentRepository,
	featureRepository repository.FeatureRepository,
	userStoryRepository repository.UserStoryRepository,
) *GetUserStoryQueryHandler {
	return &GetUserStoryQueryHandler{
		epicRepository:      epicRepository,
		commentRepository:   commentRepository,
		featureRepository:   featureRepository,
		userStoryRepository: userStoryRepository,
	}
}

func (c *GetUserStoryQueryHandler) Handle(ctx context.Context, query *GetUserStoryQuery) (*model_shared.ResultWithValue[model.UserStory], error) {
	resultUserStory, parentURL := c.userStoryRepository.GetUserStory(query.ResourceURL)
	if !resultUserStory.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_USER_STORY", "Failure getting user story.")), nil
	}
	if parentURL == "" {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_USER_STORY", "User story doesn't have a parent.")), nil
	}
	userStory := resultUserStory.Result()
	resultComments := c.commentRepository.GetComments(query.ResourceURL)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	userStory.Comments = resultComments.Result()

	//Feature
	resultFeature, parentURLFeature := c.featureRepository.GetFeature(parentURL)
	if !resultFeature.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_FEATURE", "Failure getting feature.")), nil
	}
	if parentURLFeature == "" {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_FEATURE", "Feature doesn't have a parent.")), nil
	}
	if resultFeature.Result().WorkItemType != model_shared.FeatureType {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_FEATURE", "Parent for User story is not feature.")), nil
	}

	feature := resultFeature.Result()
	resultCommentsFeature := c.commentRepository.GetComments(parentURL)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	feature.Comments = resultCommentsFeature.Result()

	//Epic
	resultEpic := c.epicRepository.GetEpic(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_EPIC", "Failure getting epic.")), nil
	}
	if resultEpic.Result().WorkItemType != model_shared.EpicType {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_EPIC", "Parent is not epic.")), nil
	}
	resultCommentsEpic := c.commentRepository.GetComments(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	feature.ParentEpic = resultEpic.Result()
	feature.ParentEpic.Comments = resultCommentsEpic.Result()

	userStory.FeatureParent = feature

	return model_shared.NewResultWithValueSuccess[model.UserStory](userStory), nil
}

package application_getuserstory

import (
	"context"
	"fmt"

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
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_USER_STORY",
			fmt.Sprintf("Failure getting user story %s.", query.ResourceURL),
		)), nil
	}
	if parentURL == "" {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_USER_STORY",
			fmt.Sprintf("User story doesn't have a parent %s.", query.ResourceURL),
		)), nil
	}
	userStory := resultUserStory.Result()
	resultComments := c.commentRepository.GetComments(query.ResourceURL)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", query.ResourceURL),
		)), nil
	}
	userStory.Comments = resultComments.Result()

	//Feature
	resultFeature, parentURLFeature := c.featureRepository.GetFeature(parentURL)
	if !resultFeature.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_FEATURE",
			fmt.Sprintf("Failure getting feature %s.", parentURL),
		)), nil
	}
	if parentURLFeature == "" {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_FEATURE",
			fmt.Sprintf("Feature doesn't have a parent %s.", parentURL),
		)), nil
	}
	if resultFeature.Result().WorkItemType != model_shared.FeatureType {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_FEATURE",
			fmt.Sprintf("Parent for User story is not feature. %s.", parentURL),
		)), nil
	}

	feature := resultFeature.Result()
	resultCommentsFeature := c.commentRepository.GetComments(parentURL)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", parentURL),
		)), nil
	}
	feature.Comments = resultCommentsFeature.Result()

	//Epic
	resultEpic := c.epicRepository.GetEpic(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_EPIC",
			fmt.Sprintf("Failure getting epic %s.", parentURLFeature),
		)), nil
	}
	if resultEpic.Result().WorkItemType != model_shared.EpicType {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_EPIC",
			fmt.Sprintf("Parent is not epic %s.", parentURLFeature),
		)), nil
	}
	resultCommentsEpic := c.commentRepository.GetComments(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.UserStory](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", parentURLFeature),
		)), nil
	}
	feature.ParentEpic = resultEpic.Result()
	feature.ParentEpic.Comments = resultCommentsEpic.Result()

	userStory.FeatureParent = feature

	return model_shared.NewResultWithValueSuccess[model.UserStory](userStory), nil
}

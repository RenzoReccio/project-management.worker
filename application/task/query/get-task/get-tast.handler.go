package application_gettask

import (
	"context"
	"fmt"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type GetTaskQueryHandler struct {
	epicRepository      repository.EpicRepository
	commentRepository   repository.CommentRepository
	featureRepository   repository.FeatureRepository
	userStoryRepository repository.UserStoryRepository
	taskRepository      repository.TaskRepository
}

func NewGetTaskQueryHandler(
	epicRepository repository.EpicRepository,
	commentRepository repository.CommentRepository,
	featureRepository repository.FeatureRepository,
	userStoryRepository repository.UserStoryRepository,
	taskRepository repository.TaskRepository) *GetTaskQueryHandler {
	return &GetTaskQueryHandler{
		epicRepository:      epicRepository,
		commentRepository:   commentRepository,
		featureRepository:   featureRepository,
		userStoryRepository: userStoryRepository,
		taskRepository:      taskRepository,
	}
}

func (c *GetTaskQueryHandler) Handle(ctx context.Context, query *GetTaskQuery) (*model_shared.ResultWithValue[model.Task], error) {
	//Task
	resultTask, parentURLTask := c.taskRepository.GetTask(query.ResourceURL)
	if !resultTask.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_TASK",
			fmt.Sprintf("Failure getting task %s.", query.ResourceURL),
		)), nil
	}
	if parentURLTask == "" {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_TASK",
			fmt.Sprintf("Task doesn't have a parent %s.", query.ResourceURL),
		)), nil
	}
	task := resultTask.Result()
	resultCommentsTask := c.commentRepository.GetComments(query.ResourceURL)
	if !resultCommentsTask.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", query.ResourceURL),
		)), nil
	}
	task.Comments = resultCommentsTask.Result()

	//User story
	resultUserStory, parentURLUserStory := c.userStoryRepository.GetUserStory(parentURLTask)
	if !resultUserStory.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_USER_STORY",
			fmt.Sprintf("Failure getting user story %s.", parentURLTask),
		)), nil
	}
	if resultUserStory.Result().WorkItemType != model_shared.UserStoryType {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_USER_STORY",
			fmt.Sprintf("Parent for task is not user story %s.", parentURLTask),
		)), nil
	}
	if parentURLUserStory == "" {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_USER_STORY",
			fmt.Sprintf("User story doesn't have a parent %s.", parentURLTask),
		)), nil
	}
	userStory := resultUserStory.Result()
	resultComments := c.commentRepository.GetComments(parentURLTask)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", parentURLTask),
		)), nil
	}
	userStory.Comments = resultComments.Result()

	//Feature
	resultFeature, parentURLFeature := c.featureRepository.GetFeature(parentURLUserStory)
	if !resultFeature.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_FEATURE",
			fmt.Sprintf("Failure getting feature %s.", parentURLUserStory),
		)), nil
	}
	if resultFeature.Result().WorkItemType != model_shared.FeatureType {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_FEATURE",
			fmt.Sprintf("Parent for user story is not feature %s.", parentURLUserStory),
		)), nil
	}
	if parentURLFeature == "" {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_FEATURE",
			fmt.Sprintf("Feature doesn't have a parent %s.", parentURLUserStory),
		)), nil
	}
	feature := resultFeature.Result()
	resultCommentsFeature := c.commentRepository.GetComments(parentURLUserStory)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", parentURLUserStory),
		)), nil
	}
	feature.Comments = resultCommentsFeature.Result()

	//Epic
	resultEpic := c.epicRepository.GetEpic(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_EPIC",
			fmt.Sprintf("Failure getting epic %s.", parentURLFeature),
		)), nil
	}
	if resultEpic.Result().WorkItemType != model_shared.EpicType {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_EPIC",
			fmt.Sprintf("Parent for feature is not epic %s.", parentURLFeature),
		)), nil
	}
	resultCommentsEpic := c.commentRepository.GetComments(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError(
			"ERROR_COMMENTS",
			fmt.Sprintf("Failure getting comments %s.", parentURLFeature),
		)), nil
	}
	feature.ParentEpic = resultEpic.Result()
	feature.ParentEpic.Comments = resultCommentsEpic.Result()

	userStory.FeatureParent = feature
	task.UserStoryParent = userStory
	return model_shared.NewResultWithValueSuccess[model.Task](task), nil
}

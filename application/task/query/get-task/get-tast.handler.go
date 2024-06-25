package application_gettask

import (
	"context"

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
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_TASK", "Failure getting task.")), nil
	}
	if parentURLTask == "" {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_TASK", "Task doesn't have a parent.")), nil
	}
	task := resultTask.Result()
	resultCommentsTask := c.commentRepository.GetComments(query.ResourceURL)
	if !resultCommentsTask.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	task.Comments = resultCommentsTask.Result()

	//User story
	resultUserStory, parentURLUserStory := c.userStoryRepository.GetUserStory(parentURLTask)
	if !resultUserStory.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_USER_STORY", "Failure getting user story.")), nil
	}
	if resultUserStory.Result().WorkItemType != model_shared.UserStoryType {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_USER_STORY", "Parent for task is not User story")), nil
	}
	if parentURLUserStory == "" {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_USER_STORY", "User story doesn't have a parent.")), nil
	}
	userStory := resultUserStory.Result()
	resultComments := c.commentRepository.GetComments(parentURLTask)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	userStory.Comments = resultComments.Result()

	//Feature
	resultFeature, parentURLFeature := c.featureRepository.GetFeature(parentURLUserStory)
	if !resultFeature.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_FEATURE", "Failure getting feature.")), nil
	}
	if resultFeature.Result().WorkItemType != model_shared.FeatureType {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_FEATURE", "Parent for user story is not feature.")), nil
	}
	if parentURLFeature == "" {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_FEATURE", "Feature doesn't have a parent.")), nil
	}
	feature := resultFeature.Result()
	resultCommentsFeature := c.commentRepository.GetComments(parentURLUserStory)
	if !resultComments.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	feature.Comments = resultCommentsFeature.Result()

	//Epic
	resultEpic := c.epicRepository.GetEpic(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_EPIC", "Failure getting epic.")), nil
	}
	if resultEpic.Result().WorkItemType != model_shared.EpicType {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_EPIC", "Parent for feature is not epic.")), nil
	}
	resultCommentsEpic := c.commentRepository.GetComments(parentURLFeature)
	if !resultEpic.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Task](model_shared.NewError("ERROR_COMMENTS", "Failure getting comments.")), nil
	}
	feature.ParentEpic = resultEpic.Result()
	feature.ParentEpic.Comments = resultCommentsEpic.Result()

	userStory.FeatureParent = feature
	task.UserStoryParent = userStory
	return model_shared.NewResultWithValueSuccess[model.Task](task), nil
}

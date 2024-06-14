package application_getworkitemtype

import (
	"context"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"

	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type GetWorkItemTypeQueryHandler struct {
	workItemTypeRepository repository.WorkItemTypeRepository
}

func NewGetWorkItemTypeQueryHandler(workItemTypeRepository repository.WorkItemTypeRepository) *GetWorkItemTypeQueryHandler {
	return &GetWorkItemTypeQueryHandler{workItemTypeRepository: workItemTypeRepository}
}

func (c *GetWorkItemTypeQueryHandler) Handle(ctx context.Context, query *GetWorkItemTypeQuery) (*model_shared.ResultWithValue[model.WorkItemType], error) {

	resultWorkItemType := c.workItemTypeRepository.GetWorkItemType(&query.ResourceURL)
	if !resultWorkItemType.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.WorkItemType](model_shared.NewError("ERROR_WORK_ITEM", "Failure getting work item type.")), nil
	}

	return model_shared.NewResultWithValueSuccess[model.WorkItemType](model.NewWorkItemType(resultWorkItemType.Result())), nil
}

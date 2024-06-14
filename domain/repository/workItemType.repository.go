package repository

import (
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type WorkItemTypeRepository interface {
	GetWorkItemType(url *string) *model_shared.ResultWithValue[string]
}

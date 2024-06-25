package repository

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

type FeatureRepository interface {
	GetFeature(url string) (*model_shared.ResultWithValue[model.Feature], string)
}

package application_sendfeature

import "github.com/RenzoReccio/project-management.worker/domain/model"

type SendFeatureEvent struct {
	Data *model.Feature
}

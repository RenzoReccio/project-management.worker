package application_sendepic

import "github.com/RenzoReccio/project-management.worker/domain/model"

type SendEpicEvent struct {
	Data *model.Epic
}

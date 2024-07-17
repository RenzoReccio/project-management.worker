package repository

import (
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
)

var EventLogger EventLogRepository

type EventLogRepository interface {
	InsertLog(resourceURL string, message string) *model_shared.Result
	InsertErrorLog(resourceURL string, message string) *model_shared.Result
}

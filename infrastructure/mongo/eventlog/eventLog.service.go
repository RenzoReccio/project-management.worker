package mongoInfraestructure_eventlog

import (
	"context"
	"runtime/debug"

	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventLogService struct {
	databaseConnection *mongo.Database
}

var EventLogLogger *EventLogService

func NewEventLogService(_databaseConnection *mongo.Database) repository.EventLogRepository {
	return &EventLogService{
		databaseConnection: _databaseConnection,
	}
}

func (u *EventLogService) InsertLog(resourceURL string, message string) *model_shared.Result {
	eventLog := NewEventLog("INFO", resourceURL, message, string(debug.Stack()))
	_, err := u.databaseConnection.Collection("eventLog").InsertOne(context.TODO(), eventLog)
	if err != nil {
		return model_shared.NewResultFailure(model_shared.NewError("DATABASE_FAIL", err.Error()))
	}
	return model_shared.NewResultSuccess()
}

func (u *EventLogService) InsertErrorLog(resourceURL string, message string) *model_shared.Result {
	eventLog := NewEventLog("ERROR", resourceURL, message, string(debug.Stack()))
	_, err := u.databaseConnection.Collection("eventLog").InsertOne(context.TODO(), eventLog)
	if err != nil {
		return model_shared.NewResultFailure(model_shared.NewError("DATABASE_FAIL", err.Error()))
	}
	return model_shared.NewResultSuccess()
}

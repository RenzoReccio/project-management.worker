package mongoInfraestructure

import (
	"context"
	"time"

	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventService struct {
	databaseConnection *mongo.Database
}

func NewEventService(_databaseConnection *mongo.Database) repository.EventRepository {
	return &EventService{
		databaseConnection: _databaseConnection,
	}
}

func (u *EventService) InsertEvent(in *model.Event) *model_shared.ResultWithValue[model.Event] {
	event := NewEvent(in.EventID, in.SubscriptionID, in.EventType, in.CreatedDate, in.ResourceId, in.ResourceUrl, in.Processed)

	insertEvent := struct {
		EventID        string    `bson:"EventID"`
		SubscriptionID string    `bson:"SubscriptionID"`
		EventType      string    `bson:"EventType"`
		CreatedDate    time.Time `bson:"CreatedDate"`
		ResourceId     int       `bson:"ResourceId"`
		ResourceUrl    string    `bson:"ResourceUrl"`
		Processed      bool      `bson:"Processed"`
	}{
		EventID:        event.EventID,
		SubscriptionID: event.SubscriptionID,
		EventType:      event.EventType,
		CreatedDate:    event.CreatedDate,
		ResourceId:     event.ResourceId,
		ResourceUrl:    event.ResourceUrl,
		Processed:      event.Processed,
	}

	result, err := u.databaseConnection.Collection("event").InsertOne(context.TODO(), insertEvent)
	if err != nil {
		return model_shared.NewResultWithValueFailure[model.Event](model_shared.NewError("DATABASE_FAIL", err.Error()))
	}

	in.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return model_shared.NewResultWithValueSuccess[model.Event](in)
}

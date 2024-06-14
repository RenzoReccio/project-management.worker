package mongoInfraestructure

import "time"

type Event struct {
	Id             string    `bson:"_id"`
	EventID        string    `bson:"EventID"`
	SubscriptionID string    `bson:"SubscriptionID"`
	EventType      string    `bson:"EventType"`
	CreatedDate    time.Time `bson:"CreatedDate"`
	ResourceId     int       `bson:"ResourceId"`
	ResourceUrl    string    `bson:"ResourceUrl"`
	Processed      bool      `bson:"Processed"`
}

func NewEvent(eventId string, subscriptionId string, eventType string, createdDate time.Time, resourceId int, resourceUrl string, processed bool) *Event {
	return &Event{
		EventID:        eventId,
		SubscriptionID: subscriptionId,
		EventType:      eventType,
		CreatedDate:    createdDate,
		ResourceId:     resourceId,
		ResourceUrl:    resourceUrl,
		Processed:      processed}
}

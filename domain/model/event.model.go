package model

import "time"

type Event struct {
	Id             string
	EventID        string
	SubscriptionID string
	EventType      string
	CreatedDate    time.Time
	ResourceId     int
	ResourceUrl    string
	Processed      bool
}

func NewEvent(eventId string, subscriptionId string, eventType string, createdDate time.Time, resourceId int, resourceUrl string) *Event {
	return &Event{
		EventID:        eventId,
		SubscriptionID: subscriptionId,
		EventType:      eventType,
		CreatedDate:    createdDate,
		ResourceId:     resourceId,
		ResourceUrl:    resourceUrl,
		Processed:      false}
}

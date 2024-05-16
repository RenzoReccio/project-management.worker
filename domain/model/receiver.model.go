package model

import "time"

type Receiver struct {
	ID             string    `json:"id"`
	SubscriptionID string    `json:"subscriptionId"`
	EventType      string    `json:"eventType"`
	CreatedDate    time.Time `json:"createdDate"`
	ResourceId     int       `json:"resourceId"`
	ResourceUrl    string    `json:"resourceUrl"`
	Processed      bool      `json:"processed"`
}

func NewReceiver(id string, subscriptionId string, eventType string, createdDate time.Time, resourceId int, resourceUrl string) *Receiver {
	return &Receiver{ID: id,
		SubscriptionID: subscriptionId,
		EventType:      eventType,
		CreatedDate:    createdDate,
		ResourceId:     resourceId,
		ResourceUrl:    resourceUrl,
		Processed:      false}
}

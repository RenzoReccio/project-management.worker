package applicationreceiver

import "time"

type InsertReceiverDto struct {
	SubscriptionID  string      `json:"subscriptionId"`
	NotificationID  int         `json:"notificationId"`
	ID              string      `json:"id"`
	EventType       string      `json:"eventType"`
	PublisherID     string      `json:"publisherId"`
	Message         interface{} `json:"message"`
	DetailedMessage interface{} `json:"detailedMessage"`
	Resource        struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"resource"`
	ResourceVersion    string `json:"resourceVersion"`
	ResourceContainers struct {
		Collection struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"collection"`
		Account struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"account"`
		Project struct {
			ID      string `json:"id"`
			BaseURL string `json:"baseUrl"`
		} `json:"project"`
	} `json:"resourceContainers"`
	CreatedDate time.Time `json:"createdDate"`
}

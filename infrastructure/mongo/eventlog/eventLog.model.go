package mongoInfraestructure_eventlog

import "time"

type EventLog struct {
	LogType     string    `bson:"LogType"`
	AppName     string    `bson:"AppName"`
	CreatedDate time.Time `bson:"CreatedDate"`
	ResourceUrl string    `bson:"ResourceUrl"`
	Message     string    `bson:"Message"`
	StackTrace  string    `bson:"StackTrace"`
}

func NewEventLog(logType string, resourceUrl string, message string, stackTrace string) *EventLog {
	return &EventLog{
		AppName:     "WORKER",
		CreatedDate: time.Now(),
		ResourceUrl: resourceUrl,
		LogType:     logType,
		Message:     message,
		StackTrace:  stackTrace,
	}
}

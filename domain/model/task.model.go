package model

import "time"

type Comment struct {
	Date           time.Time
	User           string
	UserUniqueName string
}

type Task struct {
	Id                         int
	AreaPath                   string
	TeamProject                string
	IterationPath              string
	WorkItemType               string
	State                      string
	Reason                     string
	CreatedDate                time.Time
	CreatedByUser              string
	CreatedByUserUniqueName    string
	ModidifiedDate             time.Time
	ModidifiedByUser           string
	ModidifiedByUserUniqueName string
	Title                      string
	Comments                   []Comment
}

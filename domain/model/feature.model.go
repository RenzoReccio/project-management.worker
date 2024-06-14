package model

import "time"

type Feature struct {
	Id              int
	AreaPath        string
	TeamProject     string
	IterationPath   string
	WorkItemType    string
	State           string
	Reason          string
	AssignedTo      *Person
	CreatedDate     time.Time
	Title           string
	Priority        string
	ValueArea       string
	Risk            string
	TargetDate      time.Time
	BusinessValue   int
	TimeCriticality int
	Effort          int
	StartDate       time.Time
	Description     string
	Tags            string
	Url             string
	ParentEpic      *Epic
	Comments        *[]Comment
}

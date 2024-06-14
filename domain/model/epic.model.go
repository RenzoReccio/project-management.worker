package model

import "time"

type Epic struct {
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
	Description     string
	Priority        int
	ValueArea       string
	Risk            string
	BusinessValue   int
	TimeCriticality int
	Effort          int
	StartDate       time.Time
	TargetDate      time.Time
	Url             string
	PageUrl         string
	Tags            string
	Comments        *[]Comment
}

func NewEpic(Id int,
	AreaPath string,
	TeamProject string,
	IterationPath string,
	WorkItemType string,
	State string,
	Reason string,
	AssignedTo *Person,
	CreatedDate time.Time,
	Title string,
	Description string,
	Priority int,
	ValueArea string,
	Risk string,
	BusinessValue int,
	TimeCriticality int,
	Effort int,
	StartDate time.Time,
	TargetDate time.Time,
	Url string,
	Tags string,
	Comments *[]Comment,
	PageURL string) *Epic {
	return &Epic{
		Id:              Id,
		AreaPath:        AreaPath,
		TeamProject:     TeamProject,
		IterationPath:   IterationPath,
		WorkItemType:    WorkItemType,
		State:           State,
		Reason:          Reason,
		AssignedTo:      AssignedTo,
		CreatedDate:     CreatedDate,
		Title:           Title,
		Description:     Description,
		Priority:        Priority,
		ValueArea:       ValueArea,
		Risk:            Risk,
		BusinessValue:   BusinessValue,
		TimeCriticality: TimeCriticality,
		Effort:          Effort,
		StartDate:       StartDate,
		TargetDate:      TargetDate,
		Url:             Url,
		Tags:            Tags,
		Comments:        Comments,
		PageUrl:         PageURL,
	}
}

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
	Priority        int
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
	PageUrl         string
}

func NewFeature(Id int,
	AreaPath string,
	TeamProject string,
	IterationPath string,
	WorkItemType string,
	State string,
	Reason string,
	AssignedTo *Person,
	CreatedDate time.Time,
	Title string,
	Priority int,
	ValueArea string,
	Risk string,
	TargetDate time.Time,
	BusinessValue int,
	TimeCriticality int,
	Effort int,
	StartDate time.Time,
	Description string,
	Tags string,
	Url string,
	ParentEpic *Epic,
	Comments *[]Comment,
	PageUrl string) *Feature {
	return &Feature{
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
		Priority:        Priority,
		ValueArea:       ValueArea,
		Risk:            Risk,
		TargetDate:      TargetDate,
		BusinessValue:   BusinessValue,
		TimeCriticality: TimeCriticality,
		Effort:          Effort,
		StartDate:       StartDate,
		Description:     Description,
		Tags:            Tags,
		Url:             Url,
		ParentEpic:      ParentEpic,
		Comments:        Comments,
		PageUrl:         PageUrl,
	}
}

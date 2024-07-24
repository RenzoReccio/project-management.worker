package model

import "time"

type Task struct {
	Id               int
	AreaPath         string
	TeamProject      string
	IterationPath    string
	WorkItemType     string
	State            string
	Reason           string
	AssignedTo       *Person
	Title            string
	RemainingWork    int
	OriginalEstimate int
	CompletedWork    int
	Activity         string
	Priority         int
	Description      string
	Tags             string
	UserStoryParent  *UserStory
	Url              string
	Comments         *[]Comment
	PageUrl          string
	CreatedDate      time.Time
	UpdatedDate      time.Time
}

func NewTask(
	Id int,
	AreaPath string,
	TeamProject string,
	IterationPath string,
	WorkItemType string,
	State string,
	Reason string,
	AssignedTo *Person,
	Title string,
	RemainingWork int,
	OriginalEstimate int,
	CompletedWork int,
	Activity string,
	Priority int,
	Description string,
	Tags string,
	UserStoryParent *UserStory,
	Url string,
	Comments *[]Comment,
	PageUrl string,
	CreatedDate time.Time,
	UpdatedDate time.Time,
) *Task {
	return &Task{
		Id:               Id,
		AreaPath:         AreaPath,
		TeamProject:      TeamProject,
		IterationPath:    IterationPath,
		WorkItemType:     WorkItemType,
		State:            State,
		Reason:           Reason,
		AssignedTo:       AssignedTo,
		Title:            Title,
		RemainingWork:    RemainingWork,
		OriginalEstimate: OriginalEstimate,
		CompletedWork:    CompletedWork,
		Activity:         Activity,
		Priority:         Priority,
		Description:      Description,
		Tags:             Tags,
		UserStoryParent:  UserStoryParent,
		Url:              Url,
		Comments:         Comments,
		PageUrl:          PageUrl,
		CreatedDate:      CreatedDate,
		UpdatedDate:      UpdatedDate,
	}
}

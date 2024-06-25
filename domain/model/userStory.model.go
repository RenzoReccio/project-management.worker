package model

type UserStory struct {
	Id                 int
	AreaPath           string
	TeamProject        string
	IterationPath      string
	WorkItemType       string
	State              string
	Reason             string
	AssignedTo         *Person
	Title              string
	BoardColumn        string
	BoardColumnDone    bool
	Priority           int
	ValueArea          string
	Risk               string
	KanbanColumn       string
	KanbanColumnDone   bool
	Description        string
	AcceptanceCriteria string
	Tags               string
	FeatureParent      *Feature
	Url                string
	Comments           *[]Comment
}

func NewUserStory(
	Id int,
	AreaPath string,
	TeamProject string,
	IterationPath string,
	WorkItemType string,
	State string,
	Reason string,
	AssignedTo *Person,
	Title string,
	BoardColumn string,
	BoardColumnDone bool,
	Priority int,
	ValueArea string,
	Risk string,
	KanbanColumn string,
	KanbanColumnDone bool,
	Description string,
	AcceptanceCriteria string,
	Tags string,
	FeatureParent *Feature,
	Url string,
	Comments *[]Comment,
) *UserStory {
	return &UserStory{
		Id:                 Id,
		AreaPath:           AreaPath,
		TeamProject:        TeamProject,
		IterationPath:      IterationPath,
		WorkItemType:       WorkItemType,
		State:              State,
		Reason:             Reason,
		AssignedTo:         AssignedTo,
		Title:              Title,
		BoardColumn:        BoardColumn,
		BoardColumnDone:    BoardColumnDone,
		Priority:           Priority,
		ValueArea:          ValueArea,
		Risk:               Risk,
		KanbanColumn:       KanbanColumn,
		KanbanColumnDone:   KanbanColumnDone,
		Description:        Description,
		AcceptanceCriteria: AcceptanceCriteria,
		Tags:               Tags,
		FeatureParent:      FeatureParent,
		Url:                Url,
		Comments:           Comments,
	}
}

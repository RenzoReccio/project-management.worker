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

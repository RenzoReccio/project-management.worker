package model

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
	Priority         int
	Description      string
	Tags             string
	UserStoryParent  *UserStory
	Url              string
	Comments         *[]Comment
}

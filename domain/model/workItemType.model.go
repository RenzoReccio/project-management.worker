package model

type WorkItemType struct {
	Type *string
}

func NewWorkItemType(typeWork *string) *WorkItemType {
	return &WorkItemType{
		Type: typeWork,
	}
}

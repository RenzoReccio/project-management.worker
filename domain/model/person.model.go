package model

type Person struct {
	DisplayName string
	Id          string
	UniqueName  string
}

func NewPerson(DisplayName string, Id string, UniqueName string) *Person {
	return &Person{
		DisplayName: DisplayName,
		Id:          Id,
		UniqueName:  UniqueName,
	}
}

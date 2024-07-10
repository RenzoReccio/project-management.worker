package model

import "time"

type Comment struct {
	Date      time.Time
	Text      string
	CreatedBy *Person
}

func NewComment(date time.Time, CreatedBy *Person, text string) *Comment {
	return &Comment{
		Date:      date,
		CreatedBy: CreatedBy,
		Text:      text,
	}
}

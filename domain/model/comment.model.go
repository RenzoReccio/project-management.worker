package model

import "time"

type Comment struct {
	Date           time.Time
	User           string
	UserUniqueName string
	Text           string
}

func NewComment(date time.Time, user string, userUniqueName string, text string) *Comment {
	return &Comment{
		Date:           date,
		User:           user,
		UserUniqueName: userUniqueName,
		Text:           text,
	}
}

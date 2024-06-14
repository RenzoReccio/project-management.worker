package model_shared

import (
	"fmt"
)

// Error struct to represent an error with a code and description
type Error struct {
	Code        string
	Description string
}

var (
	None      = Error{Code: "", Description: ""}
	NullValue = Error{Code: "Error.NullValue", Description: "Null value was provided"}
)

func NullOrEmpty(typeName string, attribute string) *Error {
	return &Error{
		Code:        fmt.Sprintf("%s.NullOrEmpty", typeName),
		Description: fmt.Sprintf("Null or Empty attribute: %s", attribute),
	}
}

func NewError(code string, description string) *Error {
	return &Error{
		Code:        code,
		Description: description,
	}
}

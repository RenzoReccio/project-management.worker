package model_shared

type ResultWithValue[TValue any] struct {
	IsSuccess bool
	Error     *Error
	value     *TValue
}

type Result struct {
	IsSuccess bool
	Error     *Error
}

func NewResultSuccess() *Result {
	return &Result{
		IsSuccess: true,
		Error:     nil,
	}
}

func NewResultWithValueSuccess[TValue any](value *TValue) *ResultWithValue[TValue] {
	return &ResultWithValue[TValue]{
		IsSuccess: true,
		Error:     &None,
		value:     value,
	}
}

func NewResultFailure(error *Error) *Result {
	return &Result{
		IsSuccess: false,
		Error:     error,
	}
}

func NewResultWithValueFailure[TValue any](error *Error) *ResultWithValue[TValue] {
	return &ResultWithValue[TValue]{
		IsSuccess: false,
		Error:     error,
		value:     nil,
	}
}

func (c ResultWithValue[TValue]) Result() *TValue {
	return c.value
}

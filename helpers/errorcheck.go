package helpers

import (
	"fmt"
)

type err struct {
	Raw      error
	Addition string
}

func NewError(e error) *err {
	return &err{
		Raw:      e,
		Addition: "",
	}
}

func NewErrorWithAddition(e error, addition string) *err {
	error := NewError(e)
	error.Addition = addition
	return error
}

func (e *err) SetAddition(a string) *err {
	e.Addition = a
	return e
}

func (e *err) String() string {
	var str string
	str = fmt.Sprintf("Error: %+v", e.Raw)
	if e.Addition != "" {
		str += fmt.Sprintf("\n\tAdditionalMessage: %s", e.Addition)
	}
	return str
}

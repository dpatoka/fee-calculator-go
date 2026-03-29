package errors

import "fmt"

type UnsupportedTermError struct {
	term int
}

func NewUnsupportedTermError(term int) *UnsupportedTermError {
	return &UnsupportedTermError{term}
}

func (e *UnsupportedTermError) Error() string {
	return fmt.Sprintf("term %d not supported", e.term)
}

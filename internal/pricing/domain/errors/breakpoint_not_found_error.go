package errors

import "fmt"

type BreakpointNotFoundError struct {
	amount float64
	msg    string
}

func NewLowerBreakpointNotFoundError(amount float64) error {
	return &BreakpointNotFoundError{
		amount,
		fmt.Sprintf("lower breakpoint not found for amount %.2f", amount),
	}
}

func NewUpperBreakpointNotFoundError(amount float64) error {
	return &BreakpointNotFoundError{
		amount,
		fmt.Sprintf("upper breakpoint not found for amount %.2f", amount),
	}
}

func NewNoBreakpointsError() error {
	return &BreakpointNotFoundError{
		0,
		"Breakpoint repository is empty",
	}
}

func (e *BreakpointNotFoundError) Error() string {
	return e.msg
}

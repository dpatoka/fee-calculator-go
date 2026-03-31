package errors

import "fmt"

type BreakpointNotFoundError struct {
	amount float64
	msg    string
}

func NewLowerBreakpointNotFoundError(amount float64) *BreakpointNotFoundError {
	return &BreakpointNotFoundError{
		amount,
		fmt.Sprintf("lower breakpoint not found for amount %.2f", amount),
	}
}

func NewUpperBreakpointNotFoundError(amount float64) *BreakpointNotFoundError {
	return &BreakpointNotFoundError{
		amount,
		fmt.Sprintf("upper breakpoint not found for amount %.2f", amount),
	}
}

func NewNotBreakpointsError() *BreakpointNotFoundError {
	return &BreakpointNotFoundError{
		0,
		fmt.Sprintf("Breakpoint repository is empty"),
	}
}

func (e *BreakpointNotFoundError) Error() string {
	return e.msg
}

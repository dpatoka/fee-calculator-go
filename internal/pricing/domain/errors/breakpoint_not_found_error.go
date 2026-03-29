package errors

import "fmt"

type BreakpointNotFountError struct {
	amount float64
	msg    string
}

func NewLowerBreakpointNotFountError(amount float64) *BreakpointNotFountError {
	return &BreakpointNotFountError{
		amount,
		fmt.Sprintf("lower breakpoint not found for amount %.2f", amount),
	}
}

func NewUpperBreakpointNotFountError(amount float64) *BreakpointNotFountError {
	return &BreakpointNotFountError{
		amount,
		fmt.Sprintf("upper breakpoint not found for amount %.2f", amount),
	}
}

func NewNotBreakpointsError() *BreakpointNotFountError {
	return &BreakpointNotFountError{
		0,
		fmt.Sprintf("Breakpoint repository is empty"),
	}
}

func (e *BreakpointNotFountError) Error() string {
	return e.msg
}

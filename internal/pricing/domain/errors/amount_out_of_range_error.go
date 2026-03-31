package errors

import "fmt"

type AmountOutOfRangeError struct {
	RequestedAmount float64
	Boundary        float64
	Direction       string
}

func (e *AmountOutOfRangeError) Error() string {
	return fmt.Sprintf(
		"requested amount %.2f is %s boundary %.2f",
		e.RequestedAmount, e.Direction, e.Boundary,
	)
}

func ErrorAmountBelowLowerBoundary(requested, lower float64) error {
	return &AmountOutOfRangeError{
		requested,
		lower,
		"below lower",
	}
}

func ErrorAmountAboveUpperBoundary(requested, upper float64) error {
	return &AmountOutOfRangeError{
		requested,
		upper,
		"above upper",
	}
}

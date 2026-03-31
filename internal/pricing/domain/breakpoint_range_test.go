package domain

import (
	"fee-calculator-go/internal/pricing/domain/errors"
	"testing"

	"gotest.tools/assert"
)

type testCase struct {
	name            string
	lowerAmount     float64
	lowerFee        float64
	upperAmount     float64
	upperFee        float64
	requestedAmount float64
	want            float64
	expectedError   error
}

func TestBoundaries(t *testing.T) {
	tests := []testCase{
		{"at lower breakpoint term 12", 1000, 50, 2000, 90, 1000, 50, nil},
		{"at upper breakpoint term 12", 1000, 50, 2000, 90, 2000, 90, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got, _ := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestInterpolationWithoutRounding(t *testing.T) {
	tests := []testCase{
		{"midpoint interpolation", 1000, 50, 2000, 90, 1500, 70, nil},
		{"quarter point interpolation", 1000, 50, 2000, 90, 1250, 60, nil},
		{"equal fees at both breakpoints", 2000, 90, 3000, 90, 2500, 90, nil},
		{"midpoint interpolation with large amounts", 15000, 600, 20000, 800, 17500, 700, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got, _ := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestInterpolationWithRounding(t *testing.T) {
	tests := []testCase{
		{"basic rounding", 1000, 50, 2000, 90, 1300, 65, nil},
		{"edge case: decreasing fee rounding where upper fee < lower fee", 4000, 115, 5000, 100, 4500, 110, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got, _ := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []testCase{
		{"below lower breakpoint term 12", 1000, 50, 2000, 90, 500, 0, &errors.AmountOutOfRangeError{}},
		{"above upper breakpoint term 12", 1000, 50, 2000, 90, 2500, 0, &errors.AmountOutOfRangeError{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got, _ := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func getBreakPointRange(test testCase) *BreakpointRange {
	return NewBreakpointRange(test.lowerAmount, test.lowerFee, test.upperAmount, test.upperFee)
}

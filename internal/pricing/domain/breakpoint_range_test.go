package domain

import (
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
}

func TestBoundaries(t *testing.T) {
	tests := []testCase{
		{"at lower breakpoint term 12", 1000, 50, 2000, 90, 1000, 50},
		{"at upper breakpoint term 12", 1000, 50, 2000, 90, 2000, 90},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestInterpolationWithoutRounding(t *testing.T) {
	tests := []testCase{
		{"midpoint interpolation", 1000, 50, 2000, 90, 1500, 70},
		{"quarter point interpolation", 1000, 50, 2000, 90, 1250, 60},
		{"equal fees at both breakpoints", 2000, 90, 3000, 90, 2500, 90},
		{"midpoint interpolation with large amounts", 15000, 600, 20000, 800, 17500, 700},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestInterpolationWithRounding(t *testing.T) {
	tests := []testCase{
		{"basic rounding", 1000, 50, 2000, 90, 1300, 65},
		{"edge case: decreasing fee rounding where upper fee < lower fee", 4000, 115, 5000, 100, 4500, 110},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			breakpointRange := getBreakPointRange(test)
			got := breakpointRange.CalculateFee(test.requestedAmount)
			assert.Equal(t, test.want, got)
		})
	}
}

func getBreakPointRange(test testCase) *BreakpointRange {
	return NewBreakpointRange(
		breakpoint{test.lowerAmount, test.lowerFee},
		breakpoint{test.upperAmount, test.upperFee},
	)
}

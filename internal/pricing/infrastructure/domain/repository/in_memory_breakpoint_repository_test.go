package repository

import (
	"fee-calculator-go/internal/pricing/domain"
	"fee-calculator-go/internal/pricing/domain/errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/assert"
)

func TestForCorrectValues(t *testing.T) {
	tests := []struct {
		name               string
		term               int
		requestedAmount    float64
		expectedBreakpoint *domain.BreakpointRange
	}{
		{"term 12 amount 1000 exact match",
			12,
			1000,
			domain.NewBreakpointRange(1000, 50, 2000, 90),
		},
		{"term 12 amount 1500 in first tier",
			12,
			1500,
			domain.NewBreakpointRange(1000, 50, 2000, 90),
		},
		{"term 12 amount 20000 max amount",
			12,
			20000,
			domain.NewBreakpointRange(20000, 400, 20000, 400),
		},
		{"term 24 amount 1000 exact match",
			24,
			1000,
			domain.NewBreakpointRange(1000, 70, 2000, 100),
		},
		{"term 24 amount 20000 max amount",
			24,
			20000,
			domain.NewBreakpointRange(20000, 800, 20000, 800),
		},
	}

	repo := InMemoryBreakpointRepository{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := repo.GetForTermAndAmount(test.term, test.requestedAmount)

			assert.DeepEqual(t, test.expectedBreakpoint, got,
				cmp.Comparer(func(a, b domain.BreakpointRange) bool {
					return cmp.Equal(a, b, cmpopts.EquateComparable(domain.BreakpointRange{}))
				}),
			)
		})
	}
}

func TestForInCorrectValues(t *testing.T) {
	tests := []struct {
		name             string
		term             int
		requestedAmount  float64
		expectedError    error
		expectedErrorMsg string
	}{
		{"term 12 amount 0 below threshold",
			12,
			0.0,
			&errors.AmountOutOfRangeError{},
			"requested amount 0.00 is below lower boundary 1000.00",
		},
		{"term 12 amount negative below threshold",
			12,
			-1.0,
			&errors.AmountOutOfRangeError{},
			"requested amount -1.00 is below lower boundary 1000.00",
		},
		{"term 12 amount 200 below threshold",
			12,
			200.01,
			&errors.AmountOutOfRangeError{},
			"requested amount 200.01 is below lower boundary 1000.00",
		},
		{"term 6 unsupported term",
			6,
			1000.0,
			&errors.UnsupportedTermError{},
			"term 6 not supported",
		},
	}

	repo := InMemoryBreakpointRepository{}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := repo.GetForTermAndAmount(test.term, test.requestedAmount)

			assert.Equal(t, true, got == nil)
			assert.ErrorType(t, err, test.expectedError)
			assert.ErrorContains(t, err, test.expectedErrorMsg)
		})
	}
}

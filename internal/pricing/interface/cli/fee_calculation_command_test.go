package cli_test

import (
	"fee-calculator-go/internal/pricing/interface/cli"
	"testing"

	"gotest.tools/assert"
)

type testCases struct {
	amount      float64
	term        int
	want        string
	description string
}

func TestCalculateFeeWithSuccess(t *testing.T) {
	tests := []testCases{
		{19.25000, 12, " 385.000000", "fee not rounded up, loan + fee already divisible by 5"},
		{11.50000, 24, " 460.00", " fee not rounded up, loan + fee already divisible by 5"},
	}

	command := cli.NewFeeCalculationCommand()

	for _, test := range tests {
		test := test

		t.Run(test.description, func(t *testing.T) {
			got := command.Execute(test.amount, test.term)
			assert.Equal(t, test.want, got)
		})
	}
}

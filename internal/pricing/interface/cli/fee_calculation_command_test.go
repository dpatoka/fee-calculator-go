package cli_test

import (
	"fee-calculator-go/internal/pricing/interface/cli"
	"testing"

	"gotest.tools/assert"
)

type testCase struct {
	amount      float64
	term        int
	want        string
	description string
}

func TestCalculateFeeWithSuccess(t *testing.T) {
	tests := []testCase{
		{19250.00, 12, "385.00", "fee not rounded up, loan + fee already divisible by 5"},
		{11500.00, 24, "460.00", "fee not rounded up, loan + fee already divisible by 5"},
		{1123.00, 12, "57.00", "fee rounded up to make loan + fee divisible by 5"},
		{2567.00, 24, "113.00", "fee rounded up to make loan + fee divisible by 5"},
	}

	command := cli.NewFeeCalculationCommand()

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got, _ := command.Execute(test.amount, test.term)
			assert.Equal(t, test.want, got)
		})
	}
}

package cli

import (
	"fee-calculator-go/internal/pricing/application/query"
	"fmt"
)

type FeeCalculationCommand struct {
	queryHandler *query.FeeCalculationQueryHandler
}

func NewFeeCalculationCommand() *FeeCalculationCommand {
	return &FeeCalculationCommand{
		&query.FeeCalculationQueryHandler{},
	}
}

func (f *FeeCalculationCommand) Execute(amount float64, term int) string {
	q := query.FeeCalculationQuery{amount, term}
	result := f.queryHandler.Run(q)

	return fmt.Sprintf("%f", result)
}

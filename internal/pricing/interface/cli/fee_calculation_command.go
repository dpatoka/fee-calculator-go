package cli

import (
	"fee-calculator-go/internal/pricing/application/query"
	"fee-calculator-go/internal/pricing/infrastructure/domain/repository"
	"fmt"
)

type FeeCalculationCommand struct {
	queryHandler *query.FeeCalculationQueryHandler
}

func NewFeeCalculationCommand() *FeeCalculationCommand {
	queryHandler := query.NewFeeCalculationQueryHandler(
		&repository.InMemoryBreakpointRepository{},
	)

	return &FeeCalculationCommand{
		queryHandler,
	}
}

func (f *FeeCalculationCommand) Execute(amount float64, term int) (string, error) {
	q := query.FeeCalculationQuery{amount, term}
	result, err := f.queryHandler.Run(q)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", result), nil
}

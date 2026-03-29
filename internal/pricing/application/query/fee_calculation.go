package query

import (
	"fee-calculator-go/internal/pricing/domain"
)

type FeeCalculationQuery struct {
	Amount float64
	Term   int
}

type FeeCalculationQueryHandler struct {
	repository domain.BreakpointRepository
}

func NewFeeCalculationQueryHandler(repository domain.BreakpointRepository) *FeeCalculationQueryHandler {
	return &FeeCalculationQueryHandler{repository}
}

func (f *FeeCalculationQueryHandler) Run(query FeeCalculationQuery) (float64, error) {
	breakpoint, err := f.repository.GetForTermAndAmount(query.Term, query.Amount)
	if err != nil {
		return 0, err
	}

	fee, err := breakpoint.CalculateFee(query.Amount)
	if err != nil {
		return 0, err
	}

	return fee, nil
}

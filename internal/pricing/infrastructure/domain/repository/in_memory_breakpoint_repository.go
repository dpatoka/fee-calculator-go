package repository

import (
	"fee-calculator-go/internal/pricing/domain"
)

type InMemoryBreakpointRepository struct {
}

func (r *InMemoryBreakpointRepository) GetForTermAndAmount(term int, amount float64) (*domain.BreakpointRange, error) {
	return &domain.BreakpointRange{}, nil
}

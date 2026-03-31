package repository

import (
	"fee-calculator-go/internal/pricing/domain"
	"fee-calculator-go/internal/pricing/domain/errors"
	"maps"
	"slices"
)

var breakpoints = map[int]map[float64]float64{
	12: {
		1000.0:  50.0,
		2000.0:  90.0,
		3000.0:  90.0,
		4000.0:  115.0,
		5000.0:  100.0,
		6000.0:  120.0,
		7000.0:  140.0,
		8000.0:  160.0,
		9000.0:  180.0,
		10000.0: 200.0,
		11000.0: 220.0,
		12000.0: 240.0,
		13000.0: 260.0,
		14000.0: 280.0,
		15000.0: 300.0,
		16000.0: 320.0,
		17000.0: 340.0,
		18000.0: 360.0,
		19000.0: 380.0,
		20000.0: 400.0,
	},
	24: {
		1000.0:  70.0,
		2000.0:  100.0,
		3000.0:  120.0,
		4000.0:  160.0,
		5000.0:  200.0,
		6000.0:  240.0,
		7000.0:  280.0,
		8000.0:  320.0,
		9000.0:  360.0,
		10000.0: 400.0,
		11000.0: 440.0,
		12000.0: 480.0,
		13000.0: 520.0,
		14000.0: 560.0,
		15000.0: 600.0,
		16000.0: 640.0,
		17000.0: 680.0,
		18000.0: 720.0,
		19000.0: 760.0,
		20000.0: 800.0,
	},
}

type InMemoryBreakpointRepository struct {
}

type breakpoint struct {
	amount float64
	fee    float64
}

func (r *InMemoryBreakpointRepository) GetForTermAndAmount(term int, amount float64) (*domain.BreakpointRange, error) {
	breakpoints, err := getTermBreakpoints(term)
	if err != nil {
		return nil, err
	}

	err = validateAmountBounds(amount, breakpoints)
	if err != nil {
		return nil, err
	}

	lowerBreakpoint, err := findLowerBreakpoint(amount, breakpoints)
	if err != nil {
		return nil, err
	}

	upperBreakpoint, err := findUpperBreakpoint(amount, breakpoints)
	if err != nil {
		return nil, err
	}

	return domain.NewBreakpointRange(
		lowerBreakpoint.amount,
		lowerBreakpoint.fee,
		upperBreakpoint.amount,
		upperBreakpoint.fee,
	), nil
}

func getTermBreakpoints(term int) (map[float64]float64, error) {
	result, exists := breakpoints[term]
	if !exists {
		return nil, errors.NewUnsupportedTermError(term)
	}

	return result, nil
}

func validateAmountBounds(amount float64, termBreakpoints map[float64]float64) error {
	allAmounts := slices.Collect(maps.Keys(termBreakpoints))
	minAmount := slices.Min(allAmounts)
	maxAmount := slices.Max(allAmounts)

	if amount < minAmount {
		return errors.ErrorAmountBelowLowerBoundary(amount, minAmount)
	}

	if amount > maxAmount {
		return errors.ErrorAmountAboveUpperBoundary(amount, maxAmount)
	}

	return nil
}

func findLowerBreakpoint(amount float64, breakpoints map[float64]float64) (*breakpoint, error) {
	filter := func(breakpointAmount float64) bool {
		return breakpointAmount > amount
	}
	validAmounts := filterAmountsFrom(breakpoints, filter)

	if len(validAmounts) == 0 {
		return nil, errors.NewLowerBreakpointNotFoundError(amount)
	}

	lowerAmount := slices.Max(validAmounts)
	fee, ok := breakpoints[lowerAmount]
	if !ok {
		return nil, errors.NewLowerBreakpointNotFoundError(lowerAmount)
	}

	return &breakpoint{lowerAmount, fee}, nil
}

func findUpperBreakpoint(amount float64, breakpoints map[float64]float64) (*breakpoint, error) {
	filter := func(breakpointAmount float64) bool {
		return breakpointAmount <= amount
	}
	validAmounts := filterAmountsFrom(breakpoints, filter)

	if len(validAmounts) == 0 {
		return getBreakpointForMaxAmount(breakpoints)
	}

	return getBreakpointForUpperAmount(validAmounts, breakpoints)
}

func filterAmountsFrom(breakpoints map[float64]float64, filter func(amount float64) bool) []float64 {
	amounts := slices.Collect(maps.Keys(breakpoints))
	validAmounts := slices.DeleteFunc(
		amounts,
		filter,
	)

	return validAmounts
}

func getBreakpointForMaxAmount(breakpoints map[float64]float64) (*breakpoint, error) {
	allAmounts := slices.Collect(maps.Keys(breakpoints))
	if len(allAmounts) == 0 {
		return nil, errors.NewNotBreakpointsError()
	}

	upperAmount := slices.Max(allAmounts)
	fee, ok := breakpoints[upperAmount]
	if !ok {
		return nil, errors.NewUpperBreakpointNotFoundError(upperAmount)
	}

	return &breakpoint{upperAmount, fee}, nil
}

func getBreakpointForUpperAmount(validAmounts []float64, breakpoints map[float64]float64) (*breakpoint, error) {
	upperAmount := slices.Min(validAmounts)
	fee, ok := breakpoints[upperAmount]
	if !ok {
		return nil, errors.NewUpperBreakpointNotFoundError(upperAmount)
	}

	return &breakpoint{upperAmount, fee}, nil
}

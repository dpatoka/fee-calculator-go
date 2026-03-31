package domain

import (
	"fee-calculator-go/internal/pricing/domain/errors"
	"math"
)

type BreakpointRange struct {
	lowerBreakpoint breakpoint
	upperBreakpoint breakpoint
}

func NewBreakpointRange(lowerAmount, lowerFee, upperAmount, upperFee float64) *BreakpointRange {
	return &BreakpointRange{
		breakpoint{lowerAmount, lowerFee},
		breakpoint{upperAmount, upperFee},
	}
}

func (b *BreakpointRange) CalculateFee(requestedAmount float64) (float64, error) {
	err := b.validateAmount(requestedAmount)
	if err != nil {
		return 0, err
	}

	if requestedAmount == b.lowerBreakpoint.amount {
		return b.lowerBreakpoint.fee, nil
	}

	if requestedAmount == b.upperBreakpoint.amount {
		return b.upperBreakpoint.fee, nil
	}

	interpolatedFee := b.interpolateFee(requestedAmount)

	return b.roundUpToDivisibleByFive(requestedAmount, interpolatedFee), nil
}

func (b *BreakpointRange) validateAmount(amount float64) error {
	if amount < b.lowerBreakpoint.amount {
		return errors.ErrorAmountBelowLowerBoundary(amount, b.lowerBreakpoint.amount)
	}

	if amount > b.upperBreakpoint.amount {
		return errors.ErrorAmountAboveUpperBoundary(amount, b.upperBreakpoint.amount)
	}

	return nil
}

func (b *BreakpointRange) interpolateFee(requestedAmount float64) float64 {
	ratio := (requestedAmount - b.lowerBreakpoint.amount) / (b.upperBreakpoint.amount - b.lowerBreakpoint.amount)

	return b.lowerBreakpoint.fee + (b.upperBreakpoint.fee-b.lowerBreakpoint.fee)*ratio
}

func (b *BreakpointRange) roundUpToDivisibleByFive(requestedAmount float64, fee float64) float64 {
	total := requestedAmount + fee
	roundedTotal := math.Ceil(total/5) * 5

	return roundedTotal - requestedAmount
}

package domain

import "math"

type BreakpointRange struct {
	lowerBreakpoint breakpoint
	upperBreakpoint breakpoint
}

func NewBreakpointRange(lowerBreakpoint breakpoint, upperBreakpoint breakpoint) *BreakpointRange {
	return &BreakpointRange{lowerBreakpoint, upperBreakpoint}
}

func (b *BreakpointRange) CalculateFee(requestedAmount float64) float64 {
	// TODO validateAmount

	if requestedAmount == b.lowerBreakpoint.amount {
		return b.lowerBreakpoint.fee
	}

	if requestedAmount == b.upperBreakpoint.amount {
		return b.upperBreakpoint.fee
	}

	interpolatedFee := b.interpolateFee(requestedAmount)

	return b.roundUpToDivisibleByFive(requestedAmount, interpolatedFee)
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

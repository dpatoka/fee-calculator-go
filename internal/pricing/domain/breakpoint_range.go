package domain

type BreakpointRange struct {
}

func (b *BreakpointRange) CalculateFee(requestedAmount float64) float64 {
	return requestedAmount //TODO
}

package domain

type BreakpointRepository interface {
	GetForTermAndAmount(term int, amount float64) (*BreakpointRange, error)
}

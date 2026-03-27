package query

type FeeCalculationQuery struct {
	Amount float64
	Term   int
}

type FeeCalculationQueryHandler struct {
}

func (f *FeeCalculationQueryHandler) Run(query FeeCalculationQuery) float64 {
	return query.Amount * float64(query.Term)
}

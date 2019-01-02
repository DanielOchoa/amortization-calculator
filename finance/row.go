package finance

// Row for the Table property of an Amortization

type Row struct {
	Month                                                    int
	PaidInterest, PaidPrincipal, RemainingPrincipal, Payment float64
}

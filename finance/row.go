package finance

// Row for the Table property of an Amortization

type Row struct {
	Month                                                  int
	PaidInterest, PaidPrincipal, RemainingBalance, Payment float64
}

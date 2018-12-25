package finance

// Loan implements Calculator, and Mortage as well due to being a type Loan
type Calculator interface {
	DiscountFactor() float64
	PeriodicInterestRate() float64
	PaymentPerPeriod() float64
	TermInMonths()
}

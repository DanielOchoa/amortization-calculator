package finance

import "math"

// An amortization is a type of loan, but with a set time table for repayment.
type Amortization struct {
	*Loan
	Term float64 // in years (e.g. 30 year loan)
}

// Diferent formulas for loan (amortization) calculations.

// Calculates the `Discount Factor`. Divide principal / discount factor to get Monthly Payment
//
// periodicPayment = payments per year * number of years
// periodicInterestRate = see `PeriodicInterestRate` method
//
func (a *Amortization) DiscountFactor() float64 {
	return (math.Pow((1.00+a.PeriodicInterestRate()), a.TermInMonths()) - 1.00) /
		(a.PeriodicInterestRate() * math.Pow(1.00+a.PeriodicInterestRate(), a.TermInMonths()))
}

// Calculates the periodic interest rate used to calculate the (monthly?) loan payment.
//
// interestRate = e.g. 6.00
// paymentPeriods = payment periods in a year, e.g. 12
// periodicInterestRate = annual rate (decimal) / number of payment periods (6% = (.06 / 12) = .005)
//
func (a *Amortization) PeriodicInterestRate() float64 {
	return (a.InterestRate / 100.00) / a.Schedule
}

// Loan payment per pay period. For example, the amount to be paid any given month on a fixed payment schedule.
func (a *Amortization) PaymentPerPeriod() float64 {
	return a.Principal / a.DiscountFactor()
}

// Payments per year * Term (in years)
func (a *Amortization) TermInMonths() float64 {
	return a.Schedule * a.Term
}

package finance

import "math"
import "fmt"

// An amortization is a type of loan, but with a set time table for repayment.
type Amortization struct {
	*Loan
	Term  int   // in years (e.g. 30 year loan)
	Table []Row // TODO: make Table struct for extra method calls on it.
	// TODO: figure out an idiomatic go way of memoizing these fields
	_termInMonths     int
	_paymentPerPeriod float64
}

// NOTE: Not needed yet.
type Amortizable interface {
	DiscountFactor() float64
	PeriodicInterestRate() float64
	PaymentPerPeriod() float64
	TermInMonths() int
	Calculate()
	GenerateRow(int, float64) Row
	CalcInterestPaid(float64) float64
}

// Calculates the `Discount Factor`. Divide principal / discount factor to get Monthly Payment
//
// periodicPayment = payments per year * number of years
// periodicInterestRate = see `PeriodicInterestRate` method
//
func (a *Amortization) DiscountFactor() float64 {
	return (math.Pow((1.00+a.PeriodicInterestRate()), float64(a.TermInMonths())) - 1.00) /
		(a.PeriodicInterestRate() * math.Pow(1.00+a.PeriodicInterestRate(), float64(a.TermInMonths())))
}

// Calculates the periodic interest rate used to calculate the (monthly?) loan payment.
//
// interestRate = e.g. 6.00
// paymentPeriods = payment periods in a year, e.g. 12
// periodicInterestRate = annual rate (decimal) / number of payment periods (6% = (.06 / 12) = .005)
//
func (a *Amortization) PeriodicInterestRate() float64 {
	return (a.InterestRate / 100.00) / float64(a.Loan.Schedule)
}

// Loan payment per pay period. For example, the amount to be paid any given month on a fixed payment schedule.
// needs to be memoized
func (a *Amortization) PaymentPerPeriod() float64 {
	if a._paymentPerPeriod <= 0.0 {
		a._paymentPerPeriod = a.Principal / a.DiscountFactor()
	}
	return a._paymentPerPeriod
}

// Payments per year * Term (in years), or, number of total payments.
// should be memoized
func (a *Amortization) TermInMonths() int {
	if a._termInMonths <= 0 {
		a._termInMonths = a.Loan.Schedule * a.Term
	}
	return a._termInMonths
}

// Amortization calculation
func (a *Amortization) Calculate() {
	termInMonths := a.TermInMonths()
	for month := 1; month <= termInMonths; month++ {
		var remainingPrincipal float64
		// get previous row remaining principal if available
		if len(a.Table) > 0 {
			remainingPrincipal = a.Table[len(a.Table)-1].RemainingPrincipal
		} else {
			// TODO: When adding field to pay extra to principal, we need to take into account remaining principal
			// below zero.
			remainingPrincipal = a.Loan.Principal
		}
		row := a.GenerateRow(month, remainingPrincipal)
		a.Table = append(a.Table, row)
	}
}

// Calculates a row in the Amortization Table
func (a *Amortization) GenerateRow(month int, remainingPrincipal float64) Row {
	paidInterest := a.CalcInterestPaid(remainingPrincipal)
	paidPrincipal := a.PaymentPerPeriod() - paidInterest
	newPrincipal := remainingPrincipal - paidPrincipal
	// Month, PaidInterest, PaidPrincipal, RemainingPrincipal, Payment
	return Row{month, paidInterest, paidPrincipal, newPrincipal, a.PaymentPerPeriod()}
}

func (a *Amortization) CalcInterestPaid(remainingPrincipal float64) float64 {
	return (remainingPrincipal * (a.Loan.InterestRate / 100)) / float64(a.Loan.Schedule)
}

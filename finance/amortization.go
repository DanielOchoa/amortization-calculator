package finance

import "math"

// Row for the Table property of an Amortization
type Row struct {
	Month                                                  int
	PaidInterest, PaidPrincipal, RemainingBalance, Payment float64
}

// An amortization is a type of loan, but with a set time table for repayment.
type Amortization struct {
	*Loan
	Term  int // in years (e.g. 30 year loan)
	Table []Row

	// TODO: figure out an idiomatic go way of memoizing these fields
	_termInMonths     int
	_paymentPerPeriod float64
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

// Calculate amortization. Satifies Calculable interface.
func (a *Amortization) Calculate() {
	termInMonths := a.TermInMonths()
	for month := 1; month <= termInMonths; month++ {
		remainingBalance := a.GetRemainingBalance()
		row := a.GenerateRow(month, remainingBalance)
		a.Table = append(a.Table, row)
	}
}

func (a *Amortization) GetRemainingBalance() float64 {
	// get previous row remaining principal if available
	var remainingBalance float64
	if len(a.Table) > 0 {
		remainingBalance = a.Table[len(a.Table)-1].RemainingBalance
	} else {
		// TODO: When adding field to pay extra to principal, we need to take into account remaining principal
		// below zero.
		remainingBalance = a.Loan.Principal
	}
	return remainingBalance
}

// Calculates a row in the Amortization Table
func (a *Amortization) GenerateRow(month int, remainingBalance float64) Row {
	paidInterest := a.CalcInterestPaid(remainingBalance)
	paidPrincipal := a.PaymentPerPeriod() - paidInterest
	newPrincipal := remainingBalance - paidPrincipal
	//a.RemainingBalance = newPrincipal // TODO: also being set on Row...
	// Month, PaidInterest, PaidPrincipal, RemainingBalance, Payment
	return Row{month, paidInterest, paidPrincipal, newPrincipal, a.PaymentPerPeriod()}
}

func (a *Amortization) CalcInterestPaid(remainingBalance float64) float64 {
	return (remainingBalance * (a.Loan.InterestRate / 100)) / float64(a.Loan.Schedule)
}

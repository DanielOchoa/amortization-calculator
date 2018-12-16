package calc

import "math"

// Diferent formulas for amortization calculations.

// Calculates the discount factor. Divide principal / discount factor to get monthly payment
//
// periodicPayment = payments per year * number of years
//
func DiscountFactor(periodicPayment float64, periodicInterestRate float64) float64 {
	return (math.Pow((1.00+periodicInterestRate), periodicPayment) - 1.00) / (periodicInterestRate * math.Pow(1.00+periodicInterestRate, periodicPayment))
}

// Calculates the periodic interest rate used to calculate the (monthly?) loan payment.
//
// interestRate = e.g. 6.00
// paymentPeriods = payment periods in a year, e.g. 12
// periodicInterestRate = annual rate (decimal) / number of payment periods (6% = (.06 / 12) = .005)
//
func PeriodicInterestRate(interestRate float64, paymentPeriods float64) float64 {
	return (interestRate / 100) / paymentPeriods
}

// Loan payment per pay period. For example, the amount to be paid any given month on a fixed payment schedule.
func PaymentForPeriod(balance float64, discountFactor float64) float64 {
	return balance / discountFactor
}

package main

import "fmt"
import "github.com/DanielOchoa/amortization-calculator/calc"

func main() {
	var numberOfPayments float64 = 360
	var interestRate float64 = 4
	var borrowedAmount float64 = 324550
	var paymentPeriods float64 = 12

	periodicInterestRate := calc.PeriodicInterestRate(interestRate, paymentPeriods)
	discountFactor := calc.DiscountFactor(numberOfPayments, periodicInterestRate)
	monthlyPayment := calc.PaymentForPeriod(borrowedAmount, discountFactor)

	fmt.Printf("discountFactor: %f\n", discountFactor)
	fmt.Printf("monthly payment: %f\n", monthlyPayment)
}

package main

import "fmt"
import "github.com/DanielOchoa/amortization-calculator/finance"

func main() {
	amort := &finance.Amortization{
		Loan: &finance.Loan{
			Principal:    100000.00,
			InterestRate: 7.5,
			Accrues:      finance.CompoundedYearly,
			Schedule:     finance.Monthly,
		},
		Term: 30,
	}

	periodicInterestRate := amort.PeriodicInterestRate()
	expectedPeriodicInterestRate := 0.006250

	discountFactor := amort.DiscountFactor()
	monthlyPayment := amort.PaymentPerPeriod()

	fmt.Printf("discountFactor: %f\n", discountFactor)
	fmt.Printf("monthly payment: %f\n", monthlyPayment)
	fmt.Printf("periodic int rate: %f\n", periodicInterestRate)

	fmt.Print(fmt.Sprintf("%f == %f", periodicInterestRate, expectedPeriodicInterestRate))
	fmt.Printf("%v", fmt.Sprintf("%f", periodicInterestRate) == fmt.Sprintf("%f", expectedPeriodicInterestRate))
}

// todo:
// 1 calculate interest being compounded (e.g. amort tables)

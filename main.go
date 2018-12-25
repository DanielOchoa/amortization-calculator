package main

import "fmt"
import "github.com/DanielOchoa/amortization-calculator/finance"

func main() {
	loan := &finance.Loan{
		Principal:    324550.00,
		InterestRate: 4.00,
		Accrues:      finance.CompoundedYearly,
		Schedule:     finance.Monthly,
	}

	amort := &finance.Amortization{
		Loan: loan,
		Term: 30,
	}

	periodicInterestRate := amort.PeriodicInterestRate()
	discountFactor := amort.DiscountFactor()
	monthlyPayment := amort.PaymentPerPeriod()

	fmt.Printf("discountFactor: %f\n", discountFactor)
	fmt.Printf("monthly payment: %f\n", monthlyPayment)
	fmt.Printf("periodic int rate: %f\n", periodicInterestRate)
}

// todo:
// 1 calculate interest being compounded (e.g. amort tables)

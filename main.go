package main

import "fmt"
import "github.com/DanielOchoa/amortization-calculator/finance"

func main() {
	var amortTerm int = 30
	amort := &finance.Amortization{
		Loan: &finance.Loan{
			Principal:    100000.00,
			InterestRate: 7.5,
			Accrues:      finance.CompoundedYearly,
			Schedule:     finance.Monthly,
		},
		Term: amortTerm,
	}

	periodicInterestRate := amort.PeriodicInterestRate()

	discountFactor := amort.DiscountFactor()
	monthlyPayment := amort.PaymentPerPeriod()

	fmt.Printf("discountFactor: %f\n", discountFactor)
	fmt.Printf("monthly payment: %f\n", monthlyPayment)
	fmt.Printf("periodic int rate: %f\n", periodicInterestRate)

	fmt.Print("\n")
	fmt.Print("\n")

	fmt.Printf("Term: %d\n", amortTerm)

	// calculate amort
	amort.Calculate()

	for _, row := range amort.Table {
		fmt.Printf("month: %d\n", row.Month)
		fmt.Printf("interest paid: %f\n", row.PaidInterest)
		fmt.Printf("paid to principal: %f\n", row.PaidPrincipal)
		fmt.Printf("pay amount: %f\n", row.Payment)
		fmt.Printf("%f of int + %f of principal == %f\n", row.PaidInterest, row.PaidPrincipal, row.PaidInterest+row.PaidPrincipal)
		fmt.Printf("PRINCIPAL remaining: %f\n", row.RemainingPrincipal)
	}
}

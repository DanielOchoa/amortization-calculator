package main

import "fmt"
import "github.com/DanielOchoa/amortization-calculator/finance"

// TODO:
// - Display at end total interest paid

func main() {
	var amortTerm int = 30
	amort := &finance.Amortization{
		Loan: &finance.Loan{
			Principal:    324000.00,
			InterestRate: 4.0,
			Accrues:      finance.CompoundedYearly,
			Schedule:     finance.Monthly,
			ExtraPayment: 300.00, // still to be implemented
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

	// Create and run calculator
	calculator := finance.NewCalculator(amort)
	calculator.Calculate()

	for _, row := range amort.Table {
		fmt.Printf("month: %d\n", row.Month)
		fmt.Printf("interest paid: %f\n", row.PaidInterest)
		fmt.Printf("paid to principal: %f\n", row.PaidPrincipal)
		fmt.Printf("pay amount: %f\n", row.Payment)
		fmt.Printf("%f of int + %f of principal == %f\n", row.PaidInterest, row.PaidPrincipal, row.PaidInterest+row.PaidPrincipal)
		fmt.Printf("remaining balance: %f\n", row.RemainingBalance)
		fmt.Print("\n")
	}
	fmt.Printf("\n\nDuration in ms: %f\n", calculator.ElapsedInMs())
}

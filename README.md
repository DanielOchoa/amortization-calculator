# Amortization Calculator

This is an amortization calculator written in go. It's still in a very early stage. I created it to help me
understand how an amortization is calculated.

### Current features

+ Calculate an amortization based a loan amount, term and interest rate to get:
  + Paid to principal each month
  + Paid in interest each month
  + Total monthly payment
  + remaining balance each month

+ Track how long it takes for an amortization schedule to calculate.

### To do

+ Add ability to pay extra each month, at any given month.
+ Add feature to generate visualizations in an image or pdf format.
+ Calculate other type of interest accumulating loans, such as credit cards.
+ Calculate various federal student loans (from the US)

## Usage

```go
loan := &finance.Loan{
    Principal:    324000.00,
    InterestRate: 4.0,
    Accrues:      finance.CompoundedYearly,
    Schedule:     finance.Monthly,
}

amortization := &finance.Amortization{
    Loan: loan,
    Term: 30,
}

calculator := finance.NewCalculator(amortization)
calculator.Calculate()

for _, row := range amortization.Table {
    fmt.Printf("month: %d\n", row.Month)
    fmt.Printf("interest paid in month: %f\n", row.PaidInterest)
    fmt.Printf("paid to principal: %f\n", row.PaidPrincipal)
    fmt.Printf("total pay amount: %f\n", row.Payment)
    fmt.Printf("remaining balance: %f\n", row.RemainingBalance)
    fmt.Print("\n")
}
fmt.Printf("\n\nDuration in ms: %f\n", calculator.ElapsedInMs())

```

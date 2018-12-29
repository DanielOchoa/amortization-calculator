package finance

import (
	"fmt"
	"testing"
)
import "github.com/DanielOchoa/amortization-calculator/finance"

const float64EqualityThreshold = 1e-9

func TestAmortization(t *testing.T) {
	// case one
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

	if fmt.Sprintf("%f", periodicInterestRate) != fmt.Sprintf("%f", expectedPeriodicInterestRate) {
		t.Errorf("Periodic Interest rate should be %f and it is %f", expectedPeriodicInterestRate, periodicInterestRate)
	}

	discountFactor := amort.DiscountFactor()
	expectedDiscountFactor := 143.017627

	if !isStringEqual(discountFactor, expectedDiscountFactor) {
		t.Errorf("Discount factor mismatch: (%f, %f)", discountFactor, expectedDiscountFactor)
	}

	monthlyPayment := amort.PaymentPerPeriod()
	expectedMonthlyPayment := 699.214509

	if !isStringEqual(monthlyPayment, expectedMonthlyPayment) {
		t.Errorf("Monthly payment mismatch: (%f, %f)", monthlyPayment, expectedMonthlyPayment)
	}
}

// we compare floats after converting to string since floating point arithmetic is not precise and will never reach
// equality comparison.
func isStringEqual(first, second float64) bool {
	return fmt.Sprintf("%f", first) == fmt.Sprintf("%f", second)
}

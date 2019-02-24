package scheduler

// unused.
// The goal is to have a struct to handle multiple amortizations processed in paralell by goroutines.

type AmortizationTable struct {
	schedules []Schedule
}

type Schedule struct {
	principal      *float64
	monthlyPayment *float64
}

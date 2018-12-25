package scheduler

type AmortizationTable struct {
	schedules []Schedule
}

type Schedule struct {
	principal      *float64
	monthlyPayment *float64
}

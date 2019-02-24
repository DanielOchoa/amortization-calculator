package finance

import "time"

// Anything that can be calculated should implement this interface.
//
// Used for:
// - Decorating calculable objects
//
type Calculable interface {
	Calculate()
}

// Decorator struct used to wrap calculable objects so we can measure how long it takes to execute a calculation.
type Tracker struct {
	Elapsed    time.Duration
	Calculable Calculable
}

// Measure duration of wrapping object calculation
func (t *Tracker) Calculate() {
	startTime := time.Now()
	t.Calculable.Calculate()
	endTime := time.Now()
	t.Elapsed = endTime.Sub(startTime)
}

func (t *Tracker) ElapsedInMs() float64 {
	return float64(t.Elapsed) / float64(time.Millisecond)
}

// Instead of calling `Calculate()` directly on the amortization, we use the `NewCalculator(amortization)`
// constructor. It decorates the amortization
func NewCalculator(calculable Calculable) *Tracker {
	return &Tracker{Calculable: calculable}
}

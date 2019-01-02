package finance

// Enums
// Type aliases
type AccrualType = int
type ScheduleType = int

const (
	CompoundedDaily   AccrualType = 0
	CompoundedMonthly AccrualType = 1
	CompoundedYearly  AccrualType = 2

	Monthly ScheduleType = 12
	Yearly  ScheduleType = 1
)

//
// Loan type struct
//
type Loan struct {
	Principal    float64
	InterestRate float64
	Accrues      AccrualType  // Daily or Yearly
	Schedule     ScheduleType // How many payments per year (monthly = 12, yearly = 1)
}

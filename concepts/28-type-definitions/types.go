package typedefinitions

// Cents is an amount of money in US cents. It is a separate type from int64,
// so it cannot be mixed up with other numbers by accident.
type Cents int64

// Frequency is how often a recurring payment happens.
type Frequency string

// The supported frequencies.
const (
	Weekly    Frequency = "weekly"
	Monthly   Frequency = "monthly"
	Quarterly Frequency = "quarterly"
	Yearly    Frequency = "yearly"
)

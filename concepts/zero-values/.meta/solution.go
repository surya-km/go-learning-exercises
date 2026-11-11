// Package zerovalues practices the zero values concept. Read CONCEPT.md first.
package zerovalues

// ApplyDefaults replaces unset (zero-value) settings with defaults: "USD" for
// currency, 50000 for dailyLimitCents and []string{"General"} for a nil
// categories slice. An empty but non-nil categories slice is kept as is.
func ApplyDefaults(currency string, dailyLimitCents int, categories []string) (string, int, []string) {
	if currency == "" {
		currency = "USD"
	}
	if dailyLimitCents == 0 {
		dailyLimitCents = 50000
	}
	if categories == nil {
		categories = []string{"General"}
	}
	return currency, dailyLimitCents, categories
}

// AddSpend adds cents to the total for category and returns the map.
// A nil totals map is replaced with a new one first.
func AddSpend(totals map[string]int, category string, cents int) map[string]int {
	if totals == nil {
		totals = make(map[string]int)
	}
	totals[category] += cents // a missing key reads as 0
	return totals
}

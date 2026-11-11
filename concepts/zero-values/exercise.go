// Package zerovalues practices the zero values concept. Read CONCEPT.md first.
package zerovalues

// ApplyDefaults replaces unset (zero-value) settings with defaults: "USD" for
// currency, 50000 for dailyLimitCents and []string{"General"} for a nil
// categories slice. An empty but non-nil categories slice is kept as is.
func ApplyDefaults(currency string, dailyLimitCents int, categories []string) (string, int, []string) {
	// TODO: implement
	return "", 0, nil
}

// AddSpend adds cents to the total for category and returns the map.
// A nil totals map is replaced with a new one first.
func AddSpend(totals map[string]int, category string, cents int) map[string]int {
	// TODO: implement
	return nil
}

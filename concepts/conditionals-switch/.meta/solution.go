// Package conditionalsswitch practices the conditionals-switch concept. Read CONCEPT.md first.
package conditionalsswitch

// Category returns the spending category for a merchant name, or "Other"
// when the merchant is not recognised.
func Category(merchant string) string {
	switch merchant {
	case "Tesco", "Aldi", "Lidl":
		return "Groceries"
	case "Shell", "BP", "Uber":
		return "Transport"
	case "Netflix", "Spotify":
		return "Entertainment"
	default:
		return "Other"
	}
}

// CreditRating turns a credit score (300 to 850) into a rating such as
// "Excellent" or "Poor". Scores outside that range are "Invalid".
func CreditRating(score int) string {
	switch {
	case score < 300 || score > 850:
		return "Invalid"
	case score >= 800:
		return "Excellent"
	case score >= 740:
		return "Very good"
	case score >= 670:
		return "Good"
	case score >= 580:
		return "Fair"
	default:
		return "Poor"
	}
}

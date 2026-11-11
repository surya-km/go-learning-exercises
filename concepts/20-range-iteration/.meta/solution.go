// Package rangeiteration practices the range iteration concept. Read CONCEPT.md first.
package rangeiteration

import "strings"

// RunningBalance returns the balance after applying each amount, in order,
// starting from startCents. It never returns nil.
func RunningBalance(startCents int, amounts []int) []int {
	balances := make([]int, 0, len(amounts))
	balance := startCents
	for _, amt := range amounts {
		balance += amt
		balances = append(balances, balance)
	}
	return balances
}

// InstallmentPlan splits totalCents into months payments that differ by at most
// one cent, with leftover cents on the earliest payments. It never returns nil.
func InstallmentPlan(totalCents, months int) []int {
	if months <= 0 {
		return []int{}
	}
	base, leftover := totalCents/months, totalCents%months
	plan := make([]int, 0, months)
	for i := range months {
		payment := base
		if i < leftover {
			payment++
		}
		plan = append(plan, payment)
	}
	return plan
}

// MaskCard replaces every digit except the last four with '*', keeping other
// characters unchanged.
func MaskCard(card string) string {
	digits := 0
	for _, r := range card {
		if isDigit(r) {
			digits++
		}
	}

	var b strings.Builder
	seen := 0
	for _, r := range card {
		if isDigit(r) {
			seen++
			if seen <= digits-4 {
				b.WriteRune('*')
				continue
			}
		}
		b.WriteRune(r)
	}
	return b.String()
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// BiggestCategory returns the category with the largest total. Ties go to the
// alphabetically first category. An empty map returns "", 0.
func BiggestCategory(totals map[string]int) (string, int) {
	bestName, bestTotal := "", 0
	first := true
	for name, total := range totals {
		if first || total > bestTotal || (total == bestTotal && name < bestName) {
			bestName, bestTotal = name, total
			first = false
		}
	}
	return bestName, bestTotal
}

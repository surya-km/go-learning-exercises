// Package forloops practices the for loops concept. Read CONCEPT.md first.
package forloops

// MonthsToPayOff returns how many months of paymentCents it takes to pay off
// balanceCents when monthlyRatePercent interest is added each month first.
// It returns 0 if nothing is owed and -1 if the loan can never be paid off.
func MonthsToPayOff(balanceCents, monthlyRatePercent, paymentCents int) int {
	months := 0
	for balanceCents > 0 {
		interest := balanceCents * monthlyRatePercent / 100
		if paymentCents <= interest {
			return -1
		}
		balanceCents += interest - paymentCents
		months++
	}
	return months
}

// CountWeekdays returns how many of the days days starting on startDay
// (0 = Sunday ... 6 = Saturday) fall on Monday to Friday.
func CountWeekdays(startDay, days int) int {
	count := 0
	for i := 0; i < days; i++ {
		day := (startDay + i) % 7
		if day == 0 || day == 6 {
			continue
		}
		count++
	}
	return count
}

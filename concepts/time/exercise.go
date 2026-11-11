// Package timepractice practices the time concept. Read CONCEPT.md first.
package timepractice

import "time"

// DaysInMonth returns the number of days in the given month.
func DaysInMonth(year int, month time.Month) int {
	// TODO: implement
	return 0
}

// NextDueDate returns the first monthly due date strictly after the given date.
// When a month is shorter than dueDay, the bill is due on its last day.
func NextDueDate(after time.Time, dueDay int) time.Time {
	// TODO: implement
	return time.Time{}
}

// DaysBetween parses two "2006-01-02" dates and returns the days from from to to.
func DaysBetween(from, to string) (int, error) {
	// TODO: implement
	return 0, nil
}

// FormatDue formats t like "Thu 29 Feb 2024".
func FormatDue(t time.Time) string {
	// TODO: implement
	return ""
}

// Package timepractice practices the time concept. Read CONCEPT.md first.
package timepractice

import "time"

// DaysInMonth returns the number of days in the given month.
func DaysInMonth(year int, month time.Month) int {
	// Day 0 of next month is the last day of this month.
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// NextDueDate returns the first monthly due date strictly after the given date.
// When a month is shorter than dueDay, the bill is due on its last day.
func NextDueDate(after time.Time, dueDay int) time.Time {
	year, month, day := after.Date()

	due := dueDateIn(year, month, dueDay)
	if due.Day() > day {
		return due
	}

	// Already past this month's due date: move to the 1st of next month
	// (time.Date normalizes December+1 into January of the next year).
	next := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
	return dueDateIn(next.Year(), next.Month(), dueDay)
}

// dueDateIn returns the due date in the given month, clamped to its last day.
func dueDateIn(year int, month time.Month, dueDay int) time.Time {
	day := min(dueDay, DaysInMonth(year, month))
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// DaysBetween parses two "2006-01-02" dates and returns the days from from to to.
func DaysBetween(from, to string) (int, error) {
	const layout = "2006-01-02"
	start, err := time.Parse(layout, from)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse(layout, to)
	if err != nil {
		return 0, err
	}
	return int(end.Sub(start).Hours() / 24), nil
}

// FormatDue formats t like "Thu 29 Feb 2024".
func FormatDue(t time.Time) string {
	return t.Format("Mon 2 Jan 2006")
}

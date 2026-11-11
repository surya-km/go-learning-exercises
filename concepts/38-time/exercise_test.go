package timepractice

import (
	"testing"
	"time"
)

// date builds midnight UTC on the given day, so tests never depend on the local timezone.
func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		want  int
	}{
		{year: 2024, month: time.January, want: 31},
		{year: 2024, month: time.February, want: 29},
		{year: 2023, month: time.February, want: 28},
		{year: 1900, month: time.February, want: 28}, // divisible by 100: not a leap year
		{year: 2000, month: time.February, want: 29}, // divisible by 400: leap year
		{year: 2024, month: time.April, want: 30},
		{year: 2024, month: time.December, want: 31},
	}

	for _, tt := range tests {
		t.Run(tt.month.String(), func(t *testing.T) {
			if got := DaysInMonth(tt.year, tt.month); got != tt.want {
				t.Errorf("DaysInMonth(%d, %s) = %d, want %d", tt.year, tt.month, got, tt.want)
			}
		})
	}
}

func TestNextDueDate(t *testing.T) {
	tests := []struct {
		name   string
		after  time.Time
		dueDay int
		want   time.Time
	}{
		{name: "later this month", after: date(2024, time.January, 15), dueDay: 20, want: date(2024, time.January, 20)},
		{name: "same day rolls to next month", after: date(2024, time.January, 20), dueDay: 20, want: date(2024, time.February, 20)},
		{name: "already passed this month", after: date(2024, time.March, 25), dueDay: 1, want: date(2024, time.April, 1)},
		{name: "31st into leap February", after: date(2024, time.January, 31), dueDay: 31, want: date(2024, time.February, 29)},
		{name: "31st into normal February", after: date(2023, time.January, 31), dueDay: 31, want: date(2023, time.February, 28)},
		{name: "30th into February", after: date(2023, time.January, 30), dueDay: 30, want: date(2023, time.February, 28)},
		{name: "back to full day after February", after: date(2024, time.February, 29), dueDay: 31, want: date(2024, time.March, 31)},
		{name: "31st in a 30-day month", after: date(2024, time.April, 2), dueDay: 31, want: date(2024, time.April, 30)},
		{name: "clamped day this month already passed", after: date(2024, time.April, 30), dueDay: 31, want: date(2024, time.May, 31)},
		{name: "rolls into next year", after: date(2024, time.December, 31), dueDay: 5, want: date(2025, time.January, 5)},
		{name: "December due later this month", after: date(2024, time.December, 1), dueDay: 31, want: date(2024, time.December, 31)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextDueDate(tt.after, tt.dueDay)
			if !got.Equal(tt.want) {
				t.Errorf("NextDueDate(%s, %d) = %s, want %s",
					tt.after.Format("2006-01-02"), tt.dueDay, got.Format("2006-01-02"), tt.want.Format("2006-01-02"))
			}
		})
	}
}

func TestDaysBetween(t *testing.T) {
	tests := []struct {
		name string
		from string
		to   string
		want int
	}{
		{name: "same day", from: "2024-03-10", to: "2024-03-10", want: 0},
		{name: "one week", from: "2024-03-01", to: "2024-03-08", want: 7},
		{name: "across leap day", from: "2024-02-28", to: "2024-03-01", want: 2},
		{name: "across non-leap February", from: "2023-02-28", to: "2023-03-01", want: 1},
		{name: "whole leap year", from: "2024-01-01", to: "2025-01-01", want: 366},
		{name: "backwards is negative", from: "2024-01-10", to: "2024-01-01", want: -9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DaysBetween(tt.from, tt.to)
			if err != nil {
				t.Fatalf("DaysBetween(%q, %q) returned unexpected error: %v", tt.from, tt.to, err)
			}
			if got != tt.want {
				t.Errorf("DaysBetween(%q, %q) = %d, want %d", tt.from, tt.to, got, tt.want)
			}
		})
	}
}

func TestDaysBetweenInvalid(t *testing.T) {
	tests := []struct {
		name string
		from string
		to   string
	}{
		{name: "bad from", from: "03/01/2024", to: "2024-03-08"},
		{name: "bad to", from: "2024-03-01", to: "next week"},
		{name: "impossible date", from: "2023-02-29", to: "2023-03-01"},
		{name: "empty", from: "", to: "2024-03-01"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DaysBetween(tt.from, tt.to)
			if err == nil {
				t.Fatalf("DaysBetween(%q, %q) = %d, nil; want an error", tt.from, tt.to, got)
			}
			if got != 0 {
				t.Errorf("DaysBetween(%q, %q) = %d on error, want 0", tt.from, tt.to, got)
			}
		})
	}
}

func TestFormatDue(t *testing.T) {
	tests := []struct {
		t    time.Time
		want string
	}{
		{t: date(2024, time.February, 29), want: "Thu 29 Feb 2024"},
		{t: date(2024, time.March, 1), want: "Fri 1 Mar 2024"},
		{t: date(2025, time.December, 25), want: "Thu 25 Dec 2025"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := FormatDue(tt.t); got != tt.want {
				t.Errorf("FormatDue(%s) = %q, want %q", tt.t.Format("2006-01-02"), got, tt.want)
			}
		})
	}
}

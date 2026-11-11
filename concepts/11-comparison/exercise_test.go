package comparison

import "testing"

func TestCompareAmounts(t *testing.T) {
	tests := []struct {
		name   string
		aCents int
		bCents int
		want   int
	}{
		{name: "less", aCents: 1500, bCents: 2000, want: -1},
		{name: "greater", aCents: 2000, bCents: 1500, want: 1},
		{name: "equal", aCents: 999, bCents: 999, want: 0},
		{name: "one cent less", aCents: 99, bCents: 100, want: -1},
		{name: "negative vs positive", aCents: -500, bCents: 1, want: -1},
		{name: "both negative", aCents: -100, bCents: -500, want: 1},
		{name: "both zero", aCents: 0, bCents: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareAmounts(tt.aCents, tt.bCents); got != tt.want {
				t.Errorf("CompareAmounts(%d, %d) = %d, want %d", tt.aCents, tt.bCents, got, tt.want)
			}
		})
	}
}

func TestEarlierDate(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{name: "second is earlier", a: "2026-03-15", b: "2026-01-31", want: "2026-01-31"},
		{name: "first is earlier", a: "2026-01-31", b: "2026-03-15", want: "2026-01-31"},
		{name: "across a year end", a: "2026-01-01", b: "2025-12-31", want: "2025-12-31"},
		{name: "same month different day", a: "2026-05-09", b: "2026-05-10", want: "2026-05-09"},
		{name: "same date", a: "2026-07-04", b: "2026-07-04", want: "2026-07-04"},
		{name: "decades apart", a: "2030-01-01", b: "1999-12-31", want: "1999-12-31"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EarlierDate(tt.a, tt.b); got != tt.want {
				t.Errorf("EarlierDate(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

package typedefinitions

import (
	"slices"
	"testing"
)

func TestCentsFormat(t *testing.T) {
	tests := []struct {
		cents Cents
		want  string
	}{
		{cents: 1234, want: "$12.34"},
		{cents: 5, want: "$0.05"},
		{cents: 0, want: "$0.00"},
		{cents: 100, want: "$1.00"},
		{cents: 123456789, want: "$1234567.89"},
		{cents: -1250, want: "-$12.50"},
		{cents: -7, want: "-$0.07"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.cents.Format(); got != tt.want {
				t.Errorf("Cents(%d).Format() = %q, want %q", tt.cents, got, tt.want)
			}
		})
	}
}

func TestCentsSplit(t *testing.T) {
	tests := []struct {
		name  string
		cents Cents
		n     int
		want  []Cents
	}{
		{name: "even split", cents: 900, n: 3, want: []Cents{300, 300, 300}},
		{name: "one leftover cent", cents: 1000, n: 3, want: []Cents{334, 333, 333}},
		{name: "two leftover cents", cents: 101, n: 3, want: []Cents{34, 34, 33}},
		{name: "one part", cents: 4999, n: 1, want: []Cents{4999}},
		{name: "more parts than cents", cents: 2, n: 4, want: []Cents{1, 1, 0, 0}},
		{name: "zero amount", cents: 0, n: 2, want: []Cents{0, 0}},
		{name: "zero parts", cents: 1000, n: 0, want: nil},
		{name: "negative parts", cents: 1000, n: -2, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cents.Split(tt.n)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Cents(%d).Split(%d) = %v, want %v", tt.cents, tt.n, got, tt.want)
			}
		})
	}
}

func TestFrequencyTimesPerYear(t *testing.T) {
	tests := []struct {
		freq Frequency
		want int
	}{
		{freq: Weekly, want: 52},
		{freq: Monthly, want: 12},
		{freq: Quarterly, want: 4},
		{freq: Yearly, want: 1},
		{freq: Frequency("daily"), want: 0},
		{freq: Frequency(""), want: 0},
	}

	for _, tt := range tests {
		t.Run(string(tt.freq)+"_", func(t *testing.T) {
			if got := tt.freq.TimesPerYear(); got != tt.want {
				t.Errorf("Frequency(%q).TimesPerYear() = %d, want %d", tt.freq, got, tt.want)
			}
		})
	}
}

func TestYearlyCost(t *testing.T) {
	tests := []struct {
		amount Cents
		freq   Frequency
		want   Cents
	}{
		{amount: 1599, freq: Monthly, want: 19188},
		{amount: 1000, freq: Weekly, want: 52000},
		{amount: 25000, freq: Quarterly, want: 100000},
		{amount: 9900, freq: Yearly, want: 9900},
		{amount: 500, freq: Frequency("hourly"), want: 0},
	}

	for _, tt := range tests {
		t.Run(string(tt.freq), func(t *testing.T) {
			if got := YearlyCost(tt.amount, tt.freq); got != tt.want {
				t.Errorf("YearlyCost(%d, %q) = %d, want %d", tt.amount, tt.freq, got, tt.want)
			}
		})
	}
}

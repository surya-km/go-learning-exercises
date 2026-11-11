package stringspackage

import (
	"slices"
	"testing"
)

func TestCleanMerchant(t *testing.T) {
	tests := []struct {
		raw  string
		want string
	}{
		{raw: "POS starbucks #1234", want: "STARBUCKS"},
		{raw: "  SQ *blue   bottle coffee ", want: "BLUE BOTTLE COFFEE"},
		{raw: "Whole Foods Market", want: "WHOLE FOODS MARKET"},
		{raw: "Target #0042 Austin TX", want: "TARGET"},
		{raw: "  POS   shell   oil #77 ", want: "SHELL OIL"},
		{raw: "POS SQ *corner cafe", want: "SQ *CORNER CAFE"},
		{raw: "Chipotle POS terminal", want: "CHIPOTLE POS TERMINAL"},
		{raw: "Store#5", want: "STORE#5"},
		{raw: "   ", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.raw, func(t *testing.T) {
			got := CleanMerchant(tt.raw)
			if got != tt.want {
				t.Errorf("CleanMerchant(%q) = %q, want %q", tt.raw, got, tt.want)
			}
		})
	}
}

func TestParseTags(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: "food, travel", want: []string{"food", "travel"}},
		{input: "Food, travel ,, FOOD", want: []string{"food", "travel"}},
		{input: "rent", want: []string{"rent"}},
		{input: "  Bills,utilities , BILLS, bills ,Fun", want: []string{"bills", "utilities", "fun"}},
		{input: "eating out, Eating Out", want: []string{"eating out"}},
		{input: " , ,", want: nil},
		{input: "", want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ParseTags(tt.input)
			// slices.Equal treats nil and an empty slice as equal.
			if !slices.Equal(got, tt.want) {
				t.Errorf("ParseTags(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

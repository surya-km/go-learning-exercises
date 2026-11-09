package parseamount

import (
	"errors"
	"strings"
	"testing"
)

func TestParseAmountValid(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "12.34", want: 1234},
		{input: "7", want: 700},
		{input: "0.5", want: 50},
		{input: "0.05", want: 5},
		{input: "0", want: 0},
		{input: "  3.00 ", want: 300},
		{input: "1000000.99", want: 100000099},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := ParseAmount(tt.input)
			if err != nil {
				t.Fatalf("ParseAmount(%q) returned unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseAmount(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseAmountEmpty(t *testing.T) {
	for _, input := range []string{"", "   "} {
		t.Run(strings.ReplaceAll(input, " ", "_")+"empty", func(t *testing.T) {
			got, err := ParseAmount(input)
			if !errors.Is(err, ErrEmpty) {
				t.Fatalf("ParseAmount(%q) error = %v, want an error matching ErrEmpty", input, err)
			}
			if got != 0 {
				t.Errorf("ParseAmount(%q) = %d on error, want 0", input, got)
			}
		})
	}
}

func TestParseAmountInvalid(t *testing.T) {
	inputs := []string{"abc", "1.234", "1.", ".", "-5", "1.2.3", "12a", "1.x", "$5"}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			got, err := ParseAmount(input)
			if !errors.Is(err, ErrInvalid) {
				t.Fatalf("ParseAmount(%q) error = %v, want an error matching ErrInvalid", input, err)
			}
			if got != 0 {
				t.Errorf("ParseAmount(%q) = %d on error, want 0", input, got)
			}
			if !strings.Contains(err.Error(), input) {
				t.Errorf("error message %q should mention the input %q", err.Error(), input)
			}
		})
	}
}

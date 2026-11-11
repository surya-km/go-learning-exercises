package regularexpressions

import (
	"slices"
	"testing"
)

func TestValidRoutingNumber(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "021000021", want: true},
		{input: "111000025", want: true},
		{input: "02100002", want: false},
		{input: "0210000210", want: false},
		{input: "02100002a", want: false},
		{input: " 021000021", want: false},
		{input: "021000021\n", want: false},
		{input: "021-000-021", want: false},
		{input: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ValidRoutingNumber(tt.input)
			if got != tt.want {
				t.Errorf("ValidRoutingNumber(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestExtractAmounts(t *testing.T) {
	tests := []struct {
		text string
		want []int
	}{
		{text: "Subtotal $12.50, tax $1.03", want: []int{1250, 103}},
		{text: "You owe $7 now", want: []int{700}},
		{text: "Total: $1500.00", want: []int{150000}},
		{text: "$0.99 then $0.01 then $10", want: []int{99, 1, 1000}},
		{text: "$3.5 off", want: []int{300}},
		{text: "Price: 12.50", want: nil},
		{text: "Just a $ sign", want: nil},
		{text: "", want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			got := ExtractAmounts(tt.text)
			// slices.Equal treats nil and an empty slice as equal.
			if !slices.Equal(got, tt.want) {
				t.Errorf("ExtractAmounts(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}

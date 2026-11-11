package money

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name  string
		input Cents
		want  string
	}{
		{name: "dollars and cents", input: 1234, want: "$12.34"},
		{name: "under a dollar", input: 5, want: "$0.05"},
		{name: "ten cents", input: 10, want: "$0.10"},
		{name: "whole dollars", input: 700, want: "$7.00"},
		{name: "zero", input: 0, want: "$0.00"},
		{name: "large amount", input: 123456, want: "$1234.56"},
		{name: "negative", input: -1250, want: "-$12.50"},
		{name: "negative cents only", input: -5, want: "-$0.05"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.input); got != tt.want {
				t.Errorf("Format(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

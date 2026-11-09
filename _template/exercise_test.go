package template

import "testing"

func TestExample(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "describe the case", input: "in", want: "out"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Example(tt.input); got != tt.want {
				t.Errorf("Example(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

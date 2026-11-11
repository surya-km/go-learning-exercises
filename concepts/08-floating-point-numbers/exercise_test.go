package floatingpointnumbers

import (
	"fmt"
	"testing"
)

func TestApplyRate(t *testing.T) {
	tests := []struct {
		cents int
		rate  float64
		want  int
	}{
		{cents: 10000, rate: 5, want: 500},
		{cents: 1999, rate: 8.25, want: 165},
		{cents: 150, rate: 1, want: 2},
		{cents: 149, rate: 1, want: 1},
		{cents: 123456, rate: 3.5, want: 4321},
		{cents: 999, rate: 100, want: 999},
		{cents: -1000, rate: 2.5, want: -25},
		{cents: 1000, rate: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_at_%g%%", tt.cents, tt.rate), func(t *testing.T) {
			got := ApplyRate(tt.cents, tt.rate)
			if got != tt.want {
				t.Errorf("ApplyRate(%d, %g) = %d, want %d", tt.cents, tt.rate, got, tt.want)
			}
		})
	}
}

func TestAlmostEqual(t *testing.T) {
	tests := []struct {
		name      string
		a, b, tol float64
		want      bool
	}{
		{name: "0.1+0.2 vs 0.3", a: 0.1 + 0.2, b: 0.3, tol: 1e-9, want: true},
		{name: "identical", a: 2.5, b: 2.5, tol: 0, want: true},
		{name: "too far apart", a: 1.0, b: 1.1, tol: 0.01, want: false},
		{name: "exactly at tolerance", a: 5, b: 3, tol: 2, want: true},
		{name: "order does not matter", a: 3, b: 5, tol: 2, want: true},
		{name: "opposite signs", a: -1.0, b: 1.0, tol: 1.5, want: false},
		{name: "negatives close", a: -100.001, b: -100.002, tol: 0.01, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AlmostEqual(tt.a, tt.b, tt.tol)
			if got != tt.want {
				t.Errorf("AlmostEqual(%v, %v, %v) = %v, want %v", tt.a, tt.b, tt.tol, got, tt.want)
			}
		})
	}
}

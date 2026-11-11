package typeconversion

import (
	"fmt"
	"math"
	"testing"
)

func TestAverageCents(t *testing.T) {
	tests := []struct {
		total, count int
		want         float64
	}{
		{total: 1000, count: 4, want: 250},
		{total: 1000, count: 3, want: 333.3333333333333},
		{total: 5, count: 2, want: 2.5},
		{total: 1, count: 8, want: 0.125},
		{total: -900, count: 4, want: -225},
		{total: 700, count: 0, want: 0},
		{total: 700, count: -1, want: 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_over_%d", tt.total, tt.count), func(t *testing.T) {
			got := AverageCents(tt.total, tt.count)
			if math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("AverageCents(%d, %d) = %v, want %v", tt.total, tt.count, got, tt.want)
			}
		})
	}
}

func TestOrderSummaryValid(t *testing.T) {
	tests := []struct {
		quantity, unit string
		want           string
	}{
		{quantity: "3", unit: "350", want: "3 items, 1050 cents"},
		{quantity: "1", unit: "499", want: "1 item, 499 cents"},
		{quantity: "0", unit: "250", want: "0 items, 0 cents"},
		{quantity: "12", unit: "65", want: "12 items, 780 cents"},
		{quantity: "65", unit: "1", want: "65 items, 65 cents"},
	}

	for _, tt := range tests {
		t.Run(tt.quantity+"x"+tt.unit, func(t *testing.T) {
			got, err := OrderSummary(tt.quantity, tt.unit)
			if err != nil {
				t.Fatalf("OrderSummary(%q, %q) returned unexpected error: %v", tt.quantity, tt.unit, err)
			}
			if got != tt.want {
				t.Errorf("OrderSummary(%q, %q) = %q, want %q", tt.quantity, tt.unit, got, tt.want)
			}
		})
	}
}

func TestOrderSummaryInvalid(t *testing.T) {
	tests := []struct {
		quantity, unit string
	}{
		{quantity: "two", unit: "350"},
		{quantity: "2", unit: "3.50"},
		{quantity: "", unit: "100"},
		{quantity: "4", unit: ""},
		{quantity: " 4", unit: "100"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q_x_%q", tt.quantity, tt.unit), func(t *testing.T) {
			got, err := OrderSummary(tt.quantity, tt.unit)
			if err == nil {
				t.Fatalf("OrderSummary(%q, %q) = %q, nil; want an error", tt.quantity, tt.unit, got)
			}
			if got != "" {
				t.Errorf("OrderSummary(%q, %q) = %q on error, want \"\"", tt.quantity, tt.unit, got)
			}
		})
	}
}

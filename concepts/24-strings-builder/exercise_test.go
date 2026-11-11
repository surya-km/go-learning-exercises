package stringsbuilder

import (
	"strconv"
	"testing"
)

func TestReceipt(t *testing.T) {
	tests := []struct {
		name  string
		items []LineItem
		want  string
	}{
		{
			name:  "two items",
			items: []LineItem{{Name: "Coffee", Cents: 450}, {Name: "Bagel", Cents: 325}},
			want:  "Coffee: 450\nBagel: 325\nTOTAL: 775\n",
		},
		{
			name:  "one item",
			items: []LineItem{{Name: "Rent", Cents: 150000}},
			want:  "Rent: 150000\nTOTAL: 150000\n",
		},
		{
			name: "includes a refund",
			items: []LineItem{
				{Name: "Shoes", Cents: 8999},
				{Name: "Socks", Cents: 1299},
				{Name: "Refund", Cents: -1299},
			},
			want: "Shoes: 8999\nSocks: 1299\nRefund: -1299\nTOTAL: 8999\n",
		},
		{
			name:  "empty",
			items: nil,
			want:  "TOTAL: 0\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Receipt(tt.items)
			if got != tt.want {
				t.Errorf("Receipt(%+v) =\n%q\nwant\n%q", tt.items, got, tt.want)
			}
		})
	}
}

func TestGroupDigits(t *testing.T) {
	tests := []struct {
		cents int
		want  string
	}{
		{cents: 123450, want: "1,234.50"},
		{cents: 5, want: "0.05"},
		{cents: 99, want: "0.99"},
		{cents: 0, want: "0.00"},
		{cents: 100, want: "1.00"},
		{cents: 99999, want: "999.99"},
		{cents: 100000, want: "1,000.00"},
		{cents: 12345678, want: "123,456.78"},
		{cents: 100000000, want: "1,000,000.00"},
		{cents: -123450, want: "-1,234.50"},
		{cents: -7, want: "-0.07"},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.cents), func(t *testing.T) {
			got := GroupDigits(tt.cents)
			if got != tt.want {
				t.Errorf("GroupDigits(%d) = %q, want %q", tt.cents, got, tt.want)
			}
		})
	}
}

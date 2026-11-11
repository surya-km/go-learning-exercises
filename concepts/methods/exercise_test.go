package methods

import "testing"

func TestBillIsDueBy(t *testing.T) {
	tests := []struct {
		name string
		bill Bill
		day  int
		want bool
	}{
		{name: "due earlier", bill: Bill{Name: "Rent", DueDay: 1}, day: 10, want: true},
		{name: "due same day", bill: Bill{Name: "Phone", DueDay: 15}, day: 15, want: true},
		{name: "due later", bill: Bill{Name: "Gym", DueDay: 28}, day: 27, want: false},
		{name: "end of month", bill: Bill{Name: "Card", DueDay: 31}, day: 31, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bill.IsDueBy(tt.day); got != tt.want {
				t.Errorf("%+v.IsDueBy(%d) = %t, want %t", tt.bill, tt.day, got, tt.want)
			}
		})
	}
}

func TestBillLabel(t *testing.T) {
	tests := []struct {
		bill Bill
		want string
	}{
		{bill: Bill{Name: "Rent", AmountCents: 120050}, want: "Rent: $1200.50"},
		{bill: Bill{Name: "Water", AmountCents: 7}, want: "Water: $0.07"},
		{bill: Bill{Name: "Streaming", AmountCents: 1599}, want: "Streaming: $15.99"},
		{bill: Bill{Name: "Free trial", AmountCents: 0}, want: "Free trial: $0.00"},
		{bill: Bill{Name: "Insurance", AmountCents: 10000}, want: "Insurance: $100.00"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.bill.Label(); got != tt.want {
				t.Errorf("%+v.Label() = %q, want %q", tt.bill, got, tt.want)
			}
		})
	}
}

func TestBillsTotalDueBy(t *testing.T) {
	bills := Bills{
		{Name: "Rent", AmountCents: 120000, DueDay: 1},
		{Name: "Phone", AmountCents: 4500, DueDay: 15},
		{Name: "Gym", AmountCents: 3000, DueDay: 28},
	}

	tests := []struct {
		name  string
		bills Bills
		day   int
		want  int64
	}{
		{name: "nothing due yet", bills: bills, day: 0, want: 0},
		{name: "only rent", bills: bills, day: 1, want: 120000},
		{name: "rent and phone", bills: bills, day: 15, want: 124500},
		{name: "all bills", bills: bills, day: 31, want: 127500},
		{name: "empty list", bills: Bills{}, day: 31, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bills.TotalDueBy(tt.day); got != tt.want {
				t.Errorf("%+v.TotalDueBy(%d) = %d, want %d", tt.bills, tt.day, got, tt.want)
			}
		})
	}
}

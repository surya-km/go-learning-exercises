package interfaces

import "testing"

// These lines fail to compile if a type does not satisfy Payment.
var (
	_ Payment = Subscription{}
	_ Payment = Loan{}
	_ Payment = Utility{}
)

func TestPaymentMethods(t *testing.T) {
	tests := []struct {
		name        string
		payment     Payment
		wantMonthly int64
		wantName    string
	}{
		{name: "monthly subscription", payment: Subscription{Service: "Streamly", PriceCents: 1599}, wantMonthly: 1599, wantName: "Streamly"},
		{name: "yearly subscription", payment: Subscription{Service: "CloudBox", PriceCents: 12000, Yearly: true}, wantMonthly: 1000, wantName: "CloudBox"},
		{name: "yearly subscription rounds down", payment: Subscription{Service: "News", PriceCents: 9999, Yearly: true}, wantMonthly: 833, wantName: "News"},
		{name: "loan", payment: Loan{Lender: "Car", PrincipalCents: 1200000, Months: 48}, wantMonthly: 25000, wantName: "Car loan"},
		{name: "loan rounds down", payment: Loan{Lender: "Laptop", PrincipalCents: 100000, Months: 3}, wantMonthly: 33333, wantName: "Laptop loan"},
		{name: "loan with zero months", payment: Loan{Lender: "Odd", PrincipalCents: 5000, Months: 0}, wantMonthly: 0, wantName: "Odd loan"},
		{name: "utility average", payment: Utility{Provider: "PowerCo", ReadingsCents: []int64{8000, 9000, 10000}}, wantMonthly: 9000, wantName: "PowerCo"},
		{name: "utility average rounds down", payment: Utility{Provider: "WaterCo", ReadingsCents: []int64{3000, 3001}}, wantMonthly: 3000, wantName: "WaterCo"},
		{name: "utility with no readings", payment: Utility{Provider: "GasCo"}, wantMonthly: 0, wantName: "GasCo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.payment.MonthlyCents(); got != tt.wantMonthly {
				t.Errorf("%#v.MonthlyCents() = %d, want %d", tt.payment, got, tt.wantMonthly)
			}
			if got := tt.payment.Name(); got != tt.wantName {
				t.Errorf("%#v.Name() = %q, want %q", tt.payment, got, tt.wantName)
			}
		})
	}
}

func TestMonthlyTotal(t *testing.T) {
	tests := []struct {
		name     string
		payments []Payment
		want     int64
	}{
		{
			name: "one of each type",
			payments: []Payment{
				Subscription{Service: "Streamly", PriceCents: 1599},
				Loan{Lender: "Car", PrincipalCents: 1200000, Months: 48},
				Utility{Provider: "PowerCo", ReadingsCents: []int64{8000, 10000}},
			},
			want: 1599 + 25000 + 9000,
		},
		{
			name:     "same type twice",
			payments: []Payment{Subscription{Service: "A", PriceCents: 500}, Subscription{Service: "B", PriceCents: 1200, Yearly: true}},
			want:     600,
		},
		{name: "empty", payments: []Payment{}, want: 0},
		{name: "nil", payments: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthlyTotal(tt.payments); got != tt.want {
				t.Errorf("MonthlyTotal(%#v) = %d, want %d", tt.payments, got, tt.want)
			}
		})
	}
}

func TestBiggest(t *testing.T) {
	tests := []struct {
		name     string
		payments []Payment
		want     string
	}{
		{
			name: "loan is biggest",
			payments: []Payment{
				Subscription{Service: "Streamly", PriceCents: 1599},
				Loan{Lender: "Car", PrincipalCents: 1200000, Months: 48},
				Utility{Provider: "PowerCo", ReadingsCents: []int64{8000, 10000}},
			},
			want: "Car loan",
		},
		{
			name: "yearly price is compared per month",
			payments: []Payment{
				Subscription{Service: "Yearly", PriceCents: 6000, Yearly: true},
				Subscription{Service: "Monthly", PriceCents: 600},
			},
			want: "Monthly",
		},
		{
			name: "tie goes to the first",
			payments: []Payment{
				Utility{Provider: "WaterCo", ReadingsCents: []int64{2500}},
				Subscription{Service: "Gym", PriceCents: 2500},
			},
			want: "WaterCo",
		},
		{name: "single payment", payments: []Payment{Loan{Lender: "Bike", PrincipalCents: 60000, Months: 12}}, want: "Bike loan"},
		{name: "empty", payments: nil, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Biggest(tt.payments); got != tt.want {
				t.Errorf("Biggest(%#v) = %q, want %q", tt.payments, got, tt.want)
			}
		})
	}
}

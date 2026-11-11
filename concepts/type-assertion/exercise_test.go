package typeassertion

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		name  string
		event any
		want  string
	}{
		{name: "deposit", event: Deposit{From: "Acme", AmountCents: 250000}, want: "deposit of 250000 cents from Acme"},
		{name: "ATM withdrawal", event: Withdrawal{ATM: true, AmountCents: 2000}, want: "ATM withdrawal of 2000 cents"},
		{name: "other withdrawal", event: Withdrawal{ATM: false, AmountCents: 2000}, want: "withdrawal of 2000 cents"},
		{name: "fee", event: Fee{Reason: "overdraft", AmountCents: 3500}, want: "fee of 3500 cents (overdraft)"},
		{name: "string note", event: "hello", want: "note: hello"},
		{name: "nil", event: nil, want: "empty event"},
		{name: "int", event: 42, want: "unknown event of type int"},
		{name: "float", event: 1.5, want: "unknown event of type float64"},
		{name: "pointer to deposit is not a deposit", event: &Deposit{From: "Acme"}, want: "unknown event of type *typeassertion.Deposit"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Describe(tt.event); got != tt.want {
				t.Errorf("Describe(%#v) = %q, want %q", tt.event, got, tt.want)
			}
		})
	}
}

func TestLookupString(t *testing.T) {
	payload := map[string]any{
		"merchant": "Corner Cafe",
		"empty":    "",
		"amount":   int64(450),
		"tags":     []string{"food"},
	}

	tests := []struct {
		name    string
		payload map[string]any
		key     string
		want    string
		wantOK  bool
	}{
		{name: "string value", payload: payload, key: "merchant", want: "Corner Cafe", wantOK: true},
		{name: "empty string still counts", payload: payload, key: "empty", want: "", wantOK: true},
		{name: "wrong type int64", payload: payload, key: "amount", want: "", wantOK: false},
		{name: "wrong type slice", payload: payload, key: "tags", want: "", wantOK: false},
		{name: "missing key", payload: payload, key: "city", want: "", wantOK: false},
		{name: "nil map", payload: nil, key: "merchant", want: "", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := LookupString(tt.payload, tt.key)
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("LookupString(%v, %q) = (%q, %t), want (%q, %t)", tt.payload, tt.key, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

func TestLookupCents(t *testing.T) {
	payload := map[string]any{
		"amount":   int64(450),
		"zero":     int64(0),
		"refund":   int64(-125),
		"count":    3,
		"text":     "500",
		"merchant": "Corner Cafe",
	}

	tests := []struct {
		name    string
		payload map[string]any
		key     string
		want    int64
		wantOK  bool
	}{
		{name: "int64 value", payload: payload, key: "amount", want: 450, wantOK: true},
		{name: "zero still counts", payload: payload, key: "zero", want: 0, wantOK: true},
		{name: "negative int64", payload: payload, key: "refund", want: -125, wantOK: true},
		{name: "plain int is not int64", payload: payload, key: "count", want: 0, wantOK: false},
		{name: "numeric string is not int64", payload: payload, key: "text", want: 0, wantOK: false},
		{name: "missing key", payload: payload, key: "tip", want: 0, wantOK: false},
		{name: "nil map", payload: nil, key: "amount", want: 0, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := LookupCents(tt.payload, tt.key)
			if got != tt.want || ok != tt.wantOK {
				t.Errorf("LookupCents(%v, %q) = (%d, %t), want (%d, %t)", tt.payload, tt.key, got, ok, tt.want, tt.wantOK)
			}
		})
	}
}

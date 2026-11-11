package errorwrapping

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestProcessPaymentSuccess(t *testing.T) {
	p := Payment{ID: "p1", CardID: "card-ok", AmountCents: 1999}
	if err := ProcessPayment(p); err != nil {
		t.Errorf("ProcessPayment(%+v) = %v, want nil", p, err)
	}
}

func TestProcessPaymentErrors(t *testing.T) {
	tests := []struct {
		name    string
		payment Payment
		wantIs  error // sentinel to find with errors.Is (nil if not a sentinel)
		wantMsg string
	}{
		{
			name:    "zero amount",
			payment: Payment{ID: "p2", CardID: "card-ok", AmountCents: 0},
			wantIs:  ErrInvalidAmount,
			wantMsg: "payment p2: invalid amount",
		},
		{
			name:    "negative amount is checked before the card",
			payment: Payment{ID: "p3", CardID: "card-expired", AmountCents: -5},
			wantIs:  ErrInvalidAmount,
			wantMsg: "payment p3: invalid amount",
		},
		{
			name:    "expired card",
			payment: Payment{ID: "p4", CardID: "card-expired", AmountCents: 500},
			wantIs:  ErrCardExpired,
			wantMsg: "payment p4: gateway: card card-expired: card expired",
		},
		{
			name:    "unknown card",
			payment: Payment{ID: "p5", CardID: "card-lost", AmountCents: 500},
			wantIs:  ErrUnknownCard,
			wantMsg: `payment p5: gateway: unknown card "card-lost"`,
		},
		{
			name:    "declined card",
			payment: Payment{ID: "p6", CardID: "card-poor", AmountCents: 120000},
			wantMsg: "payment p6: gateway: card card-poor: card declined (insufficient_funds) for 120000 cents",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ProcessPayment(tt.payment)
			if err == nil {
				t.Fatalf("ProcessPayment(%+v) = nil, want error %q", tt.payment, tt.wantMsg)
			}
			if tt.wantIs != nil && !errors.Is(err, tt.wantIs) {
				t.Errorf("ProcessPayment(%+v) = %v, want an error matching %v", tt.payment, err, tt.wantIs)
			}
			if err.Error() != tt.wantMsg {
				t.Errorf("ProcessPayment(%+v) message = %q, want %q", tt.payment, err.Error(), tt.wantMsg)
			}
		})
	}
}

func TestDeclineCode(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode string
		wantOK   bool
	}{
		{name: "declined error itself", err: &DeclinedError{Code: "do_not_honor"}, wantCode: "do_not_honor", wantOK: true},
		{name: "wrapped twice: insufficient funds", err: ProcessPayment(Payment{ID: "p1", CardID: "card-poor", AmountCents: 100}), wantCode: "insufficient_funds", wantOK: true},
		{name: "wrapped twice: fraud", err: ProcessPayment(Payment{ID: "p2", CardID: "card-fraud", AmountCents: 100}), wantCode: "suspected_fraud", wantOK: true},
		{name: "wrapped three times", err: fmt.Errorf("checkout: %w", fmt.Errorf("payment p3: %w", ChargeCard("card-fraud", 1))), wantCode: "suspected_fraud", wantOK: true},
		{name: "a sentinel, not a declined error", err: fmt.Errorf("payment p4: %w", ErrCardExpired), wantCode: "", wantOK: false},
		{name: "nil", err: nil, wantCode: "", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, ok := DeclineCode(tt.err)
			if code != tt.wantCode || ok != tt.wantOK {
				t.Errorf("DeclineCode(%v) = (%q, %t), want (%q, %t)", tt.err, code, ok, tt.wantCode, tt.wantOK)
			}
		})
	}
}

func TestRootCause(t *testing.T) {
	declined := &DeclinedError{Code: "insufficient_funds", AmountCents: 50}
	plain := errors.New("disk full")

	tests := []struct {
		name string
		err  error
		want error
	}{
		{name: "nil", err: nil, want: nil},
		{name: "not wrapped", err: plain, want: plain},
		{name: "one layer", err: fmt.Errorf("saving: %w", plain), want: plain},
		{name: "two layers to sentinel", err: ProcessPayment(Payment{ID: "p1", CardID: "card-expired", AmountCents: 1}), want: ErrCardExpired},
		{name: "three layers to declined error", err: fmt.Errorf("a: %w", fmt.Errorf("b: %w", fmt.Errorf("c: %w", declined))), want: declined},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RootCause(tt.err); got != tt.want {
				t.Errorf("RootCause(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestProcessBatch(t *testing.T) {
	t.Run("all succeed", func(t *testing.T) {
		batch := []Payment{{ID: "p1", CardID: "card-ok", AmountCents: 100}, {ID: "p2", CardID: "card-ok", AmountCents: 200}}
		if err := ProcessBatch(batch); err != nil {
			t.Errorf("ProcessBatch(%+v) = %v, want nil", batch, err)
		}
	})

	t.Run("empty", func(t *testing.T) {
		if err := ProcessBatch(nil); err != nil {
			t.Errorf("ProcessBatch(nil) = %v, want nil", err)
		}
	})

	t.Run("mixed failures are all reported", func(t *testing.T) {
		batch := []Payment{
			{ID: "p1", CardID: "card-ok", AmountCents: 100},
			{ID: "p2", CardID: "card-expired", AmountCents: 100},
			{ID: "p3", CardID: "card-ok", AmountCents: 0},
			{ID: "p4", CardID: "card-poor", AmountCents: 100},
		}
		err := ProcessBatch(batch)
		if err == nil {
			t.Fatalf("ProcessBatch(%+v) = nil, want an error", batch)
		}
		for _, sentinel := range []error{ErrCardExpired, ErrInvalidAmount} {
			if !errors.Is(err, sentinel) {
				t.Errorf("ProcessBatch(...) = %q, want an error matching %v", err, sentinel)
			}
		}
		if code, ok := DeclineCode(err); !ok || code != "insufficient_funds" {
			t.Errorf("DeclineCode(ProcessBatch(...)) = (%q, %t), want (%q, true)", code, ok, "insufficient_funds")
		}
		msg := err.Error()
		for _, id := range []string{"payment p2:", "payment p3:", "payment p4:"} {
			if !strings.Contains(msg, id) {
				t.Errorf("ProcessBatch(...) message %q should mention %q", msg, id)
			}
		}
		if strings.Contains(msg, "payment p1:") {
			t.Errorf("ProcessBatch(...) message %q should not mention the successful payment p1", msg)
		}
		if lines := strings.Count(msg, "\n") + 1; lines != 3 {
			t.Errorf("ProcessBatch(...) message has %d lines, want 3 (one per failure):\n%s", lines, msg)
		}
	})
}

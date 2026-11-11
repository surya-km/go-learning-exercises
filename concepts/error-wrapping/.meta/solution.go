// Package errorwrapping practices the error wrapping concept. Read CONCEPT.md first.
package errorwrapping

import (
	"errors"
	"fmt"
)

// ProcessPayment validates the payment and charges the card. Errors are
// wrapped with the prefix "payment <ID>: ".
func ProcessPayment(p Payment) error {
	if p.AmountCents <= 0 {
		return fmt.Errorf("payment %s: %w", p.ID, ErrInvalidAmount)
	}
	if err := ChargeCard(p.CardID, p.AmountCents); err != nil {
		return fmt.Errorf("payment %s: %w", p.ID, err)
	}
	return nil
}

// DeclineCode finds a *DeclinedError anywhere inside err and returns its Code.
func DeclineCode(err error) (string, bool) {
	var de *DeclinedError
	if errors.As(err, &de) {
		return de.Code, true
	}
	return "", false
}

// RootCause unwraps err layer by layer and returns the innermost error.
func RootCause(err error) error {
	for {
		inner := errors.Unwrap(err)
		if inner == nil {
			return err
		}
		err = inner
	}
}

// ProcessBatch processes every payment and joins all failures into one error.
// It returns nil if nothing failed.
func ProcessBatch(payments []Payment) error {
	var errs []error
	for _, p := range payments {
		if err := ProcessPayment(p); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

# Error wrapping

## In plain English

When an error travels up through several layers of code, each layer can add a short
note saying what it was doing, like a parcel collecting shipping stickers. **Wrapping**
adds those notes while keeping the original error inside, so code at the top can still
open the parcel and ask "was this a declined card?" and read the details.

The `errors` concept covered `%w` and `errors.Is` with a sentinel. This one goes further:
several layers, custom error **types** with `errors.As`, `errors.Unwrap` and `errors.Join`.

## Several layers of wrapping

```go
// bottom layer
err := ErrCardExpired
// middle layer
err = fmt.Errorf("gateway: card %s: %w", card, err)
// top layer
err = fmt.Errorf("payment %s: %w", id, err)

fmt.Println(err)                       // payment p1: gateway: card c9: card expired
errors.Is(err, ErrCardExpired)         // true: Is looks through every layer
```

## Custom error types and `errors.As`

A sentinel can only say *which* error happened. A custom type can carry **details**.
Any type with an `Error() string` method is an error:

```go
type DeclinedError struct {
    Code string
}

func (e *DeclinedError) Error() string { return "card declined: " + e.Code }
```

`errors.As` searches the wrapped layers for an error of a given type and, if it finds
one, stores it in your variable:

```go
var de *DeclinedError                // note: a pointer, because Error has a pointer receiver
if errors.As(err, &de) {             // pass the ADDRESS of your variable
    fmt.Println(de.Code)             // "insufficient_funds"
}
```

Rule of thumb: `errors.Is` for "is it this exact error?", `errors.As` for "is it this
*kind* of error, and give me its details".

## `errors.Unwrap`: peel off one layer

```go
inner := errors.Unwrap(err)   // the error that was wrapped with %w, or nil if none
```

Call it in a loop to dig down to the original error.

## `errors.Join`: several errors in one

```go
err := errors.Join(err1, err2, err3)   // nil values are skipped; all nil → returns nil
errors.Is(err, ErrCardExpired)         // true if ANY of them matches
fmt.Println(err)                       // each message on its own line
```

(Note: `errors.Unwrap` does **not** look inside a joined error. `errors.Is` and `errors.As` do.)

> **Coming from Ruby?** Wrapping is like Ruby's `exception.cause` chain, and
> `errors.As` is like `rescue DeclinedError => e`, except it's a normal `if`.

## Your task

`types.go` contains a fake card gateway, `ChargeCard(cardID, cents)`, plus the
`Payment` struct, the `DeclinedError` type and the sentinel errors `ErrInvalidAmount`,
`ErrCardExpired` and `ErrUnknownCard`. Read it first: `ChargeCard` already wraps its
errors once (`"gateway: ..."`).

**Problem 1: `ProcessPayment(p Payment) error`**

- If `p.AmountCents <= 0`, return an error that wraps `ErrInvalidAmount` (don't call the gateway).
- Otherwise call `ChargeCard(p.CardID, p.AmountCents)`. If it fails, return its error
  wrapped one more time.
- In both cases the message must start with `"payment <ID>: "`, for example
  `payment p7: gateway: card card-expired: card expired`.
- On success return `nil`.

**Problem 2: inspecting errors**

- `DeclineCode(err error) (string, bool)`: use `errors.As` to find a `*DeclinedError`
  anywhere in `err`. Return its `Code` and `true`, or `""` and `false`.
- `RootCause(err error) error`: use `errors.Unwrap` in a loop to return the innermost
  error. `RootCause(nil)` is `nil`; an error that wraps nothing is its own root cause.
- `ProcessBatch(payments []Payment) error`: process every payment (keep going after a
  failure) and return all the failures combined with `errors.Join`. Return `nil` if
  every payment succeeds or the list is empty.

Run: `go test -v ./concepts/error-wrapping/`

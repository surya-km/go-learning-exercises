# Interfaces

## In plain English

An **interface** is a list of things a value must be able to do. Any type that has
those methods automatically "fits" the interface; you never have to say so out loud.
That lets one function work with many different types, as long as they can all do the
job it needs. Think of a power socket: it doesn't care if you plug in a lamp or a phone
charger, only that the plug has the right shape.

```go
type Shape interface {
    Area() float64
}

type Square struct{ Side float64 }
func (s Square) Area() float64 { return s.Side * s.Side }

type Circle struct{ R float64 }
func (c Circle) Area() float64 { return 3.14159 * c.R * c.R }

// Works with ANY Shape: Square, Circle, or something written next year.
func TotalArea(shapes []Shape) float64 {
    total := 0.0
    for _, s := range shapes {
        total += s.Area()
    }
    return total
}
```

There is no `implements` keyword. `Square` is a `Shape` simply because it has an
`Area() float64` method. This is called *implicit* implementation.

> **Coming from Ruby?** This is duck typing ("if it quacks like a duck..."), but checked
> by the compiler: passing a type that lacks a method is a compile error, not a crash at runtime.

## Small interfaces are best

Most Go interfaces have just one or two methods, for example `fmt.Stringer`
(`String() string`) and `io.Reader` (`Read([]byte) (int, error)`).

## Your task

A budgeting app has different kinds of recurring payments. In `types.go` you'll find the
`Payment` interface and three types: `Subscription`, `Loan` and `Utility`.

```go
type Payment interface {
    MonthlyCents() int64
    Name() string
}
```

**Problem 1: make all three types satisfy `Payment`**

Implement `MonthlyCents()` and `Name()` for each type in `exercise.go`:

| Type | `MonthlyCents()` | `Name()` |
|---|---|---|
| `Subscription{Service, PriceCents, Yearly}` | `PriceCents` if monthly; if `Yearly` is true, `PriceCents / 12` (integer division) | `Service` |
| `Loan{Lender, PrincipalCents, Months}` | `PrincipalCents / Months` (integer division); `0` if `Months <= 0` | `Lender + " loan"` |
| `Utility{Provider, ReadingsCents}` | average of the `ReadingsCents` slice (integer division); `0` if empty | `Provider` |

**Problem 2: `MonthlyTotal(payments []Payment) int64` and `Biggest(payments []Payment) string`**

- `MonthlyTotal` adds up `MonthlyCents()` of every payment. Empty list → `0`.
- `Biggest` returns the `Name()` of the payment with the largest `MonthlyCents()`.
  If two tie, return the one that comes **first**. Empty list → `""`.

Both functions take the interface, so they must not care which concrete type each payment is.

Run: `go test -v ./concepts/interfaces/`

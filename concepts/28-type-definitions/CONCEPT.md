# Type definitions

## In plain English

You can create your **own named type** on top of an existing one, like
`type Cents int64`. Underneath it's still a number, but Go now treats it as a different
type, so you can't accidentally mix up "cents" with "number of days". You can also give
your new type its own methods, so the money logic lives right on the money type.

```go
type Cents int64

func (c Cents) Dollars() int64 { return int64(c) / 100 }

price := Cents(1999)
price.Dollars()   // 19
```

## Different types don't mix without a conversion

```go
var c Cents = 500
var n int64 = 3

// c * n        // compile error: mismatched types Cents and int64
c * Cents(n)    // OK: convert explicitly
int64(c)        // back to a plain int64
c + 25          // OK: untyped constants like 25 adapt to Cents
```

This is the main benefit: the compiler catches mixed-up units for you.

## Named string types and constants

A named type built on `string` works well for a fixed set of options:

```go
type Frequency string

const (
    Weekly  Frequency = "weekly"
    Monthly Frequency = "monthly"
)

func (f Frequency) IsValid() bool {
    return f == Weekly || f == Monthly
}
```

> **Coming from Ruby?** There's no direct equivalent. It's a bit like wrapping a value in
> a tiny class (a "value object"), but with zero runtime cost: a `Cents` is stored
> exactly like an `int64`.

## Your task

`Cents`, `Frequency` and the `Weekly`/`Monthly`/`Quarterly`/`Yearly` constants are
declared in `types.go`. Implement the methods and function in `exercise.go`.

**Problem 1: methods on `Cents`**

- `func (c Cents) Format() string` returns the amount as dollars:
  `Cents(1234).Format()` → `"$12.34"`, `Cents(5).Format()` → `"$0.05"`,
  `Cents(0).Format()` → `"$0.00"`, `Cents(-1250).Format()` → `"-$12.50"`.
- `func (c Cents) Split(n int) []Cents` splits the amount into `n` parts as evenly as
  possible, so the parts always add up to the original. Leftover cents go to the
  **first** parts, one each: `Cents(1000).Split(3)` → `[334 333 333]`.
  If `n <= 0`, return `nil`. You only need to handle `c >= 0`.

**Problem 2: `Frequency`**

- `func (f Frequency) TimesPerYear() int` returns `52` for `Weekly`, `12` for `Monthly`,
  `4` for `Quarterly`, `1` for `Yearly`, and `0` for anything else.
- `YearlyCost(amount Cents, f Frequency) Cents` returns what a recurring payment
  of `amount` at frequency `f` costs per year: `YearlyCost(1599, Monthly)` → `19188`.

Run: `go test -v ./concepts/28-type-definitions/`

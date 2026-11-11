# Floating-point numbers

## In plain English

A `float64` stores numbers with a decimal point, like `3.14` or `0.1`. The computer keeps
them in binary, and most decimal fractions (even `0.1`) can't be stored exactly, so you get
a tiny error. That's fine for distances or percentages, but it's why we **never** store
money in a `float64`.

```go
var price float64 = 19.99
rate := 0.05            // Go works out that rate is a float64

price * rate            // 0.9995000000000001  (not exactly 0.9995!)
0.1 + 0.2               // 0.30000000000000004
0.1+0.2 == 0.3          // false
```

## Comparing floats: use a tolerance

Because of these tiny errors, don't compare floats with `==`. Instead, check whether
they're "close enough":

```go
import "math"

math.Abs(a-b) < 1e-9    // 1e-9 means 0.000000001
```

## Rounding

```go
math.Round(2.5)         // 3    (halves round away from zero)
math.Round(-2.5)        // -3
math.Floor(2.7)         // 2
math.Ceil(2.1)          // 3
int(math.Round(1234.6)) // 1235 (convert to int AFTER rounding)
```

`int(2.9)` is `2`. Converting a float to an int just chops off the decimal part, so
always round first.

## Floats and integers don't mix

```go
cents := 1999
float64(cents) * 0.05   // you must convert the int first
```

> **Coming from Ruby?** Ruby's `Float` is the same thing as Go's `float64`, with the same
> `0.1 + 0.2` surprise. Go has no `BigDecimal` built in, so integer cents is the usual
> answer for money.

## Your task

**Problem 1:** implement **`ApplyRate(cents int, ratePercent float64) int`**.

Interest rates and tax rates are often given as floats like `5.25` (meaning 5.25%). Work
out `cents * ratePercent / 100` and **round to the nearest cent**. Halves round away from
zero, which is what `math.Round` does.

| Call | Result |
|---|---|
| `ApplyRate(10000, 5)` | `500` |
| `ApplyRate(1999, 8.25)` | `165` (164.9175 rounds up) |
| `ApplyRate(150, 1)` | `2` (1.5 rounds up) |
| `ApplyRate(1000, 0)` | `0` |
| `ApplyRate(-1000, 2.5)` | `-25` (refunds work too) |

**Problem 2:** implement **`AlmostEqual(a, b, tolerance float64) bool`**.

Return `true` if `a` and `b` differ by **at most** `tolerance`.

| Call | Result |
|---|---|
| `AlmostEqual(0.1+0.2, 0.3, 1e-9)` | `true` |
| `AlmostEqual(1.0, 1.1, 0.01)` | `false` |
| `AlmostEqual(5, 3, 2)` | `true` (exactly at the tolerance) |
| `AlmostEqual(-1.0, 1.0, 1.5)` | `false` |

Run: `go test -v ./concepts/floating-point-numbers/`

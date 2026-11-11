# Arithmetic Operators

## In plain English

Go does maths with the usual symbols: `+`, `-`, `*`, `/` and `%` (remainder). The one
surprise is division: when both numbers are whole numbers (`int`), Go **throws away the
fraction**, so `7 / 2` is `3`, not `3.5`. That's why we store money as whole cents.

```go
a, b := 17, 5
a + b   // 22
a - b   // 12
a * b   // 85
a / b   // 3   (the .4 is dropped, it does NOT round)
a % b   // 2   (the remainder: 17 = 3*5 + 2)
```

## Shortcuts

```go
total := 100
total += 50   // same as total = total + 50  → 150
total -= 20   // 130
total *= 2    // 260
total /= 4    // 65
total++       // 66  (++ and -- are statements, not expressions: you can't write x = y++)
```

## Rounding with integer maths

Since `/` always rounds *down* (for positive numbers), you can round to the nearest
whole number by adding half the divisor first:

```go
// 1234 / 100 = 12, but we want 12.34 rounded to the nearest whole → 12
(1234 + 50) / 100   // 12
(1250 + 50) / 100   // 13   (a half rounds up)
```

You can't mix types: `int + float64` is a compile error. Convert explicitly with
`float64(x)` or `int(y)`. In this exercise everything is `int`, so you won't need to.

> **Coming from Ruby?** Same as Ruby: `7 / 2` is `3` for integers. But Go has no
> `7.fdiv(2)` shortcut and won't auto-convert. Ruby's `**` (power) doesn't exist either.

## Your task

All amounts are in integer **cents** (`$12.34` is `1234`).

**Problem 1.** Implement `TipCents(billCents, percent int) int`. Return the tip for a bill,
rounded to the **nearest** cent (exactly half a cent rounds up).

- `TipCents(1000, 15)` returns `150`.
- `TipCents(1999, 20)` returns `400` (the exact tip is 399.8 cents).
- `TipCents(1234, 18)` returns `222` (exact: 222.12).
- `TipCents(50, 1)` returns `1` (exact: 0.5, rounds up).

Inputs are never negative.

**Problem 2.** Implement `RoundUp(cents int) (roundedCents, spareChangeCents int)`. Some
banking apps "round up" each purchase to the next whole dollar and save the difference.

- `RoundUp(1234)` returns `1300, 66`.
- `RoundUp(1300)` returns `1300, 0` (already a whole dollar, nothing to save).
- `RoundUp(1)` returns `100, 99`.
- `RoundUp(0)` returns `0, 0`.

Inputs are never negative. Hint: `%` tells you how far past the last whole dollar you are.

Run: `go test -v ./concepts/arithmetic-operators/`

# Multiple Return Values

## In plain English

A Go function can hand back **more than one result at once**. Instead of packing answers
into an array or an object, you just list them: "here's the quotient *and* the remainder".
The caller receives them into separate variables in one line.

```go
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

q, r := divmod(7, 2) // q = 3, r = 1
```

Use `_` to ignore a value you don't need (Go refuses to compile unused variables):

```go
_, r := divmod(9, 4) // r = 1
```

> **Coming from Ruby?** It looks like `q, r = 7.divmod(2)`, but Go doesn't build a
> hidden array. The values really are separate, and their number and types are fixed.

## Named results

You can give results names. They start at their zero value (`0`, `""`, `false`), act like
local variables, and also document what each result means:

```go
func minMax(a, b int) (low, high int) {
    if a < b {
        return a, b
    }
    return b, a
}
```

A bare `return` (with nothing after it) returns the named results as they are. It's
handy in tiny functions, but in longer ones `return low, high` is clearer.

## The "result plus ok" pattern

A very common Go shape is `(value, ok bool)`, meaning "here's the answer, and whether it
worked". You'll see the same idea later with `(value, error)`.

```go
func safeDivide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

if result, ok := safeDivide(10, 0); !ok {
    fmt.Println("can't divide by zero")
}
```

## Useful operators

`/` on ints drops the remainder (`1000 / 3 == 333`) and `%` gives the remainder
(`1000 % 3 == 1`).

## Your task

Implement in `exercise.go`:

**Problem 1: `SplitBill(totalCents, people int) (each, leftover int)`**

Split a bill evenly. Return how many cents each person pays and how many cents are left
over. If `people` is 0 or less, return `0, totalCents` (nobody to split with).

| Call | Result |
|---|---|
| `SplitBill(1000, 4)` | `250, 0` |
| `SplitBill(1000, 3)` | `333, 1` |
| `SplitBill(2, 5)` | `0, 2` |
| `SplitBill(500, 0)` | `0, 500` |

**Problem 2: `CoinChange(cents int) (quarters, dimes, nickels, pennies int)`**

A cashier gives change using as few coins as possible: quarters (25¢) first, then dimes
(10¢), then nickels (5¢), then pennies (1¢). Return how many of each coin to hand over.
If `cents` is 0 or negative, return all zeros.

| Call | Result |
|---|---|
| `CoinChange(41)` | `1, 1, 1, 1` |
| `CoinChange(99)` | `3, 2, 0, 4` |
| `CoinChange(30)` | `1, 0, 1, 0` |
| `CoinChange(-5)` | `0, 0, 0, 0` |

> Why cents? Money should almost never be stored as a floating-point number,
> because `0.1 + 0.2 != 0.3` in floating point. Integer cents are exact.

Run: `go test -v ./concepts/15-multiple-return-values/`

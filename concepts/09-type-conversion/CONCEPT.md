# Type conversion

## In plain English

Go never quietly changes one type into another. If you have an `int` and need a
`float64`, you have to say so: `float64(n)`. Turning text into numbers (and numbers into
text) is a different job, and the `strconv` package ("string conversion") does it.

## Number to number: `T(value)`

```go
count := 3
total := 1000

total / count                    // 333   (int / int drops the remainder)
float64(total) / float64(count)  // 333.3333333333333

price := 19.99
int(price)                       // 19    (the decimal part is chopped off, NOT rounded)
```

You can't mix types in one expression. `float64(total) / count` won't compile, because
both sides of `/` must be the same type.

## Text to number and back: `strconv`

```go
import "strconv"

n, err := strconv.Atoi("42")   // 42, nil       ("ASCII to integer")
_, err = strconv.Atoi("4x2")   // 0, an error   (always check err!)
s := strconv.Itoa(42)          // "42"          ("integer to ASCII")

f, err := strconv.ParseFloat("2.5", 64)  // 2.5, nil
b, err := strconv.ParseBool("true")      // true, nil
```

**Trap:** `string(65)` gives `"A"` (the character with code 65), **not** `"65"`. Use
`strconv.Itoa` to turn a number into its digits. `go vet` warns about this mistake.

> **Coming from Ruby?** There's no `"42".to_i` that quietly returns `0` for junk input.
> `strconv.Atoi` returns an error you must check. And `1 + "2"` doesn't compile at all.

## Your task

**Problem 1:** implement **`AverageCents(totalCents, count int) float64`**.

Return the exact average (`totalCents` divided by `count`) as a `float64`, keeping the
decimal part. If `count` is 0 or less, return `0`.

| Call | Result |
|---|---|
| `AverageCents(1000, 4)` | `250` |
| `AverageCents(1000, 3)` | `333.333...` |
| `AverageCents(5, 2)` | `2.5` (not `2`!) |
| `AverageCents(700, 0)` | `0` |

**Problem 2:** implement **`OrderSummary(quantity, unitCents string) (string, error)`**.

Both values arrive as text, as they would from a web form or a CSV file. Convert them to
numbers, multiply, and return a summary like `"3 items, 1050 cents"`.

- Use `"item"` (singular) when the quantity is exactly 1: `"1 item, 499 cents"`.
- Surrounding spaces are **not** allowed. Pass the strings to `strconv.Atoi` as they are.
- If either string isn't a valid integer, return `""` and the error from `strconv.Atoi`.

| Call | Result |
|---|---|
| `OrderSummary("3", "350")` | `"3 items, 1050 cents", nil` |
| `OrderSummary("1", "499")` | `"1 item, 499 cents", nil` |
| `OrderSummary("0", "250")` | `"0 items, 0 cents", nil` |
| `OrderSummary("two", "350")` | `"", <error>` |
| `OrderSummary("2", "3.50")` | `"", <error>` |

Run: `go test -v ./concepts/09-type-conversion/`

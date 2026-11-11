# String formatting

## In plain English

`fmt.Sprintf` builds a string from a template with placeholders, called **verbs**, that
start with `%`. Each verb is replaced by the next value you pass in. Numbers between the
`%` and the letter control width, padding and decimal places, which is how you line up
columns in a bank statement.

```go
import "fmt"

fmt.Sprintf("%s paid %d cents", "Ada", 450)   // "Ada paid 450 cents"
fmt.Sprintf("%q", "hi")                        // "\"hi\""  (quoted string)
fmt.Sprintf("%v", []int{1, 2})                 // "[1 2]"   (any value, default format)
fmt.Sprintf("%t", true)                        // "true"
fmt.Sprintf("%.2f", 3.14159)                   // "3.14"    (2 decimal places)
fmt.Sprintf("%d%%", 50)                        // "50%"     (%% is a literal percent sign)
```

## Width and padding

```go
fmt.Sprintf("[%5d]", 42)     // "[   42]"   right-aligned in 5 characters
fmt.Sprintf("[%-5d]", 42)    // "[42   ]"   the - means left-align
fmt.Sprintf("[%05d]", 42)    // "[00042]"   pad with zeros
fmt.Sprintf("[%02d]", 7)     // "[07]"      handy for cents and dates
fmt.Sprintf("[%-8s]", "Tea") // "[Tea     ]"
fmt.Sprintf("[%8s]", "Tea")  // "[     Tea]"
```

If the value is wider than the width, nothing is cut off. The width is a minimum.

`fmt.Println` and `fmt.Printf` print to the screen. `fmt.Sprintf` ("S" for string)
returns the text instead, which is what you need in this exercise.

> **Coming from Ruby?** `fmt.Sprintf` is Ruby's `format` / `sprintf` / `"%05d" % 42`. The
> verbs are nearly the same. `%v` is roughly `inspect`/`to_s` for any value.

## Your task

**Problem 1:** implement **`StatementLine(description string, cents int) string`**.

Format one line of a bank statement:

- The description is **left-aligned** in a column 20 characters wide.
- Then a single space.
- Then the amount, **right-aligned** in a column 10 characters wide, written as dollars
  with exactly two decimals and a `$` sign: `$4.50`, `$1234.00`. Negative amounts put the
  minus sign before the `$`: `-$12.05`.
- Don't use `float64` for the amount. Build it from `cents / 100` and `cents % 100`.

```
StatementLine("Coffee", 450)
"Coffee                    $4.50"
 |<------- 20 ------->| |<- 10 ->|

StatementLine("Refund", -1205)
"Refund                  -$12.05"
```

**Problem 2:** implement **`SpendingShare(category string, spentCents, totalCents int) string`**.

Return a line like `"Groceries: 25.0%"`. The percentage is `spentCents / totalCents * 100`,
shown with exactly one decimal place. If `totalCents` is 0, show `0.0%`.

| Call | Result |
|---|---|
| `SpendingShare("Groceries", 2500, 10000)` | `"Groceries: 25.0%"` |
| `SpendingShare("Rent", 2, 3)` | `"Rent: 66.7%"` |
| `SpendingShare("Fun", 0, 5000)` | `"Fun: 0.0%"` |
| `SpendingShare("Misc", 100, 0)` | `"Misc: 0.0%"` |

(Here a float is fine, because a percentage for display isn't money.)

Run: `go test -v ./concepts/string-formatting/`

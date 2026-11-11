# strings.Builder

## In plain English

Go strings can't be changed, so `s += "x"` in a loop copies the whole string every
time. With thousands of pieces that gets slow. A `strings.Builder` is like a notepad you
keep writing onto. You add pieces to it, then call `.String()` once at the end to get the
finished text.

```go
import "strings"

var sb strings.Builder        // ready to use, no setup needed

sb.WriteString("Total: ")     // add a string
sb.WriteByte('$')             // add a single byte (an ASCII character)
sb.WriteRune('€')             // add any character (a rune)
sb.WriteString("12.50")

sb.String()                   // "Total: $€12.50"
sb.Len()                      // number of bytes written so far
```

It works with `fmt.Fprintf` too, which is handy for formatted pieces:

```go
fmt.Fprintf(&sb, "%-10s|", "Coffee")   // note the & (pass a pointer to the builder)
```

Some tips:

- Declare it with `var sb strings.Builder`, then use it. You don't need `new` or `make`.
- If you know roughly how big the result will be, `sb.Grow(n)` reserves space up front.
- Don't copy a Builder after writing to it. Pass `*strings.Builder` (a pointer) to helper
  functions instead.

> **Coming from Ruby?** It's like building a string with `<<` on a mutable buffer, then
> taking the result at the end. Similar to `StringIO`.

## Converting numbers

```go
import "strconv"

strconv.Itoa(42)   // "42"
```

## Your task

Each line is a `LineItem` (declared in `types.go`):

```go
type LineItem struct {
    Name  string
    Cents int
}
```

**Problem 1:** implement **`Receipt(items []LineItem) string`** using a `strings.Builder`.

Return one line per item, then a total line. Each line ends with `"\n"`:

```
Coffee: 450
Bagel: 325
TOTAL: 775
```

That's `Name + ": " + cents` for each item, then `"TOTAL: " + sum`. For an empty list,
return `"TOTAL: 0\n"`.

**Problem 2:** implement **`GroupDigits(cents int) string`** using a `strings.Builder`.

Format a cent amount as dollars with comma thousand separators and exactly two decimal
places:

| Call | Result |
|---|---|
| `GroupDigits(123450)` | `"1,234.50"` |
| `GroupDigits(5)` | `"0.05"` |
| `GroupDigits(99)` | `"0.99"` |
| `GroupDigits(100000000)` | `"1,000,000.00"` |
| `GroupDigits(-123450)` | `"-1,234.50"` |

Hints: split into `dollars := cents / 100` and `rem := cents % 100` (handle the sign
first). Turn `dollars` into a string with `strconv.Itoa`, then write its digits one at a
time, adding a `','` before each group of three counted from the **right**.

Run: `go test -v ./concepts/24-strings-builder/`

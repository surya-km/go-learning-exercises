# Numbers

## In plain English

Go has whole numbers (`int`) and numbers with a decimal point (`float64`). For money we
use `int` and count in **cents**, so $12.34 is stored as `1234`. When you divide one `int`
by another, Go throws the remainder away. The `%` operator gives you that remainder back.

```go
a := 17   // Go works out that a is an int
b := 5

a + b   // 22
a - b   // 12
a * b   // 85
a / b   // 3   (integer division drops the .4)
a % b   // 2   (the remainder: 17 = 3*5 + 2)
```

Handy shortcuts:

```go
total := 0
total += 250   // same as total = total + 250
count := 0
count++        // same as count = count + 1
```

An `int` is 64 bits on modern machines, so it holds values up to about 9.2 quintillion.
That's plenty for cents.

> **Coming from Ruby?** `17 / 5` is `3` in both languages, and `%` works the same for
> positive numbers. Unlike Ruby, Go integers have a fixed size and silently wrap around
> if they overflow, and you can't mix `int` and `float64` in one expression without
> converting first.

## A tiny bit about slices

Problem 1 returns a **slice**, which is Go's growable list (covered properly later):

```go
shares := make([]int, 3)   // [0 0 0]
shares[0] = 5              // [5 0 0]
len(shares)                // 3
```

## Your task

**Problem 1:** implement **`SplitBill(totalCents, people int) []int`**.

Split a restaurant bill as evenly as possible. Return one amount per person.

- The amounts must add up to exactly `totalCents`. No cent gets lost.
- Amounts differ by at most 1 cent. The **first** people in the list pay the extra cents.
- If `people` is 0 or less, return `nil`.

| Call | Result |
|---|---|
| `SplitBill(1000, 4)` | `[250 250 250 250]` |
| `SplitBill(1000, 3)` | `[334 333 333]` |
| `SplitBill(1001, 3)` | `[334 334 333]` |
| `SplitBill(2, 4)` | `[1 1 0 0]` |
| `SplitBill(500, 0)` | `nil` |

**Problem 2:** implement **`MakeChange(cents int) (quarters, dimes, nickels, pennies int)`**.

Give change using the **fewest coins**: quarters (25¢), then dimes (10¢), then nickels (5¢),
then pennies (1¢).

| Call | Result |
|---|---|
| `MakeChange(41)` | `1, 1, 1, 1` |
| `MakeChange(99)` | `3, 2, 0, 4` |
| `MakeChange(0)` | `0, 0, 0, 0` |
| `MakeChange(-5)` | `0, 0, 0, 0` (negative amounts get no change) |

Hint: `/` tells you how many coins fit, and `%` tells you what's left over.

Run: `go test -v ./concepts/06-numbers/`

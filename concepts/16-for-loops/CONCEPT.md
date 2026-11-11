# For Loops

## In plain English

A loop repeats a block of code. Go has **only one loop keyword: `for`**. There is no
`while`, `until` or `loop`. The same `for` covers all of those jobs, depending on what
you put after it.

## 1. The classic counting loop

```go
for i := 0; i < 3; i++ {   // start; keep going while this is true; after each pass
    fmt.Println(i)         // 0, 1, 2
}
```

## 2. The "while" loop: just a condition

```go
balance := 1000
for balance > 0 {          // keeps going while the condition is true
    balance -= 300
}
```

## 3. The forever loop, with `break`

```go
for {
    if done() {
        break              // jump out of the loop right now
    }
}
```

## Skipping a pass with `continue`

```go
for i := 0; i < 5; i++ {
    if i%2 == 0 {
        continue           // skip the rest of this pass, go to the next i
    }
    fmt.Println(i)         // 1, 3
}
```

You can also `return` from inside a loop. It leaves the loop *and* the function.

> **Coming from Ruby?** `while x > 0 ... end` becomes `for x > 0 { ... }`, and
> `loop do ... end` becomes `for { ... }`. `next` is called `continue`. Go's
> `for ... range` (like Ruby's `.each`) has its own exercise.

## Your task

Implement in `exercise.go`:

**Problem 1: `MonthsToPayOff(balanceCents, monthlyRatePercent, paymentCents int) int`**

Count how many months it takes to pay off a loan. Each month:
1. Interest is added: `balanceCents * monthlyRatePercent / 100` (rounded down).
2. Then the payment is subtracted.

Keep going until the balance is 0 or less, and return the number of months.

- If the balance is already 0 or less, return `0`.
- If a payment is not **bigger** than that month's interest, the loan can never be
  paid off. Return `-1` instead of looping forever.

| Call | Result | Why |
|---|---|---|
| `MonthsToPayOff(1000, 0, 300)` | `4` | 700, 400, 100, then paid off |
| `MonthsToPayOff(10000, 1, 2500)` | `5` | 7600, 5176, 2727, 227, then paid off |
| `MonthsToPayOff(10000, 5, 500)` | `-1` | interest (500) eats the whole payment |
| `MonthsToPayOff(0, 3, 100)` | `0` | nothing owed |

A `for balance > 0 { ... }` loop is a natural fit.

**Problem 2: `CountWeekdays(startDay, days int) int`**

Days of the week are numbered `0` = Sunday, `1` = Monday, ..., `6` = Saturday.
Starting on `startDay` (0-6) and counting forward `days` days (including the start
day), return how many of them are weekdays (Monday to Friday). If `days` is 0 or less,
return `0`.

Hint: the day `i` steps after the start is `(startDay + i) % 7`. Use a counting loop,
and `continue` past Saturdays and Sundays.

| Call | Result |
|---|---|
| `CountWeekdays(1, 7)` | `5` (a full week from Monday) |
| `CountWeekdays(6, 2)` | `0` (Saturday, Sunday) |
| `CountWeekdays(5, 3)` | `1` (Friday, Saturday, Sunday) |
| `CountWeekdays(3, 10)` | `8` |

Run: `go test -v ./concepts/16-for-loops/`

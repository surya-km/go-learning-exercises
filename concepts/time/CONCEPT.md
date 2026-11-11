# Time

## In plain English

The `time` package handles dates and times. A `time.Time` is one exact moment, and a
`time.Duration` is a length of time (like "90 minutes"). Go can do date maths for you,
like "add one month" or "how many days between these two dates", which is exactly
what you need for bills and recurring payments.

```go
t := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC) // year, month, day, h, m, s, ns, zone
t.Year(), t.Month(), t.Day()      // 2024, March, 15
t.AddDate(0, 1, 0)                // one month later: 2024-04-15
t.Add(36 * time.Hour)             // 36 hours later
later.Sub(t)                      // a time.Duration between two times
```

Always pass a location. In this exercise everything is in `time.UTC`, so the results
don't depend on the computer's timezone.

## Month-end surprises

Go *normalizes* out-of-range dates instead of failing. That's handy but can surprise you:

```go
time.Date(2024, time.February, 30, 0, 0, 0, 0, time.UTC) // March 1, 2024 (Feb has 29 days)
time.Date(2024, time.March, 0, 0, 0, 0, 0, time.UTC)     // Feb 29, 2024: "day 0" = last day of previous month
jan31.AddDate(0, 1, 0)                                   // March 2 (or 3), NOT Feb 28!
```

The "day 0" trick is the easy way to find how many days a month has.

## Parsing and formatting: the reference date

Go doesn't use `%Y-%m-%d`. Instead you write how one specific date would look:
**Mon Jan 2 15:04:05 2006** (think 1 2 3 4 5 6: month 1, day 2, hour 3pm, minute 4,
second 5, year 2006).

```go
t, err := time.Parse("2006-01-02", "2024-02-29")   // parse; result is in UTC
t.Format("2006-01-02")                             // "2024-02-29"
t.Format("Mon 2 Jan 2006")                         // "Thu 29 Feb 2024"
```

## Days between two dates

```go
d := b.Sub(a)          // time.Duration
days := int(d.Hours() / 24)
```

> **Coming from Ruby?** `time.Date(...)` is like `Time.utc(...)`, and the layout string
> replaces `strftime`'s `%Y-%m-%d`. Ruby's `Date >> 1` clamps Jan 31 to Feb 28; Go's
> `AddDate` does not, so you clamp yourself.

## Your task

All times in the tests are built with `time.Date(..., time.UTC)` at midnight, and your
results should be midnight UTC too.

**Problem 1: monthly bills**

- `DaysInMonth(year int, month time.Month) int` returns how many days the month has
  (`DaysInMonth(2024, time.February)` → `29`, `DaysInMonth(2023, time.February)` → `28`).
- `NextDueDate(after time.Time, dueDay int) time.Time`: a bill is due every month on
  `dueDay` (1 to 31). If a month is too short, it's due on that month's **last** day
  instead (due day 31 in April → April 30). Return the first due date **strictly after**
  `after`.

| `after` | `dueDay` | Result |
|---|---|---|
| 2024-01-15 | 20 | 2024-01-20 |
| 2024-01-20 | 20 | 2024-02-20 (same day doesn't count) |
| 2024-01-31 | 31 | 2024-02-29 (leap year) |
| 2023-01-31 | 31 | 2023-02-28 |
| 2024-02-29 | 31 | 2024-03-31 |
| 2024-12-31 | 5 | 2025-01-05 (rolls into the next year) |

**Problem 2: parsing and formatting**

- `DaysBetween(from, to string) (int, error)` parses both dates in `"2006-01-02"`
  format and returns the number of days from `from` to `to` (negative if `to` is
  earlier). If either string can't be parsed, return `0` and the parse error.
- `FormatDue(t time.Time) string` returns the date like `"Thu 29 Feb 2024"`
  (weekday, day without leading zero, short month, year).

Run: `go test -v ./concepts/time/`

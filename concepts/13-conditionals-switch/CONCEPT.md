# Conditionals: switch

## In plain English

A `switch` is a tidy way to write a long `if / else if` chain. You give it a value, list
the possible matches as `case`s, and Go runs the **first** case that matches. A `default`
case catches everything else.

```go
switch day {
case "Saturday", "Sunday":       // one case can list several values
    fmt.Println("weekend")
case "Friday":
    fmt.Println("almost weekend")
default:
    fmt.Println("weekday")
}
```

Unlike some languages, Go does **not** "fall through" into the next case. Each case stops
by itself, so there is no `break` to forget.

## switch with no value

Leave the value out and each `case` becomes its own true/false condition. This is great
for ranges, and it reads more cleanly than a chain of `else if`:

```go
switch {
case temp < 0:
    return "freezing"
case temp < 20:          // we only get here if temp >= 0
    return "mild"
default:
    return "hot"
}
```

Order matters: Go checks cases from top to bottom.

> **Coming from Ruby?** This is Ruby's `case / when`. `case "a", "b":` is like
> `when "a", "b"`, and a value-less `switch` is like `case` with no subject. Go has no
> ranges like `when 1..5`; write `case n >= 1 && n <= 5:` instead.

## Your task

**Problem 1.** Implement `Category(merchant string) string` using a `switch` on `merchant`:

| Merchant | Category |
|---|---|
| `"Tesco"`, `"Aldi"`, `"Lidl"` | `"Groceries"` |
| `"Shell"`, `"BP"`, `"Uber"` | `"Transport"` |
| `"Netflix"`, `"Spotify"` | `"Entertainment"` |
| anything else (including `""`) | `"Other"` |

Matching is exact and case-sensitive (`"tesco"` is `"Other"`).

**Problem 2.** Implement `CreditRating(score int) string` using a `switch` with no value.
Valid scores are `300` to `850` inclusive:

| Score | Rating |
|---|---|
| below `300` or above `850` | `"Invalid"` |
| `800`–`850` | `"Excellent"` |
| `740`–`799` | `"Very good"` |
| `670`–`739` | `"Good"` |
| `580`–`669` | `"Fair"` |
| `300`–`579` | `"Poor"` |

Run: `go test -v ./concepts/13-conditionals-switch/`

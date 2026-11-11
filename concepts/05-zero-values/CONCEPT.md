# Zero Values

## In plain English

In Go, a variable is **never** left holding random junk. If you create one without giving
it a value, Go fills it with that type's "zero value": a sensible empty default. So you
can always tell what an unset variable contains, and you often don't need to set it at all.

| Type | Zero value |
|---|---|
| `int`, `float64` (all numbers) | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| slices, maps, pointers, errors, functions | `nil` (means "nothing here yet") |

```go
var count int       // 0
var name string     // ""
var done bool       // false
var tags []string   // nil
var totals map[string]int // nil
```

## Slices and maps in one minute

A **slice** is a list: `[]string{"food", "rent"}`. A **map** is a lookup table from keys to
values: `map[string]int{"food": 1200}`. Both have a `nil` zero value, and they behave
differently when `nil`:

```go
var tags []string            // nil slice
len(tags)                    // 0: a nil slice acts like an empty list
for _, t := range tags {}    // fine: loops zero times
tags = append(tags, "food")  // fine: append works on a nil slice

empty := []string{}          // NOT nil, but also length 0
tags == nil                  // you can check for nil explicitly

var totals map[string]int    // nil map
totals["food"]               // 0: reading a missing key gives the zero value
// totals["food"] = 5        // PANIC: you can't write to a nil map!
totals = map[string]int{}    // make it first (or: make(map[string]int))
totals["food"] = 5           // now fine
```

A common Go pattern: let the zero value mean **"not set, use the default"**.

> **Coming from Ruby?** Ruby gives you `nil` for everything unset, and calling a method on
> `nil` blows up. Go's numbers, strings and bools are never `nil`; they start at `0`, `""`
> and `false`. Only slices, maps, pointers and a few others can be `nil`.

## Your task

Money is in integer **cents**.

**Problem 1.** Implement
`ApplyDefaults(currency string, dailyLimitCents int, categories []string) (string, int, []string)`.
It fills in settings the user left unset (at their zero value):

- `currency` is `""` → use `"USD"`. Otherwise keep it.
- `dailyLimitCents` is `0` → use `50000`. Otherwise keep it.
- `categories` is `nil` → use `[]string{"General"}`.
  But an **empty, non-nil** slice (`[]string{}`) means the user deliberately removed every
  category, so keep it as it is (empty). Any other slice is kept too.

**Problem 2.** Implement `AddSpend(totals map[string]int, category string, cents int) map[string]int`.
It adds `cents` to the running total for `category` and returns the map.

- If `totals` is `nil`, create a new map first (writing to a nil map would panic).
- If the category isn't in the map yet, its total starts from the zero value, `0`.
- When `totals` is not nil, update **that same map** and return it.

Run: `go test -v ./concepts/05-zero-values/`

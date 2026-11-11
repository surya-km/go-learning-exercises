# Maps

## In plain English

A **map** is a lookup table: you give it a **key** and it gives you back a **value**,
like looking up a word in a dictionary. It's Go's version of a Ruby `Hash`. All keys
have one type and all values have one type, for example `map[string]int` means
"string keys, int values".

## Creating, reading, writing

```go
totals := map[string]int{}      // empty map, ready to use (or: make(map[string]int))
totals["food"] = 500            // add or overwrite
totals["food"] += 250           // missing keys read as the zero value (0), so += just works
len(totals)                     // number of keys
delete(totals, "food")          // remove a key (no error if it isn't there)

budget := map[string]int{       // a map literal
    "food": 40000,
    "rent": 120000,
}
```

## "Comma ok": is the key really there?

A missing key reads as `0`, which looks the same as a key whose value *is* `0`.
Ask for a second value to tell them apart:

```go
limit, ok := budget["travel"]
if !ok {
    // "travel" isn't in the map at all
}
```

⚠️ `var m map[string]int` creates a **nil map**. Reading from it is fine, but *writing*
to it panics. Use `map[string]int{}` or `make(map[string]int)`.

## Looping, and the random order

```go
for category, cents := range totals {
    fmt.Println(category, cents)
}
```

Go **shuffles map order on purpose**. When you need a stable order (for example to
return a list), collect the keys into a slice and sort them:

```go
import "slices"

keys := []string{}
for k := range totals {
    keys = append(keys, k)
}
slices.Sort(keys)               // now alphabetical, every time
```

> **Coming from Ruby?** Ruby hashes remember insertion order. Go maps don't.
> `hash.fetch(k)` is closest to `v, ok := m[k]`.

## Your task

`Transaction` is already defined for you in `types.go`.

**Problem 1: `TotalByCategory(txns []Transaction) map[string]int`**

Return a map from each category to the total cents spent in it. With no transactions,
return an empty (non-nil) map.

**Problem 2: `OverBudget(spent, budget map[string]int) []string`**

Return the categories where spending went **over** the budget, sorted alphabetically.

- Spending exactly the budget is fine (not over).
- A category in `spent` that has **no entry** in `budget` counts as over budget if
  anything at all was spent (more than 0). Use "comma ok" to spot the missing entry.
- Categories that are only in `budget` (nothing spent) are never over.
- Return an empty, non-nil slice when nothing is over.

```go
spent  := map[string]int{"food": 45000, "rent": 120000, "games": 2000}
budget := map[string]int{"food": 40000, "rent": 120000}
OverBudget(spent, budget) // ["food", "games"]
```

Run: `go test -v ./concepts/maps/`

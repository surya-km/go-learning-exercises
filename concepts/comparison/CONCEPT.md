# Comparison

## In plain English

Comparison operators ask a question about two values and answer with `true` or `false`:
"are these equal?", "is this one smaller?". They work on numbers and on strings. Both
sides must be the **same type**; Go won't compare an `int` with a `string`.

```go
a, b := 1500, 2000

a == b   // false  equal
a != b   // true   not equal
a < b    // true   less than
a <= b   // true   less than or equal
a > b    // false  greater than
a >= b   // false  greater than or equal
```

## Comparing strings

Strings are compared **character by character, from the left**, using each character's
code. That gives dictionary order for lowercase words, with a few gotchas:

```go
"apple" < "banana"   // true
"apple" < "apples"   // true  (a prefix comes first)
"Zebra" < "apple"    // true! all uppercase letters come before lowercase ones
"10" < "9"           // true! it compares '1' with '9', not ten with nine
"Go" == "go"         // false: comparison is case-sensitive
```

A handy consequence: dates written as `"YYYY-MM-DD"` (like `"2026-03-15"`) sort correctly
as plain strings, because the biggest unit (the year) comes first and every part has a
fixed width.

> **Coming from Ruby?** There's no `<=>` spaceship operator and no `Comparable` mixin.
> You compare with the six operators above. `==` on strings compares their contents
> (like Ruby's `==`), not whether they are the same object.

## Your task

**Problem 1.** Implement `CompareAmounts(aCents, bCents int) int`. Return:

- `-1` if `aCents` is less than `bCents`,
- `0` if they are equal,
- `1` if `aCents` is greater than `bCents`.

(Sorting code uses this "-1 / 0 / 1" style a lot.)

**Problem 2.** Implement `EarlierDate(a, b string) string`. Both inputs are dates in
`"YYYY-MM-DD"` format. Return whichever one is earlier. If they are the same date,
return `a`.

- `EarlierDate("2026-03-15", "2026-01-31")` returns `"2026-01-31"`.
- `EarlierDate("2025-12-31", "2026-01-01")` returns `"2025-12-31"`.

You don't need to split the strings apart: compare them directly.

Run: `go test -v ./concepts/comparison/`

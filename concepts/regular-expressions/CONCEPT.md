# Regular expressions

## In plain English

A regular expression ("regex") is a tiny pattern language for describing text, like
"three digits, a dash, then four digits". Go's `regexp` package lets you check whether
text matches a pattern, find the matching parts, and pull out pieces of it. Use regexes
when the `strings` package isn't enough, and keep them simple.

## Pattern cheat sheet

| Pattern | Means |
|---|---|
| `\d` | one digit `0`-`9` |
| `[A-Z]` | one upper-case letter |
| `[^,]` | any character except a comma |
| `.` | any character |
| `x*` / `x+` / `x?` | zero or more / one or more / optional `x` |
| `x{4}` / `x{2,3}` | exactly 4 / 2 to 3 `x` |
| `^` / `$` | start / end of the text |
| `(...)` | a capture group: remember this part |
| `(?:...)` | a group that is NOT remembered (just for `?`, `*`, etc.) |
| `\.` `\$` | a literal `.` or `$` |

## Using `regexp`

Write patterns in **backtick raw strings** so you don't have to double every backslash.
Compile them once, at package level, with `regexp.MustCompile`, which panics if the
pattern itself is wrong:

```go
import "regexp"

var zipRe = regexp.MustCompile(`^\d{5}$`)

zipRe.MatchString("90210")   // true
zipRe.MatchString("9021")    // false

var amountRe = regexp.MustCompile(`\$(\d+)\.(\d{2})`)

amountRe.FindString("Paid $12.50 today")          // "$12.50"
amountRe.FindStringSubmatch("Paid $12.50 today")  // ["$12.50" "12" "50"]
                                                    //  whole   group1 group2
amountRe.FindAllString("$1.00 and $2.50", -1)      // ["$1.00" "$2.50"]  (-1 = all)
amountRe.ReplaceAllString("$1.00", "X")            // "X"
```

`FindStringSubmatch` returns `nil` if nothing matched. Check that before indexing into it.

> **Coming from Ruby?** Go has no `/.../` literal and no `=~`. Use
> `regexp.MustCompile(`...`)` and methods like `MatchString`. Go's regex engine (RE2) has
> no lookahead or backreferences, but it's guaranteed to run fast on any input.

## Your task

**Problem 1:** implement **`ValidRoutingNumber(s string) bool`**.

A US bank routing number is exactly **9 digits** and nothing else. Return `true` if `s`
matches. Use a regular expression and anchor it with `^` and `$`.

| Call | Result |
|---|---|
| `ValidRoutingNumber("021000021")` | `true` |
| `ValidRoutingNumber("02100002")` | `false` (8 digits) |
| `ValidRoutingNumber("0210000210")` | `false` (10 digits) |
| `ValidRoutingNumber("02100002a")` | `false` |
| `ValidRoutingNumber(" 021000021")` | `false` |

**Problem 2:** implement **`ExtractAmounts(text string) []int`**.

Find every dollar amount in a piece of text (like an email receipt) and return each one
in cents, in order. An amount is a `$` followed by one or more digits, optionally
followed by a `.` and **exactly two** digits.

| Text | Result |
|---|---|
| `"Subtotal $12.50, tax $1.03"` | `[1250 103]` |
| `"You owe $7 now"` | `[700]` |
| `"Price: 12.50"` (no `$`) | `[]` |
| `"$3.5 off"` | `[300]` (`.5` isn't two digits, so only `$3` matches) |

If there are no amounts, return an empty slice or `nil` (the tests accept either).

Hint: use two capture groups, `(\d+)` for dollars and `(\d{2})` for cents, with the
`.` and cents made optional together: `(?:\.(\d{2}))?`. Then call
`FindAllStringSubmatch(text, -1)`. An optional group that didn't match comes back as
`""`. Convert the digits with `strconv.Atoi`.

Run: `go test -v ./concepts/regular-expressions/`

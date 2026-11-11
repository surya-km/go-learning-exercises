# The strings package

## In plain English

The standard library's `strings` package is a toolbox of ready-made functions for text:
trimming spaces, changing case, searching, splitting and joining. Before writing a loop
over a string, check whether `strings` already has a function that does the job. It
usually does.

```go
import "strings"

strings.TrimSpace("  hello \n")          // "hello"
strings.ToUpper("coffee")                // "COFFEE"
strings.ToLower("COFFEE")                // "coffee"

strings.Contains("Starbucks #123", "#")  // true
strings.HasPrefix("POS AMAZON", "POS ")  // true
strings.TrimPrefix("POS AMAZON", "POS ") // "AMAZON"
strings.TrimSuffix("file.csv", ".csv")   // "file"
strings.Index("a,b,c", ",")              // 1   (-1 if not found)

strings.Split("a,b,c", ",")              // []string{"a", "b", "c"}
strings.Fields("  lots   of  space ")    // []string{"lots", "of", "space"}
strings.Join([]string{"a", "b"}, "-")    // "a-b"

strings.Cut("key=value", "=")            // "key", "value", true
strings.Repeat("ab", 3)                  // "ababab"
strings.ReplaceAll("1,234,567", ",", "") // "1234567"
strings.EqualFold("Go", "GO")            // true (case-insensitive compare)
```

`Split` and `Fields` return a **slice** of strings (a list). You can loop over it with
`for _, part := range parts { ... }`.

> **Coming from Ruby?** Most Ruby `String` methods have a twin here, but they're plain
> functions: `s.strip` is `strings.TrimSpace(s)`, `s.split(",")` is `strings.Split(s, ",")`,
> `s.split` (no argument) is `strings.Fields(s)`, and `arr.join("-")` is `strings.Join(arr, "-")`.

## Your task

**Problem 1:** implement **`CleanMerchant(raw string) string`**.

Bank statements show messy merchant names. Clean one up like this:

1. Remove the prefix `"POS "` or `"SQ *"` if the name starts with it (only one, and only at
   the start). Do this after trimming surrounding spaces.
2. Remove everything from the first `" #"` onwards (store numbers like `" #1234"`).
3. Collapse runs of spaces into a single space and trim the ends.
4. Upper-case the result.

| Call | Result |
|---|---|
| `CleanMerchant("POS starbucks #1234")` | `"STARBUCKS"` |
| `CleanMerchant("  SQ *blue   bottle coffee ")` | `"BLUE BOTTLE COFFEE"` |
| `CleanMerchant("Whole Foods Market")` | `"WHOLE FOODS MARKET"` |
| `CleanMerchant("Target #0042 Austin TX")` | `"TARGET"` |
| `CleanMerchant("   ")` | `""` |

**Problem 2:** implement **`ParseTags(s string) []string`**.

Users type tags for a transaction separated by commas, like `"Food, travel ,, FOOD"`.
Return the tags:

- Each tag trimmed of spaces and lower-cased.
- Empty tags skipped.
- Duplicates removed (keep the first one), in the order they first appear.
- If there are no tags, return an empty slice or `nil` (the tests accept either).

| Call | Result |
|---|---|
| `ParseTags("food, travel")` | `["food" "travel"]` |
| `ParseTags("Food, travel ,, FOOD")` | `["food" "travel"]` |
| `ParseTags(" , ,")` | `[]` |

Hint: you can check "have I seen this tag already?" with a loop over the tags you've kept
so far, or with `slices.Contains(kept, tag)` from the `slices` package.

Run: `go test -v ./concepts/strings-package/`

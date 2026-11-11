# Runes

## In plain English

A Go string is a row of **bytes**. Plain English letters take 1 byte each, but characters
like `é`, `€` or `😀` take 2 to 4 bytes. A **rune** is one whole character, however many
bytes it needs. If you want to count, cut or change *characters*, work with runes and not
bytes, or you'll chop an emoji in half.

```go
s := "café"
len(s)                       // 5  -> bytes! "é" takes 2 bytes
utf8.RuneCountInString(s)    // 4  -> characters (needs import "unicode/utf8")
s[3]                         // 195 -> a single byte, half of "é"
```

## Looping over runes

`for ... range` over a string gives you one **rune** at a time, plus the byte index where
it starts:

```go
for i, r := range "€5" {
    fmt.Println(i, string(r))
}
// 0 €
// 3 5     <- "€" took bytes 0, 1 and 2
```

## Converting to and from `[]rune`

```go
runes := []rune("café")   // [99 97 102 233] -> one element per character
len(runes)                 // 4
string(runes[:3])          // "caf"
string('é')                // "é"  (a rune converts to a one-character string)
```

A rune literal uses single quotes: `'a'`, `'é'`, `'😀'`. Its type is `rune` (another
name for `int32`).

## The `unicode` package

```go
unicode.ToUpper('é')     // 'É'
unicode.IsLetter('7')    // false
unicode.IsSpace('\t')    // true
```

> **Coming from Ruby?** Ruby's `"café".length` is 4 because Ruby counts characters. Go's
> `len` counts bytes. `[]rune(s)` is roughly Ruby's `s.chars`.

## Your task

**Problem 1:** implement **`Truncate(s string, max int) string`**.

Bank apps cut long merchant names to fit on screen. Return `s` unchanged if it has at
most `max` **characters**. Otherwise, keep the first `max-1` characters and add `"…"`
(a single ellipsis character), so the result is exactly `max` characters long.

- If `max` is 0 or less, return `""`.
- Never split a multi-byte character.

| Call | Result |
|---|---|
| `Truncate("Coffee", 10)` | `"Coffee"` |
| `Truncate("Starbucks Coffee", 10)` | `"Starbucks…"` |
| `Truncate("Café Crème", 6)` | `"Café …"` |
| `Truncate("😀😀😀", 2)` | `"😀…"` |
| `Truncate("abc", 1)` | `"…"` |
| `Truncate("abc", 0)` | `""` |

**Problem 2:** implement **`Initials(name string) string`**.

Return the first character of each word, in upper case. Words are separated by one or
more whitespace characters (use `unicode.IsSpace`). Names can include accented letters.

| Call | Result |
|---|---|
| `Initials("ada lovelace")` | `"AL"` |
| `Initials("Émile Zola")` | `"ÉZ"` |
| `Initials("  élodie   dupont ")` | `"ÉD"` |
| `Initials("ørjan")` | `"Ø"` |
| `Initials("")` | `""` |

Run: `go test -v ./concepts/22-runes/`

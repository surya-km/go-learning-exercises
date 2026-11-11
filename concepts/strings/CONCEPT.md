# Strings

## In plain English

A string is a piece of text, written in double quotes: `"hello"`. Once a string is made,
it can't be changed. Every "change" actually builds a new string. You can join strings
with `+`, get their length with `len`, and cut out pieces with `s[start:end]`.

```go
first := "Ada"
last := "Lovelace"
full := first + " " + last   // "Ada Lovelace"

len(full)        // 12
full[0:3]        // "Ada"       (from index 0 up to, but NOT including, 3)
full[4:]         // "Lovelace"  (from index 4 to the end)
full[len(full)-4:] // "lace"    (the last 4 characters)
full == "Ada Lovelace"  // true
```

Building a string a bit at a time in a loop:

```go
stars := ""
for i := 0; i < 3; i++ {
    stars += "*"
}
// stars is "***"
```

Raw strings use backticks. Nothing inside them is special, not even `\n`:

```go
path := `C:\Users\ada`   // the backslashes are kept as-is
```

**Note:** `len` and `s[i:j]` count **bytes**. For plain English letters and digits, one
byte is one character, which is all this exercise uses. The **runes** concept covers
characters like `é` and emoji.

> **Coming from Ruby?** Ruby strings can be changed in place (`<<`). Go strings can't.
> `s[-4..]` in Ruby is `s[len(s)-4:]` in Go, because Go has no negative indexes.

## Your task

**Problem 1:** implement **`MaskAccount(number string) string`**.

Hide an account number except for its last 4 characters. Replace every other character
with `*`.

- If the number has 4 or fewer characters, return it unchanged.

| Call | Result |
|---|---|
| `MaskAccount("1234567890")` | `"******7890"` |
| `MaskAccount("12345")` | `"*2345"` |
| `MaskAccount("1234")` | `"1234"` |
| `MaskAccount("")` | `""` |

**Problem 2:** implement **`ShortID(id string) string`**.

Transaction IDs are long. Show the first 4 characters, then `"..."`, then the last 4.

- If the ID has 11 or fewer characters, return it unchanged. Shortening it wouldn't save
  any space.

| Call | Result |
|---|---|
| `ShortID("txn_9f8e7d6c5b4a")` | `"txn_...5b4a"` |
| `ShortID("abcdefghijkl")` | `"abcd...ijkl"` |
| `ShortID("abcdefghijk")` | `"abcdefghijk"` |
| `ShortID("ab")` | `"ab"` |

Run: `go test -v ./concepts/strings/`

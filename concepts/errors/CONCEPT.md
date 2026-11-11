# Errors

## In plain English

Go has **no exceptions**. A function that can fail returns its result *plus* an `error`.
The caller checks it right away. Failures are ordinary values that you handle in
plain sight, instead of jumping somewhere else in the program.

```go
n, err := strconv.Atoi("42")
if err != nil {
    return 0, err   // pass the problem up to whoever called us
}
// safe to use n here
```

You will write `if err != nil` a *lot*. That's normal, idiomatic Go.

> **Coming from Ruby?** Instead of `raise` / `rescue`, the error comes back as a
> second return value, and you have to deal with it on the next line.

## Creating errors

```go
import "errors"

errors.New("amount is empty")               // a simple error
fmt.Errorf("bad amount %q", s)              // a formatted error
```

## Sentinel errors: errors callers can check for

When callers need to react to a *specific* failure, declare it once as a package-level
variable. These are called "sentinel" errors:

```go
var ErrEmpty = errors.New("amount is empty")
```

## Wrapping: adding context without losing the original

`%w` in `fmt.Errorf` *wraps* an error. The message gets extra detail, but callers can
still find the original error inside:

```go
return 0, fmt.Errorf("parsing %q: %w", s, ErrInvalid)
// message: parsing "abc": invalid amount
```

## Checking: `errors.Is`

```go
if errors.Is(err, ErrInvalid) {   // true even if ErrInvalid is wrapped inside err
    // show a "please enter a valid amount" message
}
```

Don't compare messages (`err.Error() == "..."`). Messages change, but sentinels don't.

## Useful helpers for this exercise

```go
strings.TrimSpace("  12.50 ")     // "12.50"
strings.Cut("12.50", ".")         // "12", "50", true
strconv.Atoi("12")                // 12, nil   (string → int)
len("50")                         // 2
```

## Your task

Implement **`ParseAmount(s string) (int, error)`** in `exercise.go`. It turns a dollar
string into integer cents:

| Input | Result |
|---|---|
| `"12.34"` | `1234, nil` |
| `"7"` | `700, nil` |
| `"0.5"` | `50, nil` (one decimal digit means tenths) |
| `"  3.00 "` | `300, nil` (surrounding spaces are ignored) |
| `""` or `"   "` | `0`, an error where `errors.Is(err, ErrEmpty)` |
| `"abc"`, `"1.234"`, `"1."`, `"-5"`, `"1.2.3"` | `0`, an error where `errors.Is(err, ErrInvalid)` |

Rules:
- At most **two** digits after the decimal point. If there's a `.`, at least one digit must follow it.
- No negative amounts.
- Invalid-input errors should **wrap** `ErrInvalid` and include the original input in the message.

`ErrEmpty` and `ErrInvalid` are already declared in `errors.go`.

Run: `go test -v ./concepts/errors/`

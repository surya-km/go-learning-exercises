# Type assertion

## In plain English

Sometimes you hold a value whose exact type you don't know yet: it's stored as an
interface like `any` (which means "could be anything"). A **type assertion** asks
"is this actually a string?" and, if so, gives you the string back. A **type switch**
asks that question for several types in one go, like sorting mail into the right pile.

## Comma-ok assertion: "is it this type?"

```go
var v any = "hello"

s, ok := v.(string)   // s = "hello", ok = true
n, ok := v.(int)      // n = 0 (zero value), ok = false
```

Always use the two-value `x, ok :=` form unless you are certain. The one-value form
`v.(int)` **crashes the program** (panics) if `v` isn't an `int`.

## Type switch: "which type is it?"

```go
func Kind(v any) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("int %d", x)      // x is an int here
    case string:
        return "string " + x                 // x is a string here
    case nil:
        return "nothing"
    default:
        return fmt.Sprintf("other %T", x)    // %T prints the type name
    }
}
```

Inside each `case`, `x` already has that case's type, so no extra conversion is needed.

> **Coming from Ruby?** A type switch is like `case value when Integer ... when String ... end`,
> and the comma-ok form is like `value.is_a?(String)` combined with getting the value back.

## Your task

A bank sends us events of different kinds. They arrive as `any`. The event types
`Deposit`, `Withdrawal` and `Fee` are declared in `types.go`.

**Problem 1: `Describe(event any) string`**

Use a **type switch** to return:

| Event | Result |
|---|---|
| `Deposit{From: "Acme", AmountCents: 250000}` | `"deposit of 250000 cents from Acme"` |
| `Withdrawal{ATM: true, AmountCents: 2000}` | `"ATM withdrawal of 2000 cents"` |
| `Withdrawal{ATM: false, AmountCents: 2000}` | `"withdrawal of 2000 cents"` |
| `Fee{Reason: "overdraft", AmountCents: 3500}` | `"fee of 3500 cents (overdraft)"` |
| a `string` such as `"hello"` | `"note: hello"` |
| `nil` | `"empty event"` |
| anything else (e.g. `42`) | `"unknown event of type int"` (use `%T`) |

**Problem 2: `LookupString` and `LookupCents` with comma-ok assertions**

Events also come with a payload: a `map[string]any` of extra fields.

- `LookupString(payload map[string]any, key string) (string, bool)` returns the value at
  `key` and `true` if it exists **and** is a `string`. Otherwise `""` and `false`.
- `LookupCents(payload map[string]any, key string) (int64, bool)` returns the value at
  `key` and `true` if it exists **and** is an `int64`. Otherwise `0` and `false`.
  (An `int` or a `string` like `"500"` does **not** count; only `int64`.)

A `nil` map should work too (reading from a nil map is safe in Go).

Run: `go test -v ./concepts/type-assertion/`

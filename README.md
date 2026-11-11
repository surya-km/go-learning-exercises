# Go Learning Exercises

A hands-on repo for learning Go as a team. It has **39 concepts**, named after the
[Exercism Go track's concepts](https://github.com/exercism/go/tree/main/concepts). Each
concept folder has 1–2 small coding problems.

Every concept folder contains:

| File | What it is |
|---|---|
| `CONCEPT.md` | A short plain-English explanation of the idea, with examples. Read this first. |
| `exercise.go` | The code **you** edit. Functions are stubbed out with `TODO`s. |
| `exercise_test.go` | Tests that pass once your solution is correct. Don't edit these. |
| `types.go` (some folders) | Types the exercise uses. Read it, but don't edit it. |
| `.meta/solution.go` | A reference answer. Try not to peek until your tests pass! |

**The tests are the teacher.** Each exercise starts with failing tests, and you're done
when they all pass.

## Setup

1. Install Go 1.25+ with `brew install go`, and check it with `go version`.
2. Clone this repo and `cd` into it.
3. (Recommended) Install the Go extension for VS Code, or use GoLand.

## How to do an exercise

```sh
# 1. Read the concept
cat concepts/01-basics/CONCEPT.md

# 2. Run the tests. They fail at first, which is expected.
go test ./concepts/01-basics/

# 3. Edit concepts/01-basics/exercise.go until the tests pass
go test -v ./concepts/01-basics/
```

Tip: `go test -v -run TestSplitBill ./concepts/01-basics/` runs a single test.

> Running `go test ./...` from the root shows failures for every exercise you haven't
> solved yet. That's normal, so work on one folder at a time.

## Learning path

The folders are numbered in learning order (`01-basics` → `39-randomness`), so work
through them top to bottom. Later concepts build on earlier ones.

### 1. Foundations
| # | Concept | What you build |
|---|---|---|
| 1 | [basics](concepts/01-basics) | `Greeting`, `SplitBill` |
| 2 | [variables](concepts/02-variables) | `MonthEnd`, `Transfer` |
| 3 | [comments](concepts/03-comments) | Doc comments, `FormatAmount`, `Describe` |
| 4 | [constants](concepts/04-constants) | A `Frequency` enum with `iota`, `YearlyCostCents` |
| 5 | [zero-values](concepts/05-zero-values) | `ApplyDefaults`, `AddSpend` |

### 2. Numbers and logic
| # | Concept | What you build |
|---|---|---|
| 6 | [numbers](concepts/06-numbers) | `SplitBill` (fair split), `MakeChange` |
| 7 | [arithmetic-operators](concepts/07-arithmetic-operators) | `TipCents`, `RoundUp` |
| 8 | [floating-point-numbers](concepts/08-floating-point-numbers) | `ApplyRate`, `AlmostEqual` |
| 9 | [type-conversion](concepts/09-type-conversion) | `AverageCents`, `OrderSummary` |
| 10 | [booleans](concepts/10-booleans) | `CanWithdraw`, `NeedsReview` |
| 11 | [comparison](concepts/11-comparison) | `CompareAmounts`, `EarlierDate` |
| 12 | [conditionals-if](concepts/12-conditionals-if) | `OverdraftFee`, `BudgetStatus` |
| 13 | [conditionals-switch](concepts/13-conditionals-switch) | `Category`, `CreditRating` |

### 3. Functions and loops
| # | Concept | What you build |
|---|---|---|
| 14 | [functions](concepts/14-functions) | `Greeting`, `TipCents`, `TotalWithTip` |
| 15 | [multiple-return-values](concepts/15-multiple-return-values) | `SplitBill`, `CoinChange` |
| 16 | [for-loops](concepts/16-for-loops) | `MonthsToPayOff`, `CountWeekdays` |
| 17 | [variadic-functions](concepts/17-variadic-functions) | `MaxAmount`, `AddToBill` |

### 4. Collections
| # | Concept | What you build |
|---|---|---|
| 18 | [slices](concepts/18-slices) | `Sum`, `Above`, `Pages` |
| 19 | [maps](concepts/19-maps) | `TotalByCategory`, `OverBudget` |
| 20 | [range-iteration](concepts/20-range-iteration) | `RunningBalance`, `InstallmentPlan`, `MaskCard`, `BiggestCategory` |

### 5. Text
| # | Concept | What you build |
|---|---|---|
| 21 | [strings](concepts/21-strings) | `MaskAccount`, `ShortID` |
| 22 | [runes](concepts/22-runes) | `Truncate`, `Initials` (handles é, € and emoji) |
| 23 | [strings-package](concepts/23-strings-package) | `CleanMerchant`, `ParseTags` |
| 24 | [strings-builder](concepts/24-strings-builder) | `Receipt`, `GroupDigits` |
| 25 | [string-formatting](concepts/25-string-formatting) | `StatementLine`, `SpendingShare` |
| 26 | [regular-expressions](concepts/26-regular-expressions) | `ValidRoutingNumber`, `ExtractAmounts` |

### 6. Your own types
| # | Concept | What you build |
|---|---|---|
| 27 | [structs](concepts/27-structs) | `NewAccount`, `Summarize` |
| 28 | [type-definitions](concepts/28-type-definitions) | `Cents.Format`, `Cents.Split`, `YearlyCost` |
| 29 | [methods](concepts/29-methods) | `Bill.IsDueBy`, `Bill.Label`, `Bills.TotalDueBy` |
| 30 | [pointers](concepts/30-pointers) | `Account.Deposit`, `Account.Withdraw`, `Transfer` |
| 31 | [stringers](concepts/31-stringers) | `String()` for `Money`, `Transaction`, `Statement` |
| 32 | [interfaces](concepts/32-interfaces) | A `Payment` interface on 3 types, `MonthlyTotal`, `Biggest` |
| 33 | [type-assertion](concepts/33-type-assertion) | `Describe` (type switch), `LookupString`, `LookupCents` |

### 7. Errors
| # | Concept | What you build |
|---|---|---|
| 34 | [errors](concepts/34-errors) | `ParseAmount` with sentinel errors |
| 35 | [error-wrapping](concepts/35-error-wrapping) | `ProcessPayment`, `DeclineCode` (`errors.As`), `RootCause`, `ProcessBatch` (`errors.Join`) |

### 8. More tools
| # | Concept | What you build |
|---|---|---|
| 36 | [first-class-functions](concepts/36-first-class-functions) | `Filter`, `SpendingTracker` (closure) |
| 37 | [packages](concepts/37-packages) | A `money` sub-package, `ReceiptLine`, `Total` |
| 38 | [time](concepts/38-time) | `DaysInMonth`, `NextDueDate` (month-end safe), `DaysBetween`, `FormatDue` |
| 39 | [randomness](concepts/39-randomness) | `VerificationCode`, `ShuffledCopy` |

### Planned next
- `concurrency`: goroutines, channels, `sync.WaitGroup`, `context`
- `testing`: writing your own table-driven tests
- `real-world`: HTTP handlers, JSON, and the patterns our services actually use

Want to add one? See [CONTRIBUTING.md](CONTRIBUTING.md).

## Checking the reference solutions

Maintainers can confirm that every reference solution passes its tests:

```sh
./scripts/verify-solutions.sh
```

CI runs this on every push.

## More practice

- [exercism.org Go track](https://exercism.org/tracks/go): free, ordered, with optional mentoring
- [A Tour of Go](https://go.dev/tour/): the official interactive intro
- [Go by Example](https://gobyexample.com/): short annotated snippets
- [Effective Go](https://go.dev/doc/effective_go): idiomatic style

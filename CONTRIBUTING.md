# Contributing an exercise

Anyone on the team can add an exercise. Keep each one **small** (one concept, 15–45 minutes of work).

## Steps

1. Copy the template into `concepts/`, numbered where it fits in the learning path:
   ```sh
   cp -R _template concepts/40-my-concept
   ```
   To insert a concept in the middle, renumber the folders after it and update the
   README links. Package names stay the same, so no Go code changes are needed (except
   import paths in `37-packages`, if it moves).
2. Rename the package in all `.go` files (`package template` → e.g. `package myexercise`).
   Package names are lowercase, with no dashes or underscores.
3. Write `CONCEPT.md`: a plain-English explanation, short code examples, then a "Your task" section.
4. Write `exercise_test.go` first. Prefer **table-driven tests** with clear failure messages
   that show the input, what came back, and what was expected.
5. Write `.meta/solution.go`: a clean, idiomatic reference answer.
6. Write `exercise.go`: the same function signatures as the solution, with bodies
   replaced by `// TODO: implement` and a zero-value return.
7. Check it:
   ```sh
   go vet ./...                      # stubs must compile
   ./scripts/verify-solutions.sh     # solution must pass
   go test ./path/to/exercise/       # stub should FAIL (otherwise the tests aren't testing anything)
   ```
8. Add a row to the right section of the **Learning path** in `README.md`.

## Guidelines

- **Use realistic examples.** Money, transactions and HTTP requests beat foo/bar.
- **Money is integer cents**, never `float64`.
- **Standard library only** unless the exercise is specifically about a library.
- **Copying from Exercism?** The Exercism Go track is MIT licensed. Keep its license
  notice and add the exercise to `THIRD_PARTY_NOTICES.md`.

## Why `.meta` and `_template`?

`go test ./...` and `go vet ./...` skip directories whose names start with `.` or `_`.
That keeps solutions and the template out of normal runs, so they don't clash with the
stub in the same package.

# Demo: what "compiled" means

The exercises in `concepts/` are packages you check with `go test`. This folder is
different: it's a complete **program** you can build and run on its own.

## In plain English

- **Ruby is interpreted.** `ruby app.rb` reads and runs your code line by line every
  time, so Ruby has to be installed wherever it runs. It's like reading a recipe
  while you cook.
- **Go is compiled.** `go build` translates the whole program into machine code
  **once**, and saves it as a single ready-to-run file. That file runs without Go
  installed. It's like cooking the meal in advance and handing over a packed box.

## Try it

From the repo root:

```sh
cd demo/hello

go run .            # build and run in one step (handy while coding)

go build -o hello   # compile once: creates a file called "hello"
ls -lh              # main.go = your code (the recipe), hello = the program (the packed box)
./hello             # run it directly, no "go" command needed
```

You'll see:

```
Hello from a compiled program!
```

## Things to notice

- **One file.** `hello` is the whole program: no Go, gems or dependencies needed to run it.
  You could copy it to another Mac and run it there.
- **Build for another system** from your Mac, for example a Linux server:
  ```sh
  GOOS=linux GOARCH=amd64 go build -o hello-linux
  ```
- **Mistakes are caught before running.** Change `fmt.Println` to `fmt.Printn` and run
  `go build`. Go refuses to build and tells you exactly what's wrong. Ruby would only
  find a typo like that when the line actually runs.
- **Clean up** with `rm hello hello-linux`. Built files are ignored by Git (see
  `.gitignore`), so you won't commit them by accident.

## What makes a program a program

Two things turn Go code into a runnable program:

1. `package main` at the top
2. a `func main()`, which is where the program starts

The exercises don't have these, which is why you use `go test` on them instead of
`go run`.

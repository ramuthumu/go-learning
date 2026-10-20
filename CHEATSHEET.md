# Go quick reference

## Everyday commands

```sh
go version             # installed toolchain
go mod init example/app # start a module
go run .                # build and run the current package
go test ./...           # run package tests recursively
go test -run TestName   # run matching test(s)
go fmt ./...            # format Go source
go doc strings          # read package docs
```

## Language shapes

```go
const limit = 10
count := 3

if count > limit {
	// ...
}

for i := 0; i < count; i++ {
	// Go's loop form
}

items := []string{"a", "b"} // slice
scores := map[string]int{}   // map

type User struct {
	Name string
}

func (u User) Label() string { // method with value receiver
	return u.Name
}

func find(name string) (User, error) { // multiple results
	return User{}, nil
}
```

## Habits that make Go easier

- Read the whole signature first: types and return values often explain a lot.
- Handle errors where you can add context or make a real decision.
- Keep interfaces small and define them near the code that consumes them.
- Prefer clear, direct control flow; avoid hiding work behind clever helpers.
- Run `gofmt` early. Go's formatter is part of the normal workflow.
- Start with the standard library. Add a dependency when it solves a real need.
- A slice is a descriptor over an underlying array; appending can allocate a
  new array, so do not assume two slices remain linked after append.
- A map is reference-like and must be initialized before writing to it.
- Use `context.Context` to carry cancellation and deadlines through operations;
  pass it explicitly, usually as the first parameter.

## Tests at a glance

Put tests in a file ending `_test.go`. A basic test receives `*testing.T`:

```go
func TestDouble(t *testing.T) {
	if got := Double(4); got != 8 {
		t.Fatalf("Double(4) = %d, want 8", got)
	}
}
```

For many inputs, use a table of cases and loop over it. Keep each failure
message specific enough to show the input, actual result, and expected result.

## A professional development loop

```sh
go fmt ./...                    # canonical formatting
go test ./...                   # behavior checks
go test -race ./...             # exercised shared-memory races
go vet ./...                    # common suspicious constructs
go test ./examples/loglens -run=^$ -bench=. -benchmem
go doc net/http.Server          # inspect the exact API contract
go mod tidy                     # reconcile module dependencies
```

| Concern | Starting point | Rule to remember |
|---|---|---|
| CLI input/output | `io.Reader`, `io.Writer`, `flag` | Results to stdout; diagnostics to stderr |
| File streaming | `bufio.Scanner` | Bound token size and check `Err()` |
| HTML | `html/template`, `embed` | Escape data; do not mark user input trusted |
| HTTP | `net/http`, `httptest` | Set headers before writing the response |
| JSON | `encoding/json` | Separate input types; validate after decoding |
| SQL | `database/sql` | Parameterize values; close rows; check `rows.Err()` |
| Errors | `errors.Is`, `errors.As`, `%w` | Preserve causes; handle at a useful boundary |
| Cancellation | `context.Context` | Every blocking operation needs an exit path |
| Shared state | `sync.Mutex` | Protect the full invariant, not just one field |
| Concurrent work | channels, `sync.WaitGroup` | Bound workers and define who closes channels |
| Diagnostics | `log/slog`, benchmarks, `pprof` | Measure before optimizing |

## Context lifetime

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, query, args...)
```

Request work normally uses the request's context. Durable background work uses
an application lifecycle context and a deadline per attempt. Shutdown draining
uses a fresh context, not the already-cancelled signal context.

## Transaction shape

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil { return err }
defer tx.Rollback()
// Perform every operation through tx, checking each error.
return tx.Commit()
```

A transaction protects related database operations. It does not automatically
make remote HTTP effects atomic or prevent duplicate external deliveries.

## Delivery review questions

Can the operation be cancelled? What happens after a partial failure? Which
values are shared? Who owns cleanup? Is data scoped to the verified user or
tenant? Can a request be replayed safely? How will an operator diagnose a failure?
Use the [full guide](GUIDE.md) for explanations and worked examples.

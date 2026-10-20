# Go: from first program to professional practice

36 guides • 6 stages • practical projects

Use this guide with the runnable examples in the repository. All shell commands assume the repository root unless a lesson explicitly creates a separate module.


## 1. Your first program and development loop

You need a small program that can be built and run on another computer. Start with the tools, then learn what each line does.

### A program is a set of instructions

An editor saves source code as text. The Go compiler checks that code and produces a program your operating system can execute. A terminal is where you ask the tools to do things. You do not need another language first; use this first stage slowly if programming is new.

### Understand the smallest Go program

package main identifies an executable package. import brings another package into this file; fmt provides formatted input and output. func main() declares the starting function. Statements inside its braces run in order. := creates a local variable and assigns its first value. A quoted string is text; \n is a newline.

### A module gives code an identity

go.mod names a module and records its minimum Go version and dependencies. A package groups Go files in a directory. go run . builds and runs the current package; go build writes an executable. gofmt standardizes source layout. Compiler errors such as “declared and not used” are feedback: read the filename and line first.

### Build it yourself

1. Install a currently supported Go release from the official installation page. Run go version in a terminal. Use Go 1.25 or newer for these examples.
2. Make a scratch folder; run go mod init example.com/first-go. Save the example below as main.go, then run go run .
3. Change the name, add a second print, and deliberately misspell a variable. Read the compiler message and repair it.
4. Run go fmt ./... and go build -o hello .; run ./hello (hello.exe on Windows).

```sh
go version
go run ./examples/01-hello
```


<details><summary>Reference / excerpt</summary>

```
package main

import "fmt"

func main() {
	name := "Go learner"
	fmt.Printf("Hello, %s!\n", name)
}
```

</details>

### Checkpoints

- You can run your own program from an empty folder.
- You can explain package, import, func, variable, and module in your own words.

**Watch for:** Run commands in the folder containing go.mod. Do not paste terminal commands inside a .go file.

**Recall:** What is the difference between source code and an executable?

Source is text you edit. An executable is the compiler’s output that your OS runs. go run performs compilation for you before running.


Further reading: [Install Go](https://go.dev/doc/install) · [First Go program](https://go.dev/doc/tutorial/getting-started)


## 2. Values, decisions, loops, and functions

Build an order calculator that totals quantities and labels larger orders.

### Values have types

int stores whole numbers, float64 stores approximate real numbers, bool stores true or false, and string stores text. Go checks types at compile time. Zero values are useful: 0, false, and an empty string. Use var when you need a declaration without an initial value, := inside a function, and const for a fixed value.

### Control the path through a program

if runs a block when a condition is true. for repeats work; it can be a counted loop or a condition-only loop. switch selects among named cases. Braces make the scope visible: a variable declared inside a block is normally only available there.

### Functions make a decision reusable

A function declares input types and output types. Return the calculated value so callers can print it, test it, or combine it with other values. Keep money in integer cents for this exercise; floating-point arithmetic is approximate. Production monetary rules also need explicit rounding, currency, and overflow decisions.

### Build it yourself

1. Predict the output before running the example.
2. Add a shippingCents(total int) int function: zero shipping at 5000 cents, otherwise 499.
3. Handle zero or negative quantity explicitly. Print the subtotal, shipping, and final total.
4. Add three orders in a loop and accumulate their total.

```sh
go run ./examples/02-values
```


<details><summary>Reference / excerpt</summary>

```
package main

import "fmt"

func total(unitCents, quantity int) int { return unitCents * quantity }

func main() {
	cents := total(1250, 3)
	label := "small order"
	if cents >= 3000 {
		label = "large order"
	}
	fmt.Printf("%s: $%d.%02d\n", label, cents/100, cents%100)
	for i := 1; i <= 3; i++ {
		fmt.Println("package", i)
	}
}
```

</details>

### Checkpoints

- 1250 cents × 3 prints $37.50.
- An order exactly at the free-shipping threshold behaves correctly.
- You can distinguish assignment (=) from equality (==).

**Watch for:** 1250/100 uses integer division. Use quotient and remainder to display cents, rather than introducing float rounding accidentally.

**Recall:** Why should the calculation return a value rather than print it?

Returning separates the business calculation from its presentation, so the same function works in a CLI, test, or web handler.


Further reading: [Go basics](https://go.dev/tour/basics/1)


## 3. Slices, maps, strings, and ownership

Count words in a support log and produce output that can be compared across runs.

### Choose by the question you need to answer

A slice is an ordered sequence with a length and capacity. append may grow it, so keep the returned slice: items = append(items, value). A map looks up values by a key. Reading a missing key returns the value type’s zero value; value, ok := m[key] distinguishes missing from present.

### Understand shared storage

A slice describes a region of a backing array. Copying the slice does not copy all elements. A subslice can change the original array. Make an explicit copy when a function needs its own data. Maps also refer to shared data: assigning a map variable does not clone its entries.

### Text and stable output

Go strings contain bytes, often UTF-8. len(text) counts bytes; range over a string yields rune values. A rune is a Unicode code point, which is not always a whole displayed character. Map iteration order is unspecified; sort keys when generating reports, snapshots, or predictable CLI output.

### Build it yourself

1. Run the word-count example and explain why the repeated word has count 2.
2. Add a word that contains non-ASCII text; compare byte length and rune count.
3. Modify a subslice and observe the original. Then repeat using a separate copy.
4. Build a top-three report: highest count first, alphabetical order for ties.

```sh
go run ./examples/03-collections
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, word := range strings.Fields("Go ships Go tools") {
		counts[strings.ToLower(word)]++
	}
	keys := make([]string, 0, len(counts))
	for word := range counts {
		keys = append(keys, word)
	}
	sort.Strings(keys)
	for _, word := range keys {
		fmt.Println(word, counts[word])
	}
	original := []int{1, 2, 3}
	view := original[:2]
	view[0] = 99
	fmt.Println("shared array:", original)
	copyOf := append([]int(nil), original...)
	copyOf[0] = 1
	fmt.Println("independent copy:", copyOf)
}
```

</details>

### Checkpoints

- Empty input produces an empty report without panicking.
- Repeated runs order the output identically.
- You can explain which values share a backing array.

**Watch for:** A nil map can be read but cannot be assigned to. Initialize writable maps with make or a literal.

**Recall:** Why can changing view[0] change original[0]?

Both slices point into the same backing array. A new slice header is not a deep copy of the elements.


Further reading: [Slices explained](https://go.dev/blog/slices-intro) · [Strings and runes](https://go.dev/blog/strings)


## 4. Structs, pointers, and methods

Model a support ticket whose state changes when it is completed.

### Give related fields a name

A struct groups fields into one value. Ticket{Title: "Fix export"} sets a field and leaves the rest at zero values. Assigning a struct copies it, but reference-like fields such as slices still point to shared storage. Capitalized names are exported to other packages.

### A pointer refers to a value

&ticket gets its address. A *Ticket parameter receives a pointer, allowing a function to modify that original ticket. You usually do not need explicit dereferencing for field access: Go allows t.Title. Use pointers for intentional mutation or when copying a large value would be inappropriate.

### Methods express behavior

A method is a function with a receiver. A pointer receiver can update the original; a value receiver gets a copy. Keep receiver choices consistent and make invalid state transitions explicit. Go uses composition rather than class inheritance: types can contain other types without an object hierarchy.

### Build it yourself

1. Run the example: close the original ticket, then inspect its earlier copy.
2. Add an assigned-to field and a method that rejects an empty assignee.
3. Return an error if a completed ticket is completed again.
4. Explain whether a slice of Ticket and a slice of *Ticket have the same mutation behavior.

```sh
go run ./examples/04-model
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"errors"
	"fmt"
)

type Ticket struct {
	Title  string
	Closed bool
}

func (t *Ticket) Close() error {
	if t.Closed {
		return errors.New("ticket already closed")
	}
	t.Closed = true
	return nil
}

func main() {
	ticket := Ticket{Title: "Fix the export"}
	copyOf := ticket
	if err := ticket.Close(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ticket.Closed, copyOf.Closed)
	fmt.Println(ticket.Close())
}
```

</details>

### Checkpoints

- Completing one copied struct does not set the other copy’s Closed field.
- Repeated completion is handled deliberately.
- You can explain when a pointer can be nil.

**Watch for:** A struct containing a sync.Mutex should not be copied after use. This matters when you add concurrency later.

**Recall:** When does a value receiver fail to update the caller’s struct?

The receiver is a copy. Changing a scalar field on it changes that copy; use a pointer receiver when mutation is part of the API.


Further reading: [Methods and interfaces tour](https://go.dev/tour/methods/1)


## 5. Errors, cleanup, and useful failure messages

A command receives invalid quantities and missing files. Give callers useful information without crashing.

### Errors are ordinary returned values

Go commonly returns (value, error). A nil error means success. Check the error before using the value. Return errors when a caller can choose what to do; a CLI may print a message and exit, while an HTTP handler may return a status.

### Preserve causes while adding context

fmt.Errorf with %w wraps an error, keeping its cause available to errors.Is and errors.As. Use Is for a known error value and As for an error type. Do not compare error text. Avoid logging at every layer; add context as errors travel and log once where the action is decided.

### Release resources on every path

After a successful os.Open, defer file.Close() schedules cleanup when the function returns. Defers run in reverse order. A panic is for a broken assumption or unrecoverable programming error, not ordinary bad input. os.Exit skips deferred calls, so put the work in a function that returns before main exits.

### Build it yourself

1. Run the three inputs in the example and trace the different error paths.
2. Wrap a file-open error with the filename and use errors.Is(err, os.ErrNotExist).
3. Write a function that returns early after opening a file; confirm cleanup still occurs.
4. Add a maximum quantity and a clearly named validation error.

```sh
go run ./examples/05-errors
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidQuantity = errors.New("quantity must be positive")

func parseQuantity(raw string) (int, error) {
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parse quantity %q: %w", raw, err)
	}
	if n < 1 {
		return 0, ErrInvalidQuantity
	}
	return n, nil
}

func main() {
	for _, raw := range []string{"4", "0", "oops"} {
		n, err := parseQuantity(raw)
		if errors.Is(err, ErrInvalidQuantity) {
			fmt.Println("please order at least one item")
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("accepted", n)
	}
}
```

</details>

### Checkpoints

- Invalid input returns an error; it does not panic.
- The caller can recognize a missing file after wrapping.
- You can explain the cleanup consequences of os.Exit.

**Watch for:** Defer inside a long loop runs at the end of the surrounding function. Put per-file work in a helper so files close each iteration.

**Recall:** Why use %w instead of %v when returning another error?

%w preserves a wrapped cause for programmatic inspection. %v only includes its text in the new message.


Further reading: [Working with errors](https://go.dev/blog/go1.13-errors) · [Defer, panic, recover](https://go.dev/blog/defer-panic-and-recover)


## 6. Packages, interfaces, and generics

The same report should write to a terminal, a file, or a buffer used by a test.

### Make package boundaries meaningful

A module may contain several packages. A package should represent a coherent responsibility, with a small exported API. internal directories restrict import access. Keep main responsible for wiring dependencies together. Start with a few packages and split when responsibilities become clear, not one package per type.

### Describe the behavior a caller needs

An interface is a method set. A type satisfies it implicitly by providing those methods. report accepts io.Writer, so it works with os.Stdout and *bytes.Buffer without knowing their concrete types. Define small interfaces near the consumer when you need substitution. Accepting an interface does not mean every struct needs one.

### Use type parameters for shared algorithms

Generics let an algorithm operate on a constrained set of types. The unique function requires comparable because map keys need equality. Prefer an ordinary concrete function until you have repeated logic whose behavior really is the same. A nil interface and an interface holding a typed nil pointer differ; return an explicit nil error on success.

### Build it yourself

1. Run the report example using the terminal and buffer.
2. Move report into an internal/report package and import it from your command.
3. Make unique work with integers and strings; explain the comparable constraint.
4. Try passing a type without Write to report and interpret the compile error.

```sh
go run ./examples/06-interfaces
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func report(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "report: %s\n", name)
	return err
}

func unique[T comparable](items []T) []T {
	seen := make(map[T]bool)
	out := make([]T, 0, len(items))
	for _, item := range items {
		if !seen[item] {
			out = append(out, item)
			seen[item] = true
		}
	}
	return out
}

func main() {
	var buffer bytes.Buffer
	if err := report(&buffer, "daily"); err != nil {
		panic(err)
	}
	fmt.Print(buffer.String())
	if err := report(os.Stdout, "weekly"); err != nil {
		panic(err)
	}
	fmt.Println(unique([]string{"go", "sql", "go"}))
}
```

</details>

### Checkpoints

- The report can be tested without intercepting the real terminal.
- Your package API exposes only what callers need.
- You can explain the difference between interfaces and generics.

**Watch for:** Import cycles are forbidden. If packages depend on each other, reconsider the ownership of shared concepts instead of adding more layers.

**Recall:** Why does bytes.Buffer work as an io.Writer without an implements declaration?

Its pointer type has Write([]byte) (int, error), the method io.Writer requires. Go checks that structural match.


Further reading: [Module tutorial](https://go.dev/doc/tutorial/create-module) · [Generics tutorial](https://go.dev/doc/tutorial/generics)


## 7. Read files and streams without loading everything

Operations gives you a multi-gigabyte log. Count matching lines without reading the entire file into memory.

### Separate data sources from processing

io.Reader describes bytes coming in; io.Writer describes bytes going out. A function accepting io.Reader can process a file, standard input, or an in-memory string. Let the caller open files and decide where output belongs. This also makes failure tests much simpler.

### Stream deliberately

bufio.Scanner reads tokens, usually lines, one at a time. A successful Scan makes Text available; after the loop check Err. The default token limit is too small for some logs, so set a deliberate maximum or use bufio.Reader for another strategy. Streaming reduces memory use, but unlimited line size can still create resource problems.

### Treat end-of-input differently from failure

EOF is the normal end of a stream, while a disk read error is a failed operation. A final line may have no newline. Report what your result means if the read fails halfway through: the reference command treats an incomplete scan as failure rather than printing a misleading success summary.

### Build it yourself

1. Write scan(reader, needle) returning line and match counts.
2. Run it with a file, then a pipe. Use the same processing function for both.
3. Try empty input, a final line with no newline, and a line above your chosen maximum.
4. Use a custom Reader that returns an error to observe the failure path.

```sh
go run ./examples/loglens examples/loglens/sample.log
# Expected: 2 matches in 5 lines
```


<details><summary>Reference / excerpt</summary>

```
// LogLens counts matching lines in a file or standard input.
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Summary struct {
	Lines   int `json:"lines"`
	Matches int `json:"matches"`
}

func scan(r io.Reader, needle string) (Summary, error) {
	var summary Summary
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	for scanner.Scan() {
		summary.Lines++
		if strings.Contains(scanner.Text(), needle) {
			summary.Matches++
		}
	}
	return summary, scanner.Err()
}

func run(args []string, in io.Reader, out, stderr io.Writer) int {
	fs := flag.NewFlagSet("loglens", flag.ContinueOnError)
	fs.SetOutput(stderr)
	needle := fs.String("contains", "ERROR", "substring to count")
	asJSON := fs.Bool("json", false, "write a JSON summary")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if fs.NArg() > 1 {
		fmt.Fprintln(stderr, "usage: loglens [flags] [file]")
		return 2
	}
	if fs.NArg() == 1 {
		file, err := os.Open(fs.Arg(0))
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		defer file.Close()
		in = file
	}
	summary, err := scan(in, *needle)
	if err != nil {
		fmt.Fprintln(stderr, "read log:", err)
		return 1
	}
	if *asJSON {
		err = json.NewEncoder(out).Encode(summary)
	} else {
		_, err = fmt.Fprintf(out, "%d matches in %d lines\n", summary.Matches, summary.Lines)
	}
	if err != nil {
		fmt.Fprintln(stderr, "write summary:", err)
		return 1
	}
	return 0
}

func main() { os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)) }
```

</details>

### Checkpoints

- Memory does not scale with the full input size.
- Read errors do not silently produce a successful report.
- The scanner’s line-size limit is explicit.

**Watch for:** Ignoring scanner.Err() hides truncated lines and read failures. Always inspect it after the scan loop.

**Recall:** Why accept io.Reader rather than a filename in the counting function?

It keeps opening and closing files outside the algorithm and allows the same code to read pipes, files, and test data.


Further reading: [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) · [io contracts](https://pkg.go.dev/io)


## 8. Design flags, output, and exit codes

Your tool must work both for a human in a terminal and in a scheduled script.

### Treat the CLI as an API

Document arguments, flags, defaults, and exit codes. flag.FlagSet with ContinueOnError lets run return a code rather than exiting deep inside parsing. In the standard flag package, flags normally precede positional arguments. For subcommands, give each command a separate FlagSet and dispatch explicitly.

### Separate useful output from diagnostics

stdout carries the result so a shell pipeline can consume it. stderr carries warnings and errors. JSON output should contain only JSON: banners and progress messages would break downstream tools. Stable output is a compatibility promise, just like an HTTP response schema.

### Keep process control at the edge

A run(args, in, out, stderr) int function can be tested with ordinary buffers. main is a thin wrapper calling os.Exit only after run returns, allowing cleanup in run to execute. A reasonable contract is 0 for success, 1 for operational failure, and 2 for invalid usage; document your own choices.

### Build it yourself

1. Add -contains and -json flags to LogLens.
2. Handle -help as success and unknown flags as a usage error.
3. Build the executable; redirect stdout and stderr into separate files.
4. Add a version subcommand without mixing version text into scan output.

```sh
go build -o bin/loglens ./examples/loglens
./bin/loglens -json examples/loglens/sample.log
./bin/loglens -help
```


<details><summary>Reference / excerpt</summary>

```
// LogLens counts matching lines in a file or standard input.
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Summary struct {
	Lines   int `json:"lines"`
	Matches int `json:"matches"`
}

func scan(r io.Reader, needle string) (Summary, error) {
	var summary Summary
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	for scanner.Scan() {
		summary.Lines++
		if strings.Contains(scanner.Text(), needle) {
			summary.Matches++
		}
	}
	return summary, scanner.Err()
}

func run(args []string, in io.Reader, out, stderr io.Writer) int {
	fs := flag.NewFlagSet("loglens", flag.ContinueOnError)
	fs.SetOutput(stderr)
	needle := fs.String("contains", "ERROR", "substring to count")
	asJSON := fs.Bool("json", false, "write a JSON summary")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if fs.NArg() > 1 {
		fmt.Fprintln(stderr, "usage: loglens [flags] [file]")
		return 2
	}
	if fs.NArg() == 1 {
		file, err := os.Open(fs.Arg(0))
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		defer file.Close()
		in = file
	}
	summary, err := scan(in, *needle)
	if err != nil {
		fmt.Fprintln(stderr, "read log:", err)
		return 1
	}
	if *asJSON {
		err = json.NewEncoder(out).Encode(summary)
	} else {
		_, err = fmt.Fprintf(out, "%d matches in %d lines\n", summary.Matches, summary.Lines)
	}
	if err != nil {
		fmt.Fprintln(stderr, "write summary:", err)
		return 1
	}
	return 0
}

func main() { os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)) }
```

</details>

### Checkpoints

- JSON mode can be parsed without stripping status messages.
- A missing file returns a nonzero exit code.
- Help documents the default substring and file/stdin behavior.

**Watch for:** go run may translate the program’s nonzero status into its own exit status. Build a binary when checking the exact CLI exit-code contract.

**Recall:** Where should a missing-file message go?

stderr. stdout is reserved for the report so scripts never confuse an error message with valid data.


Further reading: [flag package](https://pkg.go.dev/flag) · [os.Exit](https://pkg.go.dev/os#Exit)


## 9. Configuration and files that survive mistakes

Users need defaults, a configuration file, and environment overrides. A typo should fail visibly instead of silently changing behavior.

### Choose and document precedence

A useful order is defaults, config file, environment, then explicit flags. Do not let a flag’s default overwrite an environment value merely because the flag exists; track whether it was supplied. Parse strings into typed values once at startup and validate the final configuration.

### Decode data with intent

encoding/json maps exported struct fields using tags. DisallowUnknownFields catches misspelled keys. For a file that must contain one JSON object, decode again and require EOF. Decide whether absent, empty, and explicit zero values have different meanings; pointers or custom option types can represent that distinction.

### Write updates carefully

Replacing a file directly can leave it truncated after a crash. Write a temporary file in the same directory, set appropriate permissions, check write and close errors, then rename it. Atomic rename and crash durability are different concerns, and platform guarantees differ. Never put credentials into example config committed to Git.

### Build it yourself

1. Run the configuration example with and without APP_ADDR.
2. Read JSON from an optional file and reject unknown keys or trailing objects.
3. Add a -workers flag that overrides the environment only when explicitly supplied.
4. Save a settings file via a temporary file and test a failed write without damaging the old file.

```sh
go run ./examples/config
APP_ADDR=127.0.0.1:9090 go run ./examples/config
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Addr    string `json:"addr"`
	Workers int    `json:"workers"`
}

func main() {
	cfg := Config{Addr: "127.0.0.1:8080", Workers: 2}
	decoder := json.NewDecoder(strings.NewReader(`{"workers":4}`))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if addr := os.Getenv("APP_ADDR"); addr != "" {
		cfg.Addr = addr
	}
	if cfg.Workers < 1 {
		fmt.Fprintln(os.Stderr, "workers must be positive")
		os.Exit(1)
	}
	fmt.Printf("addr=%s workers=%d\n", cfg.Addr, cfg.Workers)
}
```

</details>

### Checkpoints

- Invalid worker counts fail before work begins.
- Configuration precedence is written down and demonstrated.
- A secret is never printed in a full config dump.

**Watch for:** Environment variables are strings, not validated settings. An invalid timeout must not fall back silently to an unlimited operation.

**Recall:** How can an omitted value differ from an explicit zero?

Use an optional representation, such as *int, when nil means absent and a pointer to 0 means explicitly zero. Validate after applying precedence.


Further reading: [encoding/json](https://pkg.go.dev/encoding/json) · [File operations](https://pkg.go.dev/os)


## 10. Test behavior with tables and useful failures

Change your log parser without breaking scripts that already depend on its output.

### Start from a behavior contract

A test is Go code in a file ending _test.go. TestXxx receives *testing.T. Arrange input, call the behavior, compare the result, and print enough context to diagnose a failure. Avoid tests that simply repeat the implementation: specify examples that would catch a plausible mistake.

### Use tables for related cases

A table of named inputs and expected results makes missing cases visible. t.Run gives each case a clear name. Include empty input, boundaries, malformed data, and missing final newlines. For io-based functions, strings.Reader and bytes.Buffer replace real files without mocking an operating system.

### Make tests isolated and repeatable

Use t.TempDir for files and t.Cleanup for resources. Avoid external websites, real credentials, shared ports, and fixed sleeps as synchronization. Integration tests with a real temporary database complement small unit tests. Coverage shows executed code; it does not prove you asserted the right behavior.

### Build it yourself

1. Read the reference tests and predict which change would make each fail.
2. Add a case where the needle is absent and a case with an empty needle.
3. Test unreadable input and a writer that fails when output is written.
4. Break line counting deliberately and verify the failure message helps locate the bug.

```sh
go test ./examples/loglens -v
go test ./examples/loglens -cover
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	cases := []struct {
		name, input, needle string
		want                Summary
	}{
		{"empty", "", "ERROR", Summary{}},
		{"mixed", "INFO start\nERROR failed\n", "ERROR", Summary{2, 1}},
		{"no newline", "ERROR failed", "ERROR", Summary{1, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := scan(strings.NewReader(tc.input), tc.needle)
			if err != nil || got != tc.want {
				t.Fatalf("got %+v, %v; want %+v", got, err, tc.want)
			}
		})
	}
}

func TestRunJSONAndUsage(t *testing.T) {
	var out, stderr bytes.Buffer
	code := run([]string{"-json"}, strings.NewReader("ERROR bad\n"), &out, &stderr)
	if code != 0 || out.String() != "{\"lines\":1,\"matches\":1}\n" {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), stderr.String())
	}
	if got := run([]string{"-unknown"}, strings.NewReader(""), &out, &stderr); got != 2 {
		t.Fatalf("usage code=%d", got)
	}
}

func FuzzScan(f *testing.F) {
	f.Add("INFO ok\nERROR failed\n", "ERROR")
	f.Fuzz(func(t *testing.T, input, needle string) {
		got, err := scan(strings.NewReader(input), needle)
		if err == nil && (got.Matches < 0 || got.Matches > got.Lines) {
			t.Fatalf("invalid summary: %+v", got)
		}
	})
}

func BenchmarkScan(b *testing.B) {
	input := strings.Repeat("INFO request complete\nERROR request failed\n", 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := scan(strings.NewReader(input), "ERROR"); err != nil {
			b.Fatal(err)
		}
	}
}
```

</details>

### Checkpoints

- Tests run without internet access or a fixed working-directory assumption.
- Failure output contains the case, actual result, and expected result.
- A regression case is added for every parser bug you find.

**Watch for:** Calling log.Fatal or os.Exit inside reusable code makes it hard to test and prevents caller-controlled cleanup.

**Recall:** What does a passing test suite actually tell you?

The behaviors asserted by those tests passed for the exercised inputs and environment. It does not prove correctness for every input or schedule.


Further reading: [testing package](https://pkg.go.dev/testing) · [Table-driven tests](https://go.dev/wiki/TableDrivenTests)


## 11. Integrate with the operating system and ship a binary

Your CLI needs to call another tool, honor Ctrl-C, and be usable without a development setup.

### Pass arguments without invoking a shell

exec.CommandContext runs a program with an argument slice. Passing user text as separate arguments avoids the shell interpretation that happens with sh -c or string-built commands. Still validate paths and intended operations. Capture stdout and stderr deliberately; neither is unlimited free storage.

### Make interruption part of the design

signal.NotifyContext turns an interrupt into cancellation. Long-running operations should observe that context. CommandContext can terminate its direct process when cancellation occurs; managing a whole child process tree needs platform-specific care. Document whether partial output remains usable after an interrupted command.

### Package a reproducible artifact

go build creates a binary for the target GOOS and GOARCH. Pure Go programs are straightforward to cross-compile, while C dependencies and runtime files need additional care. Version your releases, document supported platforms, and test the built binary—not only go run. Use go install for Go users, or publish binaries and checksums.

### Build it yourself

1. Run the subprocess example and then invoke a failing command.
2. Add a configurable command deadline and inspect the returned error.
3. Build LogLens and run it from another directory using an absolute path.
4. Create a release note showing usage, version, supported platforms, and exit codes.

```sh
go run ./examples/process
go build -trimpath -o bin/loglens ./examples/loglens
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, "go", "version").CombinedOutput()
	if err != nil {
		fmt.Printf("go version failed: %v (%s)\n", err, output)
		return
	}
	fmt.Print(string(output))
}
```

</details>

### Checkpoints

- User input is passed as arguments, never interpolated into a shell command.
- Cancellation has a visible, non-success outcome.
- The binary does not depend on the source working directory.

**Watch for:** Cross-compilation is not the same as verification on that platform. Record which artifacts you actually ran.

**Recall:** Why use exec.CommandContext(ctx, tool, arg) instead of a shell command string?

The argument stays one argument rather than becoming shell syntax, and the context supplies a cancellation boundary for the subprocess.


Further reading: [os/exec](https://pkg.go.dev/os/exec) · [go build](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)


## 12. Project: ship LogLens to another developer

Build a log-analysis CLI a teammate can use to investigate failures in a production export.

### Start from user stories

The operator should scan a file or piped input, filter by a substring, and receive text or machine-readable output. The program should handle large files, bad flags, missing files, and oversized lines predictably. Write these requirements before adding features.

### Build in small increments

First stream and count lines. Next add flags and stable JSON. Then separate run from main and test both parsing and exit behavior. The supplied reference implementation is deliberately small: understand every line before extending it. Add features one at a time and preserve the command contract.

### Practice a real handoff

A professional deliverable includes source, a binary or reproducible build, help, sample input, examples, and a short explanation of limitations. A teammate should be able to use it without asking you where the input comes from. Keep a change log when modifying output fields or defaults.

### Build it yourself

1. Rebuild the tool in a fresh module using the requirements; use the reference only after your first attempt.
2. Add -ignore-case and test its interaction with both text and JSON output.
3. Add a top-N count of error messages with deterministic tie ordering.
4. Give the binary and README to another person or evaluate it from a clean folder using only your instructions.

```sh
go test ./examples/loglens
go build -o bin/loglens ./examples/loglens
./bin/loglens -contains ERROR -json examples/loglens/sample.log
```


<details><summary>Reference / excerpt</summary>

```
// LogLens counts matching lines in a file or standard input.
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Summary struct {
	Lines   int `json:"lines"`
	Matches int `json:"matches"`
}

func scan(r io.Reader, needle string) (Summary, error) {
	var summary Summary
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	for scanner.Scan() {
		summary.Lines++
		if strings.Contains(scanner.Text(), needle) {
			summary.Matches++
		}
	}
	return summary, scanner.Err()
}

func run(args []string, in io.Reader, out, stderr io.Writer) int {
	fs := flag.NewFlagSet("loglens", flag.ContinueOnError)
	fs.SetOutput(stderr)
	needle := fs.String("contains", "ERROR", "substring to count")
	asJSON := fs.Bool("json", false, "write a JSON summary")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}
	if fs.NArg() > 1 {
		fmt.Fprintln(stderr, "usage: loglens [flags] [file]")
		return 2
	}
	if fs.NArg() == 1 {
		file, err := os.Open(fs.Arg(0))
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
		defer file.Close()
		in = file
	}
	summary, err := scan(in, *needle)
	if err != nil {
		fmt.Fprintln(stderr, "read log:", err)
		return 1
	}
	if *asJSON {
		err = json.NewEncoder(out).Encode(summary)
	} else {
		_, err = fmt.Fprintf(out, "%d matches in %d lines\n", summary.Matches, summary.Lines)
	}
	if err != nil {
		fmt.Fprintln(stderr, "write summary:", err)
		return 1
	}
	return 0
}

func main() { os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)) }
```

</details>

### Checkpoints

- Reference sample: 2 matches across 5 lines.
- A 100 MB input does not require a 100 MB buffer.
- Help, exact exit codes, JSON shape, and error cases are documented and checked.
- You can explain the design without reading the code.

**Watch for:** Adding a CLI framework does not fix unclear command behavior. Introduce one only when the command tree or completion needs justify it.

**Recall:** What would demonstrate that you can build a CLI independently?

You can implement a new command from a short requirement, exercise its failure paths, test it through the CLI boundary, and deliver usable documentation.


Further reading: [Go command documentation](https://pkg.go.dev/cmd/go)


## 13. Understand the request and response lifecycle

Build an issue tracker that receives requests from both browsers and automation.

### HTTP is a contract between two programs

A request has a method, URL, headers, and optional body. A response has a status, headers, and optional body. net/http calls a handler with a ResponseWriter and a Request. Register routes with a ServeMux; method-aware patterns and PathValue let you distinguish operations without writing your own router.

### Headers come before the body

Set Content-Type and other headers before WriteHeader or writing bytes. Writing a body implicitly starts a 200 response if you have not chosen a status. Once bytes are sent, you cannot replace the response with a different status. Use 201 for a created resource, 400 for invalid input, and 404 for a missing resource as part of a documented contract.

### A server handles overlapping requests

Handlers may run concurrently. State stored on an app struct is shared by those requests; passing a pointer does not make it synchronized. Keep mutable state in a safe store or protect it. A production server also needs header/body limits, timeouts, configuration, and shutdown behavior; later lessons develop each.

### Build it yourself

1. Start the supplied TeamDesk app and open http://localhost:8080/.
2. Use curl -i on /health and /api/tickets and identify the status, headers, and body.
3. Create a small server from scratch with GET /health and GET /hello/{name}.
4. Try a wrong method, unknown route, and extra path segment; explain the responses.

```sh
go run ./examples/teamdesk -db /tmp/teamdesk-learning.db
# In another terminal:
curl -i http://localhost:8080/health
```


<details><summary>Reference / excerpt</summary>

```
// TeamDesk is a local teaching app. Add identity and authorization before exposing it.
package main

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

//go:embed page.html
var pageHTML string
var page = template.Must(template.New("tickets").Parse(pageHTML))

type app struct{ store *Store }

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /tickets", a.createForm)
	mux.HandleFunc("POST /tickets/{id}/complete", a.complete)
	mux.HandleFunc("GET /api/tickets", a.listAPI)
	mux.HandleFunc("POST /api/tickets", a.createAPI)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	// Modern browsers send Origin / Sec-Fetch-Site. This protects unsafe cross-origin requests.
	return http.NewCrossOriginProtection().Handler(mux)
}

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	var buf bytes.Buffer
	if err := page.Execute(&buf, tickets); err != nil {
		http.Error(w, "cannot render page", 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}

func (a *app) createForm(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", 400)
		return
	}
	title := strings.TrimSpace(r.PostForm.Get("title"))
	if title == "" || len(title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	if _, err := a.store.Create(r.Context(), title); err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) complete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "invalid id", 400)
		return
	}
	err = a.store.Complete(r.Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "cannot update ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) listAPI(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tickets); err != nil {
		slog.Error("encode tickets", "error", err)
	}
}

func (a *app) createAPI(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input struct {
		Title string `json:"title"`
	}
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "invalid JSON", 400)
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		http.Error(w, "send one JSON object", 400)
		return
	}
	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" || len(input.Title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	ticket, err := a.store.Create(r.Context(), input.Title)
	if err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		slog.Error("encode ticket", "error", err)
	}
}

func run() error {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	dbPath := flag.String("db", "teamdesk.db", "SQLite file")
	flag.Parse()
	store, err := openStore(*dbPath)
	if err != nil {
		return err
	}
	defer store.db.Close()
	srv := &http.Server{Addr: *addr, Handler: (&app{store}).routes(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	result := make(chan error, 1)
	go func() { result <- srv.ListenAndServe() }()
	slog.Info("TeamDesk started", "addr", *addr)
	select {
	case err := <-result:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			_ = srv.Close()
			return err
		}
		return nil
	}
}

func main() {
	if err := run(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
```

</details>

### Checkpoints

- A handler sets headers before its body.
- The exact home route does not swallow all unmatched paths.
- You can trace one request from mux to handler to storage to response.

**Watch for:** The pattern / matches a subtree. Use /{$} for an exact root route when that is what you intend.

**Recall:** What happens if you call http.Error after already writing a response body?

The original status may already be sent. You can append bytes but cannot reliably turn that response into the intended error response.


Further reading: [net/http](https://pkg.go.dev/net/http) · [Routing enhancements](https://go.dev/blog/routing-enhancements)


## 14. Build an API with a precise input contract

A client creates tickets. Unknown fields, multiple JSON objects, and invalid titles should be rejected consistently.

### Separate input from stored data

Define an input struct containing only fields the caller may set. Do not decode directly into a database model with owner IDs, roles, or server-generated timestamps. A dedicated output representation also lets your public contract evolve independently of storage.

### Parsing is only the first validation step

Limit the request body before decoding. Disallow unknown fields if your contract requires strictness. Decode one object, then require EOF so a second object is not silently ignored. After parsing, validate meaningful domain constraints: a trimmed title must not be empty and should have a deliberate length policy.

### Make errors useful without leaking internals

Clients need stable status codes and, in a mature API, machine-readable error codes with field details. Logs can retain the wrapped internal cause. Do not return SQL statements, stack traces, tokens, or raw dependency errors. Distinguish malformed requests, forbidden actions, missing resources, and transient service failure.

### Build it yourself

1. Create a ticket with POST /api/tickets, then list it with GET /api/tickets.
2. Try empty titles, unknown fields, invalid JSON, oversized bodies, and two objects.
3. Add GET /api/tickets/{id}; distinguish invalid ID syntax from a well-formed missing ID.
4. Refactor the reference app’s text error responses into a stable JSON error envelope for API routes.

```sh
curl -i -X POST http://localhost:8080/api/tickets -H 'Content-Type: application/json' -d '{"title":"Fix export"}'
curl -i http://localhost:8080/api/tickets
```


<details><summary>Reference / excerpt</summary>

```
// TeamDesk is a local teaching app. Add identity and authorization before exposing it.
package main

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

//go:embed page.html
var pageHTML string
var page = template.Must(template.New("tickets").Parse(pageHTML))

type app struct{ store *Store }

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /tickets", a.createForm)
	mux.HandleFunc("POST /tickets/{id}/complete", a.complete)
	mux.HandleFunc("GET /api/tickets", a.listAPI)
	mux.HandleFunc("POST /api/tickets", a.createAPI)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	// Modern browsers send Origin / Sec-Fetch-Site. This protects unsafe cross-origin requests.
	return http.NewCrossOriginProtection().Handler(mux)
}

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	var buf bytes.Buffer
	if err := page.Execute(&buf, tickets); err != nil {
		http.Error(w, "cannot render page", 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}

func (a *app) createForm(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", 400)
		return
	}
	title := strings.TrimSpace(r.PostForm.Get("title"))
	if title == "" || len(title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	if _, err := a.store.Create(r.Context(), title); err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) complete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "invalid id", 400)
		return
	}
	err = a.store.Complete(r.Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "cannot update ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) listAPI(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tickets); err != nil {
		slog.Error("encode tickets", "error", err)
	}
}

func (a *app) createAPI(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input struct {
		Title string `json:"title"`
	}
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "invalid JSON", 400)
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		http.Error(w, "send one JSON object", 400)
		return
	}
	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" || len(input.Title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	ticket, err := a.store.Create(r.Context(), input.Title)
	if err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		slog.Error("encode ticket", "error", err)
	}
}

func run() error {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	dbPath := flag.String("db", "teamdesk.db", "SQLite file")
	flag.Parse()
	store, err := openStore(*dbPath)
	if err != nil {
		return err
	}
	defer store.db.Close()
	srv := &http.Server{Addr: *addr, Handler: (&app{store}).routes(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	result := make(chan error, 1)
	go func() { result <- srv.ListenAndServe() }()
	slog.Info("TeamDesk started", "addr", *addr)
	select {
	case err := <-result:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			_ = srv.Close()
			return err
		}
		return nil
	}
}

func main() {
	if err := run(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
```

</details>

### Checkpoints

- Clients cannot set a server-owned ID through the request body.
- Malformed or ambiguous input is rejected.
- An empty collection serializes as [] rather than null.

**Watch for:** CORS controls which browser origins can read responses. It is not authentication or authorization, and it does not protect an API from non-browser clients.

**Recall:** Why decode into an input type rather than a full model?

It explicitly defines what the caller may change and prevents unintended field assignment as the stored model grows.


Further reading: [JSON decoder](https://pkg.go.dev/encoding/json#Decoder)


## 15. Persist data with database/sql

Tickets must survive a restart, and a failed query should not corrupt application state.

### The driver and pool have different jobs

database/sql provides a common API and manages a pool of connections. A driver implements database-specific communication. The blank import of modernc.org/sqlite registers the SQLite driver. sql.Open configures a handle; use PingContext or an actual query to establish that the database is reachable.

### Use parameters and release rows

Pass values separately from SQL text. SQLite uses ? placeholders in this example; PostgreSQL drivers commonly use $1, $2. Do not assemble SQL by interpolating user input. QueryContext returns rows: defer Close, scan each row, and check rows.Err after iteration. ExecContext handles writes without a result set.

### Make ownership and errors explicit

A Store owns the database dependency. Handlers pass their request context, and the process closes the pool after request work finishes. Translate sql.ErrNoRows into the application’s missing-resource behavior. The teaching app uses one SQLite connection for simplicity; revisit connection counts, busy handling, and transactions under load.

### Build it yourself

1. Read store.go, then create, list, and complete tickets using the UI.
2. Stop and restart the app with the same database path; confirm the data remains.
3. Add a due-date column with a migration rather than deleting an existing database.
4. Add a database test with t.TempDir and prove missing IDs do not report success.

```sh
go test ./examples/teamdesk -v
go run ./examples/teamdesk -db /tmp/teamdesk-learning.db
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type Ticket struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
type Store struct{ db *sql.DB }

func openStore(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// One connection keeps this small SQLite lesson simple. Revisit pooling later.
	db.SetMaxOpenConns(1)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tickets (
        id INTEGER PRIMARY KEY, title TEXT NOT NULL, done INTEGER NOT NULL DEFAULT 0
    )`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("create schema: %w", err)
	}
	return &Store{db: db}, nil
}

func (s *Store) Create(ctx context.Context, title string) (Ticket, error) {
	result, err := s.db.ExecContext(ctx, "INSERT INTO tickets(title) VALUES (?)", title)
	if err != nil {
		return Ticket{}, fmt.Errorf("create ticket: %w", err)
	}
	id, err := result.LastInsertId()
	return Ticket{ID: id, Title: title}, err
}

func (s *Store) List(ctx context.Context) ([]Ticket, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, done FROM tickets ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("list tickets: %w", err)
	}
	defer rows.Close()
	tickets := make([]Ticket, 0)
	for rows.Next() {
		var t Ticket
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, rows.Err()
}

func (s *Store) Complete(ctx context.Context, id int64) error {
	result, err := s.db.ExecContext(ctx, "UPDATE tickets SET done=1 WHERE id=?", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
```

</details>

### Checkpoints

- Every query uses arguments for caller-controlled values.
- Rows are closed and iteration errors checked.
- Stored tickets survive a restart.

**Watch for:** Running several database calls is not automatically a transaction. Use a transaction when multiple changes must succeed together.

**Recall:** Why is sql.DB usually created once rather than once per request?

It is a concurrency-safe connection pool intended for reuse. Recreating it loses pooling and complicates resource ownership.


Further reading: [Go database guide](https://go.dev/doc/database/) · [SQLite driver](https://pkg.go.dev/modernc.org/sqlite)


## 16. Serve HTML, forms, and embedded assets

Give the ticket API a useful browser interface without requiring a separate frontend application.

### Templates combine trusted structure with untrusted data

html/template escapes values according to HTML context. Parse trusted templates at startup. Do not convert user content into template.HTML to bypass escaping. Render into a buffer first when template execution can fail, so you can still send a clean error status before any response bytes are committed.

### Forms have a lifecycle

GET displays a page. POST validates submitted data and changes state. After a successful POST, redirect with 303 so a refresh repeats the GET rather than submitting the form again. Keep GET operations free of state changes. Validation failures should preserve useful input and give clear field feedback.

### Bundle only what the app needs

go:embed includes templates or static assets in the binary, avoiding working-directory surprises. A simple server-rendered app can later add small JavaScript interactions while keeping the same server validation and authorization. Accessibility still matters: use labels, keyboard-operable controls, visible errors, and meaningful page titles.

### Build it yourself

1. Open TeamDesk, add a task, complete it, and refresh the page.
2. Add a title containing HTML-like text and confirm it displays as text.
3. Add a filter for open/completed tickets using query parameters.
4. Add validation feedback to the form instead of returning only plain text errors.

```sh
go build -o bin/teamdesk ./examples/teamdesk
./bin/teamdesk -db /tmp/teamdesk-learning.db
```


<details><summary>Reference / excerpt</summary>

```
<!doctype html>
<html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>TeamDesk</title>
<style>body{font:18px/1.6 system-ui;max-width:680px;margin:60px auto;padding:20px}input,button{font:inherit;padding:8px}li{margin:16px 0}form{display:inline}small{color:#666}</style>
<h1>TeamDesk</h1><p>A small issue tracker. Add a task, complete it, restart the app: your data stays.</p>
<form method="post" action="/tickets"><label>New task <input name="title" required maxlength="200"></label><button>Add</button></form>
<ul>{{range .}}<li>{{if .Done}}<s>{{.Title}}</s> <small>Done</small>{{else}}{{.Title}} <form method="post" action="/tickets/{{.ID}}/complete"><button>Complete</button></form>{{end}}</li>{{else}}<li>No tasks yet.</li>{{end}}</ul>
</html>
```

</details>

### Checkpoints

- Refreshing after a successful form submission does not create a duplicate.
- User-controlled text is escaped in HTML.
- The built binary can find its template from another working directory.

**Watch for:** html/template prevents many output-escaping errors; it does not implement login, access control, or CSRF protection by itself.

**Recall:** Why redirect after a successful form POST?

It moves the browser to a GET view so refresh and bookmarking do not repeat the state-changing form submission.


Further reading: [html/template](https://pkg.go.dev/html/template) · [Embedding files](https://pkg.go.dev/embed)


## 17. Add identity, sessions, and authorization

A shared ticket system must ensure people can only read and change records they are allowed to access.

### Separate three questions

Authentication establishes who made the request. Authorization decides whether that identity may perform this operation on this resource. Session management connects later browser requests to the authenticated user. Checking that someone is logged in is not enough: every resource lookup needs its access policy applied.

### Use established identity mechanisms

For an enterprise app, use a maintained OIDC integration with your identity provider and validate issuer, audience, signature, expiry, state, and nonce as appropriate. For a browser session, use an unpredictable opaque identifier, store session state server-side, rotate it on login, and expire or revoke it. Secure, HttpOnly, and suitable SameSite cookies have separate roles.

### Protect state changes and outputs

Cookie authentication needs CSRF protection, such as a tested token mechanism or carefully configured origin protection for supported clients. The reference app uses Go’s CrossOriginProtection for unsafe browser requests, but intentionally has no login. Add authorization before sharing it. Use HTTPS, bound request sizes, parameterized SQL, and logging that excludes credentials.

### Build it yourself

1. Write an access matrix: anonymous, member, and administrator versus view, create, complete, and delete.
2. Add session middleware through a maintained provider/session library; do not invent cryptographic token signing.
3. Make user B unable to view or mutate user A’s ticket by guessing its ID.
4. Test cross-origin form submission, session expiry, logout, and an unauthorized object ID.

<details><summary>Reference / excerpt</summary>

```
session := authenticate(r)
// The query must include the verified identity or tenant.
ticket, err := store.FindForUser(r.Context(), session.UserID, ticketID)
// Return a consistent forbidden/not-found response without exposing other users’ data.
```

</details>

### Checkpoints

- Identity comes from verified session state, not a caller-supplied owner ID.
- Every protected operation checks resource access.
- The local reference app’s limitations are understood before deployment.

**Watch for:** A signed JWT is not automatically safe: validation, key rotation, revocation strategy, storage, and audience all matter.

**Recall:** Why is a logged-in user still able to cause an access-control bug?

Authentication proves identity, but a missing ownership or role check can let that identity access another user’s resource.


Further reading: [OWASP session management](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html) · [Go cross-origin protection](https://pkg.go.dev/net/http#CrossOriginProtection)


## 18. Project: build and extend TeamDesk

Deliver an issue tracker with a browser UI, API, persistent storage, and a clear security boundary.

### Build one complete vertical feature

Start with creating a ticket: form or JSON input, validation, store call, response, then a test. Continue with listing and completing. You should be able to explain the full path before extracting abstractions. The supplied implementation includes persistence, HTML escaping, basic cross-origin protection, and graceful Ctrl-C shutdown.

### Grow beyond the reference

The reference is a local single-user learning application. To make it shared, add authenticated users, per-user or per-tenant authorization, migrations, and security-focused tests. Add error handling and logs that help an operator without exposing internal state to clients. Prefer a complete small feature over many unfinished routes.

### Write acceptance scenarios

A new user can create and complete a ticket; data survives restart; a malformed request is rejected; one user cannot read another’s data; shutdown does not close the database under active requests. Automate the important scenarios and document what remains unsupported.

### Build it yourself

1. Rebuild the create/list/complete flow without copying the reference first.
2. Add edit and delete operations with validation and tests.
3. Implement the access matrix from the previous lesson and verify tenant/user isolation.
4. Write a short README with setup, migrations, API examples, security assumptions, backup, and restore.

```sh
go test -race ./examples/teamdesk
go run ./examples/teamdesk -db /tmp/teamdesk-learning.db
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestAPIAndTemplates(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tickets.db")
	s, err := openStore(path)
	if err != nil {
		t.Fatal(err)
	}
	defer s.db.Close()
	h := (&app{s}).routes()
	for _, tc := range []struct {
		body   string
		status int
	}{
		{`{"title":"export"}`, 201}, {`{"title":""}`, 400},
		{`{"title":"x","admin":true}`, 400}, {`{"title":"x"}{"title":"y"}`, 400},
	} {
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, httptest.NewRequest("POST", "/api/tickets", strings.NewReader(tc.body)))
		if rec.Code != tc.status {
			t.Errorf("body=%s status=%d want=%d", tc.body, rec.Code, tc.status)
		}
	}
	if _, err = s.Create(context.Background(), "<script>alert(1)</script>"); err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/", nil))
	if rec.Code != 200 || strings.Contains(rec.Body.String(), "<script>alert") {
		t.Fatalf("template escaped incorrectly: %s", rec.Body.String())
	}
	rec = httptest.NewRecorder()
	req := httptest.NewRequest("POST", "http://example.com/tickets", strings.NewReader("title=x"))
	req.Header.Set("Origin", "https://attacker.example")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("cross-origin POST status=%d", rec.Code)
	}
	if err := s.db.Close(); err != nil {
		t.Fatal(err)
	}
	reopened, err := openStore(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.db.Close()
	tickets, err := reopened.List(context.Background())
	if err != nil || len(tickets) != 2 {
		t.Fatalf("persisted=%v err=%v", tickets, err)
	}
}
```

</details>

### Checkpoints

- Both the browser UI and API are usable.
- The database survives restart and has a migration path.
- Authorization is checked at the resource boundary.
- A new developer can run the app and tests using the README.

**Watch for:** An enterprise web app is not defined by how many layers it has. Clear business rules, safe data access, and operational behavior matter more.

**Recall:** When is it reasonable to add a router or web framework?

When it solves demonstrated requirements such as routing or middleware ergonomics, and you still understand the HTTP, state, and security behavior underneath.


Further reading: [HTTP testing tools](https://pkg.go.dev/net/http/httptest)


## 19. Understand goroutines and shared state

A request counter is correct in a single request but loses increments when traffic overlaps.

### Concurrency is overlapping work

A goroutine runs a function concurrently with other goroutines. It is not a promise of a separate CPU core or of faster execution. The go statement does not wait for completion; when main returns, the process ends. Decide who starts each goroutine, how it stops, and who waits for it.

### Shared memory needs synchronization

Incrementing an integer is a read-modify-write operation. Two goroutines can read the same old value and overwrite one another. sync.Mutex protects a critical section; sync.WaitGroup waits for a known set of work. Call Add before starting the goroutine and defer Done in it. Atomics fit some individual operations but do not make a multi-field invariant atomic.

### Reason about ownership first

The simplest concurrency design often gives each value one owner, or passes independent copies. Use a mutex when several goroutines need shared mutable state. Do not copy a mutex after use, and do not hold a lock while doing slow network I/O unless that is a deliberate requirement. Run the race detector on executed concurrent paths.

### Build it yourself

1. Run the counter example and explain why the final result is 100.
2. Temporarily remove the mutex in your scratch copy and use go run -race to observe the failure.
3. Put the mutex back and write a concurrent test that checks the total.
4. Describe the invariant the lock protects and which functions must acquire it.

```sh
go run -race ./examples/counter
go test -race ./examples/counter
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Add()       { c.mu.Lock(); defer c.mu.Unlock(); c.value++ }
func (c *Counter) Value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.value }

func main() {
	var c Counter
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Add() }()
	}
	wg.Wait()
	fmt.Println(c.Value())
}
```

</details>

### Checkpoints

- All goroutines finish before their result is inspected.
- Shared mutable state has an explicit owner or synchronization strategy.
- You can distinguish a data race from a broader business-level race.

**Watch for:** The race detector only sees paths that run. A clean run does not prove the absence of every possible race.

**Recall:** Why is checking a value and then updating it under two separate locks unsafe?

Another goroutine can change the value between those steps. The entire invariant-preserving operation needs one critical section or equivalent atomic control.


Further reading: [Race detector](https://go.dev/doc/articles/race_detector) · [sync package](https://pkg.go.dev/sync)


## 20. Bound work with channels and worker pools

A batch has 50,000 jobs. Starting a goroutine for every job would overload the destination and grow memory use.

### Channels transfer values and coordinate work

An unbuffered send waits for a receiver. A buffered channel can hold a bounded number of values. Buffering changes where backpressure appears; it does not remove capacity limits. A worker pool starts a fixed number of workers and feeds them jobs, bounding active work even when the input is large.

### Give closing responsibility to the producer

The code that knows no more values will be sent should close the channel. Receivers can range until it closes, or use value, ok := <-ch. A receive from a closed channel immediately returns the zero value with ok false. Sending to a closed channel panics. A nil channel blocks forever unless excluded by another select case.

### Make every blocking edge stoppable

Cancellation must be able to unblock job production, waiting workers, and result publication. The reference example uses select around sends and receives. One coordinator waits for workers and closes results, allowing the caller to range safely. Decide how overload is handled: wait, reject, shed optional work, or persist it elsewhere.

### Build it yourself

1. Run the worker example and note that output order may vary.
2. Change the number of workers and confirm each of the five jobs produces exactly one result.
3. Cancel while producing work and check that all goroutines can exit.
4. Replace integer work with a function that returns a result and error; keep failures observable.

```sh
go run ./examples/workers
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			select {
			case results <- job * job:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jobs, results := make(chan int), make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, jobs, results, &wg)
	}
	go func() {
		defer close(jobs)
		for i := 1; i <= 5; i++ {
			select {
			case jobs <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() { wg.Wait(); close(results) }()
	for result := range results {
		fmt.Println(result)
	}
}
```

</details>

### Checkpoints

- The maximum active work is bounded by the worker count.
- Closed job channels do not produce endless zero-value jobs.
- The receiver never closes a channel while producers may still send.

**Watch for:** Channels are not automatically better than a mutex. Use a channel for ownership transfer or coordination; use a lock for a small shared-state operation.

**Recall:** Who should close the results channel when several workers send to it?

A coordinator that knows all workers have finished, typically after WaitGroup.Wait, rather than any individual worker.


Further reading: [Pipelines and cancellation](https://go.dev/blog/pipelines)


## 21. Propagate cancellation and shut down cleanly

A user abandons a slow request. A deployment also needs to stop the server without dropping accepted work.

### Context carries an operation lifetime

Pass context.Context as the first parameter through calls that can block or perform external work. Derive a deadline with WithTimeout and defer its cancel function. Context cancellation is a signal: the called code must observe it or use an API that does. Avoid storing a request context on a long-lived struct.

### Request work and durable work have different lifetimes

A database query answering a request should normally use r.Context. A job already accepted into a durable queue must not use that request’s context after the response ends. Give background processing an application lifecycle context and a separate per-attempt timeout. Do not hide required dependencies in context values.

### Shutdown is a sequence

Stop new intake, allow bounded time for in-flight requests and jobs, and leave unfinished durable jobs recoverable. http.Server.Shutdown waits for ordinary active HTTP requests; it does not automatically wait for unrelated background goroutines or hijacked connections. Use a fresh timeout context for draining, not the already-cancelled signal context.

### Build it yourself

1. Read the TeamDesk run function and identify startup, serve, and shutdown ownership.
2. Add a handler whose database operation observes a short context deadline.
3. Press Ctrl-C during a slow request and observe whether the process drains or forces close after its deadline.
4. Write down shutdown order for an API, worker pool, and database.

```sh
go run ./examples/teamdesk -db /tmp/teamdesk-learning.db
# Submit a request, then press Ctrl-C and observe shutdown.
```


<details><summary>Reference / excerpt</summary>

```
// TeamDesk is a local teaching app. Add identity and authorization before exposing it.
package main

import (
	"bytes"
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

//go:embed page.html
var pageHTML string
var page = template.Must(template.New("tickets").Parse(pageHTML))

type app struct{ store *Store }

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /tickets", a.createForm)
	mux.HandleFunc("POST /tickets/{id}/complete", a.complete)
	mux.HandleFunc("GET /api/tickets", a.listAPI)
	mux.HandleFunc("POST /api/tickets", a.createAPI)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	// Modern browsers send Origin / Sec-Fetch-Site. This protects unsafe cross-origin requests.
	return http.NewCrossOriginProtection().Handler(mux)
}

func (a *app) home(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	var buf bytes.Buffer
	if err := page.Execute(&buf, tickets); err != nil {
		http.Error(w, "cannot render page", 500)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}

func (a *app) createForm(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", 400)
		return
	}
	title := strings.TrimSpace(r.PostForm.Get("title"))
	if title == "" || len(title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	if _, err := a.store.Create(r.Context(), title); err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) complete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "invalid id", 400)
		return
	}
	err = a.store.Complete(r.Context(), id)
	if errors.Is(err, sql.ErrNoRows) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, "cannot update ticket", 500)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *app) listAPI(w http.ResponseWriter, r *http.Request) {
	tickets, err := a.store.List(r.Context())
	if err != nil {
		http.Error(w, "cannot load tickets", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tickets); err != nil {
		slog.Error("encode tickets", "error", err)
	}
}

func (a *app) createAPI(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 4096)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var input struct {
		Title string `json:"title"`
	}
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "invalid JSON", 400)
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		http.Error(w, "send one JSON object", 400)
		return
	}
	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" || len(input.Title) > 200 {
		http.Error(w, "title must be 1–200 bytes", 400)
		return
	}
	ticket, err := a.store.Create(r.Context(), input.Title)
	if err != nil {
		http.Error(w, "cannot create ticket", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		slog.Error("encode ticket", "error", err)
	}
}

func run() error {
	addr := flag.String("addr", "127.0.0.1:8080", "listen address")
	dbPath := flag.String("db", "teamdesk.db", "SQLite file")
	flag.Parse()
	store, err := openStore(*dbPath)
	if err != nil {
		return err
	}
	defer store.db.Close()
	srv := &http.Server{Addr: *addr, Handler: (&app{store}).routes(), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	result := make(chan error, 1)
	go func() { result <- srv.ListenAndServe() }()
	slog.Info("TeamDesk started", "addr", *addr)
	select {
	case err := <-result:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			_ = srv.Close()
			return err
		}
		return nil
	}
}

func main() {
	if err := run(); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
```

</details>

### Checkpoints

- Every derived cancel function is released.
- The shutdown timeout is based on a still-live context.
- The database stays open until dependent work has stopped.

**Watch for:** select does not prioritize cancellation over other ready cases. Design shutdown semantics explicitly if no new job may start once draining begins.

**Recall:** Why not pass r.Context to a background delivery that should run after returning 202?

The request context ends with the request. Durable accepted work needs its own lifecycle so a client disconnect cannot erase the processing promise.


Further reading: [Context package](https://pkg.go.dev/context) · [Cancel database operations](https://go.dev/doc/database/cancel-operations)


## 22. Call external services with limits

Your application checks several partner endpoints. One is slow, one returns 500, and one closes the connection.

### Reuse the client and transport

http.Client is safe for concurrent use and normally should be reused. Its transport manages connection reuse. A zero client timeout can wait indefinitely unless another deadline bounds the call. Combine an overall client policy with per-operation context deadlines as appropriate.

### HTTP failures are not all Go errors

client.Do returns an error for failures such as transport problems; a 500 response is still a successfully received HTTP response. Inspect the status code separately. Close every response body. Reading an appropriate bounded amount and closing enables safer resource use; fully consuming small bodies helps connection reuse.

### Bound more than time

Limit response size, concurrency, redirects, and retry attempts. The Sitecheck CLI intentionally accepts operator-supplied URLs. If a public server accepts arbitrary destinations, address SSRF: restrict schemes and destinations, validate resolved IPs and redirects, and enforce network egress rules. A hostname string check alone is insufficient.

### Build it yourself

1. Start TeamDesk locally and use Sitecheck to probe its health and missing routes.
2. Use httptest.NewServer to simulate 200, 500, delayed responses, and malformed bodies.
3. Add structured JSON output with URL, status, elapsed time, and categorized error.
4. Explain which limits protect latency, memory, and the remote service.

```sh
go run ./examples/sitecheck -workers 2 -timeout 2s http://localhost:8080/health http://localhost:8080/missing
```


<details><summary>Reference / excerpt</summary>

```
// Sitecheck probes operator-supplied URLs with bounded concurrency.
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"sync"
	"time"
)

type Result struct {
	URL    string
	Status int
	Err    error
}

func check(ctx context.Context, client *http.Client, urls []string, workers int) []Result {
	if workers < 1 {
		workers = 1
	}
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				result := Result{URL: url}
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
				if err == nil {
					var resp *http.Response
					resp, err = client.Do(req)
					if err == nil {
						result.Status = resp.StatusCode
						resp.Body.Close()
					}
				}
				result.Err = err
				results <- result
			}
		}()
	}
	go func() {
		defer close(jobs)
		for _, url := range urls {
			select {
			case jobs <- url:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() { wg.Wait(); close(results) }()
	out := make([]Result, 0, len(urls))
	for result := range results {
		out = append(out, result)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].URL < out[j].URL })
	return out
}

func main() {
	workers := flag.Int("workers", 3, "maximum simultaneous requests")
	timeout := flag.Duration("timeout", 3*time.Second, "timeout for each request")
	flag.Parse()
	if *workers < 1 || *workers > 64 || *timeout <= 0 || flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: sitecheck -workers 3 -timeout 3s URL...")
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	client := &http.Client{Timeout: *timeout}
	failed := false
	for _, result := range check(ctx, client, flag.Args(), *workers) {
		if result.Err != nil {
			fmt.Printf("FAIL %s: %v\n", result.URL, result.Err)
			failed = true
			continue
		}
		fmt.Printf("%d %s\n", result.Status, result.URL)
		if result.Status >= 400 {
			failed = true
		}
	}
	if failed || ctx.Err() != nil {
		os.Exit(1)
	}
}
```

</details>

### Checkpoints

- Slow destinations have a deadline.
- A 500 is classified as an HTTP failure, not mistakenly treated as success.
- Response bodies are closed on every successful Do call.

**Watch for:** Creating a fresh transport for each request loses pooling and can exhaust resources. Reuse it unless there is a specific isolation requirement.

**Recall:** Does err == nil after client.Do mean the server returned a successful status?

No. It means an HTTP response was obtained. Inspect resp.StatusCode to decide whether the application operation succeeded.


Further reading: [http.Client](https://pkg.go.dev/net/http#Client)


## 23. Retry without multiplying side effects

A partner processes a request, but its response is lost. Retrying might create a duplicate invoice or notification.

### A timeout creates uncertainty

The caller may not know whether an operation failed before it ran or after it committed. Retrying a state-changing request is safe only when the operation is idempotent or the receiver deduplicates a stable operation ID. Sending an Idempotency-Key header helps only if the receiver implements that contract.

### Use a retry budget

Choose retryable conditions deliberately. Transient network failures, selected 5xx responses, and sometimes 429 may warrant retry. Validation and authorization errors usually do not. Honor Retry-After where supported. Bound attempts and total elapsed time, use exponential backoff with a cap, and add jitter so many clients do not retry together.

### Persist long-lived retry decisions

A synchronous request should have a small retry budget. Durable background jobs can persist attempts and next-attempt time, then release the worker rather than sleeping for minutes. Record a terminal/dead state after the budget is exhausted and provide a deliberate operator replay procedure.

### Build it yourself

1. Build a fake destination that responds 503 twice, then 204.
2. Add retry classification, a capped delay function, and an overall deadline.
3. Test a permanent 400 and a 429 with Retry-After.
4. Use the same operation ID across attempts and design receiver-side deduplication.

<details><summary>Reference / excerpt</summary>

```
delay := min(baseDelay * (1 << attempt), maxDelay)
// Add jitter within a documented bound.
// Wait with a timer/select that also observes ctx.Done().
// Retry only when the operation and response policy permit it.
```

</details>

### Checkpoints

- Cancellation stops waiting between attempts.
- The retry budget limits load and completion time.
- The design acknowledges at-least-once delivery and possible duplicates.

**Watch for:** Retrying at several layers can multiply attempts exponentially. Decide which layer owns retries and share the deadline budget.

**Recall:** Why does a stable request ID not guarantee exactly-once effects by itself?

The receiver must atomically associate that ID with the effect and its result. A header alone cannot prevent a duplicate after a network ambiguity.


Further reading: [HTTP semantics: idempotency](https://www.rfc-editor.org/rfc/rfc9110.html#section-9.2.2)


## 24. Project: build a concurrent site checker

Operations needs a bounded, cancellable tool to check a list of service endpoints during a release.

### Define behavior before parallelizing

The tool accepts URLs, worker count, and timeout. It emits a result for every attempted URL and exits nonzero if any fail or the operation is cancelled. Completion order may differ from input order; decide whether to sort the final report or stream results with an explicit index.

### Draw the lifecycle

A producer sends jobs. A fixed pool performs requests. One collector consumes every result. The producer closes jobs; a coordinator closes results after workers stop. The supplied implementation keeps collecting until workers complete, which is why its result sends can finish. If you change that behavior, add cancellation to the send path too.

### Exercise failures rather than only speed

Prove that concurrency never exceeds the configured maximum. Test local fake servers rather than relying on public websites. Include one cancelled request, one invalid URL, one 500, and one slow response. Measure the sequential and parallel cases and explain when extra workers stop helping.

### Build it yourself

1. Rebuild the reference tool with only one worker first.
2. Add bounded concurrency and preserve the output contract.
3. Add -json and file input using the CLI techniques from stage 2.
4. Write a short design note showing ownership and shutdown paths for every goroutine.

```sh
go test -race ./examples/sitecheck
go build -o bin/sitecheck ./examples/sitecheck
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestCheckBoundsAndFailures(t *testing.T) {
	var active, peak atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := active.Add(1)
		defer active.Add(-1)
		for old := peak.Load(); n > old; old = peak.Load() {
			if peak.CompareAndSwap(old, n) {
				break
			}
		}
		time.Sleep(time.Millisecond)
		if r.URL.Path == "/bad" {
			w.WriteHeader(503)
			return
		}
		w.WriteHeader(204)
	}))
	defer srv.Close()
	urls := []string{srv.URL, srv.URL, srv.URL, srv.URL + "/bad"}
	got := check(context.Background(), srv.Client(), urls, 2)
	if len(got) != 4 || peak.Load() > 2 {
		t.Fatalf("results=%d peak=%d", len(got), peak.Load())
	}
	bad := 0
	for _, r := range got {
		if r.Err != nil {
			t.Fatal(r.Err)
		}
		if r.Status == 503 {
			bad++
		}
	}
	if bad != 1 {
		t.Fatalf("503 count=%d", bad)
	}
}

func TestCancelledRequest(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	got := check(ctx, &http.Client{Timeout: time.Second}, []string{"http://127.0.0.1:1"}, 1)
	for _, r := range got {
		if r.Err == nil {
			t.Fatal("cancelled request succeeded")
		}
	}
}
```

</details>

### Checkpoints

- Tests cover failure responses and cancellation.
- Concurrency is bounded and race-detector checks pass on exercised paths.
- A new teammate can understand why the program cannot wait forever on an unconsumed result.

**Watch for:** A faster happy-path demo is not enough. A concurrency implementation must also have a convincing termination story.

**Recall:** What is your strongest evidence that the tool is correct under load?

Bounded-concurrency assertions, cancellation and error tests, race-detector runs, and a clear ownership/lifecycle argument together provide stronger evidence than timing alone.


Further reading: [Concurrency patterns](https://go.dev/blog/pipelines)


## 25. Organize a business application around rules

An application has HTTP endpoints, scheduled jobs, and a CLI. All must enforce the same business rules.

### Keep business decisions in one place

Transport code parses requests and writes responses. Application logic decides whether an action is allowed and coordinates work. Storage code handles queries. These are responsibilities, not a requirement for a huge folder tree. A modular monolith is often easier to develop, test, and operate than several premature services.

### Make dependencies visible

Construct concrete dependencies in main and pass them into the code that uses them. Define a small interface at a consumer boundary only when behavior needs substitution. Avoid service locators, hidden global databases, and an interface for every struct. Context carries cancellation and request-scoped metadata, not your dependency container.

### Preserve invariants at every entry point

If only HTTP handlers validate a rule, a worker or import command can bypass it. Put domain validation where every use case invokes it, then enforce hard data constraints in the database too. Define expected errors such as conflict or not found and translate them at each transport boundary.

### Build it yourself

1. Draw the dependencies of TeamDesk and identify business decisions hidden in handlers.
2. Extract a CreateTicket use case that receives typed input and returns a ticket or categorized error.
3. Call the same use case from HTTP and a small CLI import command.
4. Introduce only the storage methods that use case needs; test it with a small fake and a real store integration test.

<details><summary>Reference / excerpt</summary>

```
type TicketWriter interface {
    Create(ctx context.Context, title string) (Ticket, error)
}

type TicketService struct { store TicketWriter }
// The HTTP handler and import command both call this service.
```

</details>

### Checkpoints

- Business rules are not duplicated across transport layers.
- Dependencies are assembled explicitly.
- Package imports form an understandable acyclic graph.

**Watch for:** Framework-shaped folders can obscure the domain. Name packages for responsibilities and concepts that the team actually discusses.

**Recall:** What is a useful reason to introduce an interface?

A consumer needs a small behavior that can be supplied by different implementations, such as real storage and a deterministic fake. Interface creation should follow that need.


Further reading: [Organizing a Go module](https://go.dev/doc/modules/layout) · [Go code review comments](https://go.dev/wiki/CodeReviewComments)


## 26. Protect invariants with transactions and migrations

Two buyers try to reserve the final item. A simple read-then-write can sell it twice even when there are no Go data races.

### A business race can live in the database

A mutex in one process cannot protect against another process or an external client. Prefer a single conditional update, a database constraint, or a transaction with a locking/isolation strategy appropriate to the database. Inspect affected-row counts. Learn what your database isolation level guarantees and which operations must retry.

### Use one transaction handle consistently

BeginTx creates a transaction. Execute all related statements through tx, not the outer db, then commit only after every step succeeds. Defer Rollback as cleanup for early exits. Commit itself can fail, so propagate that error. Avoid remote HTTP calls inside a transaction: locks may be held while the network stalls.

### Evolve schemas compatibly

A migration is a versioned change to stored structure and data. Test it on existing data and plan mixed-version deployment behavior. An expand/backfill/contract sequence adds new structure, migrates data, then removes obsolete structure after callers change. A CREATE TABLE IF NOT EXISTS statement is not a complete migration strategy.

### Build it yourself

1. Implement a stock decrement with WHERE available > 0 and check RowsAffected.
2. Test concurrent attempts for one remaining item and prove only one succeeds.
3. Use a transaction to write an event and pending delivery together.
4. Write a migration adding a nullable field, backfill it, and explain when it can become required.

<details><summary>Reference / excerpt</summary>

```
result, err := tx.ExecContext(ctx,
    "UPDATE stock SET available = available - 1 WHERE sku = ? AND available > 0", sku)
// Check err, then RowsAffected: 0 means reservation was not possible.
// Insert the reservation using this same tx, then check tx.Commit().
```

</details>

### Checkpoints

- The invariant holds across multiple connections or processes.
- All operations in the unit of work use tx.
- Migration behavior is tested against an older schema with data.

**Watch for:** SQLite and PostgreSQL have different concurrency and locking behavior. Passing SQLite tests does not prove PostgreSQL isolation behavior.

**Recall:** Why is a Go mutex insufficient to prevent overselling in a multi-replica application?

Each replica has its own memory and lock. The shared database must enforce the invariant across all writers.


Further reading: [Executing transactions](https://go.dev/doc/database/execute-transactions) · [PostgreSQL concurrency control](https://www.postgresql.org/docs/current/mvcc.html)


## 27. Build durable jobs with an outbox and leases

An event is accepted just before the process crashes. The delivery must still be discoverable when another worker starts.

### Commit intent before acknowledging it

In one transaction, save the business event and its pending delivery record. Only after commit may the API acknowledge acceptance. A worker polls or claims durable pending work. The database outbox avoids the gap between committing business state and publishing to another system, but it does not magically remove duplicates.

### A claim needs an expiry and identity

A worker claims a due job atomically and records a lease expiry plus a unique lease token. If it crashes, another worker can recover the job after expiry. When recording success, require the same lease token so an old worker cannot overwrite a newer owner’s state. A single-node demo can use a simple approach; multiple replicas require a database-specific claim algorithm.

### Assume at-least-once delivery

If the destination succeeds but the worker dies before recording success, a later worker may deliver again. Stable delivery IDs and receiver-side idempotency are part of the contract. Persist attempt count, next attempt, result category, and a dead-letter state. Manual replay should be auditable and should preserve the intended deduplication semantics.

### Build it yourself

1. Use the supplied schema to record an event and delivery in one transaction.
2. Implement an atomic claim with a lease token and expiry.
3. Kill the worker after claiming a job and show that a later worker can reclaim it.
4. Simulate success followed by a crash before the local acknowledgement; demonstrate why the receiver needs deduplication.

<details><summary>Reference / excerpt</summary>

```
-- Learning schema for a single-node durable webhook service using SQLite.
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS events (
    tenant_id TEXT NOT NULL,
    id TEXT NOT NULL,
    idempotency_key TEXT NOT NULL,
    body_hash TEXT NOT NULL,
    payload BLOB NOT NULL,
    created_at INTEGER NOT NULL,
    PRIMARY KEY (tenant_id, id),
    UNIQUE (tenant_id, idempotency_key)
);
CREATE TABLE IF NOT EXISTS deliveries (
    tenant_id TEXT NOT NULL,
    id TEXT NOT NULL,
    event_id TEXT NOT NULL,
    destination_id TEXT NOT NULL,
    state TEXT NOT NULL CHECK (state IN ('pending','leased','sent','dead')),
    attempts INTEGER NOT NULL DEFAULT 0,
    next_attempt INTEGER NOT NULL,
    lease_until INTEGER,
    lease_token TEXT,
    PRIMARY KEY (tenant_id, id),
    FOREIGN KEY (tenant_id, event_id) REFERENCES events(tenant_id, id)
);
CREATE INDEX IF NOT EXISTS deliveries_ready ON deliveries(state, next_attempt);
```

</details>

### Checkpoints

- An accepted event has durable pending work.
- An expired worker cannot complete another worker’s lease.
- Retry timing survives process restarts.
- The design explicitly documents duplicate delivery.

**Watch for:** A channel is an in-memory queue, not durable storage. It can coordinate workers after jobs are persisted but cannot replace the persistence guarantee.

**Recall:** Why store both lease expiry and lease token?

Expiry enables recovery after a crash. The token identifies the particular claim so a stale worker cannot incorrectly update a later claim.


Further reading: [Transactions](https://go.dev/doc/database/execute-transactions)


## 28. Make authorization and tenant isolation structural

Two organizations share your service. A guessed object ID, cached response, or background job must not cross the boundary.

### Resolve identity from a trusted boundary

Verify the session or token, derive the actor and tenant, then pass an explicit identity into the application operation. Never trust an X-Tenant-ID header or request body field without checking that the authenticated actor is allowed to select it. Separate human permissions from worker service identity.

### Carry the tenant into every data path

Scope reads, writes, unique constraints, idempotency keys, caches, and queued jobs by tenant. A function named Find(ctx, id) can make accidental unscoped queries easy; FindForTenant(ctx, tenant, id) makes the policy visible. Database row-level security can add protection, but it requires careful connection/session handling and tests.

### Test hostile combinations

Use at least two tenants and two users in integration tests. Try reading, mutating, searching, exporting, and replaying each other’s records. Verify logs and error messages do not leak sensitive fields. Audit security-relevant actions with actor, tenant, action, target, outcome, and time while excluding secrets.

### Build it yourself

1. Add tenant_id to the event and delivery schema and to composite constraints.
2. Derive the tenant from verified identity rather than arbitrary client input.
3. Write tests swapping tenant and object IDs across read, update, list, and replay endpoints.
4. Inspect cache keys and worker payloads for the same isolation rule.

<details><summary>Reference / excerpt</summary>

```
SELECT id, payload FROM events
WHERE tenant_id = ? AND id = ?;

-- Scope idempotency uniqueness by tenant too.
UNIQUE (tenant_id, idempotency_key)
```

</details>

### Checkpoints

- All resource access is tenant-scoped.
- Two tenants can use the same idempotency key without colliding.
- An unauthorized actor cannot infer private object contents through error messages.

**Watch for:** Adding tenant_id to a table does not enforce isolation unless every access path or a database policy actually applies it.

**Recall:** Why must cache keys include the tenant?

The same object identifier or query may exist in multiple tenants. A shared cache entry can leak one tenant’s data into another’s response.


Further reading: [OWASP authorization guidance](https://cheatsheetseries.owasp.org/cheatsheets/Authorization_Cheat_Sheet.html)


## 29. Version APIs, events, and integration boundaries

Another team upgrades on a different schedule. Your service must evolve without breaking its clients.

### Compatibility is a product decision

A public contract includes fields, types, statuses, error codes, ordering, pagination, and delivery semantics. Prefer additive changes when possible, document defaults, and deprecate before removing. Returning a new field is often safe, but strict decoders may reject it; validate compatibility with actual client behavior rather than assumptions.

### Choose the protocol for the requirement

JSON over HTTP is easy to inspect and use from browsers and scripts. gRPC offers generated contracts and streaming where those features justify the tooling and operational requirements. Protobuf evolution requires preserving field numbers and reserving removed ones. Messaging adds asynchronous delivery, schema evolution, deduplication, and consumer-lag concerns.

### Keep distributed failures visible

Propagate deadlines and trace/request identifiers across calls. Bound retries and avoid multiplying them at every layer. A transaction spanning independent services is a coordination problem; an outbox or workflow with compensating actions may fit better than pretending all services commit atomically. Start with one process until independent deployment or scaling is needed.

### Build it yourself

1. Write a small OpenAPI document for event creation and delivery lookup, including errors.
2. Add a response field and test it with an older client fixture.
3. Design an event envelope with event ID, type, schema version, tenant, and occurrence time.
4. As an extension, expose one read-only lookup through gRPC and compare client complexity and deadline behavior.

<details><summary>Reference / excerpt</summary>

```
{
  "id": "evt_123",
  "type": "invoice.created",
  "schema_version": 1,
  "tenant_id": "tenant_a",
  "occurred_at": "2026-01-01T12:00:00Z",
  "data": {"invoice_id": "inv_42"}
}
```

</details>

### Checkpoints

- An older supported client survives an additive server change.
- The API contract defines pagination and failure responses.
- Event consumers have a documented deduplication key and compatibility policy.

**Watch for:** Microservices do not remove coupling; they can move it into networks, deployment schedules, and data contracts.

**Recall:** When is gRPC a reasonable addition?

When generated typed clients, efficient internal communication, or streaming solve a concrete requirement and the team can operate that protocol well.


Further reading: [gRPC Go tutorial](https://grpc.io/docs/languages/go/basics/) · [Protobuf updating messages](https://protobuf.dev/programming-guides/proto3/#updating)


## 30. Project: build Dispatch, a webhook delivery service

Implement an application that accepts business events and reliably delivers them to configured partner destinations.

### An explicit acceptance contract

POST /events accepts a topic, payload, and idempotency key. After atomically saving the event and pending delivery, return 202 with an event ID. The same key and payload should return the original result; the same key with a different payload should conflict. GET /deliveries/{id} reports state for the authorized tenant.

### A deliberately bounded implementation

Start with one process and SQLite, fixed operator-configured destinations, a small worker count, and a documented retry budget. Use the schema as a scaffold, not a completed service. Add leases and recovery before running several workers or replicas. Configure destinations outside untrusted requests to keep the first version’s network boundary understandable.

### Professional evidence

Demonstrate survival of a crash after acceptance, duplicate-request handling, tenant isolation, bounded worker concurrency, destination timeouts, exhausted retries, and controlled shutdown. Document remaining limitations. The important artifact is the combination of working software, tests, and an explanation of what it guarantees—and what it does not.

### Build it yourself

1. Implement HTTP acceptance, durable storage, and delivery lookup before adding workers.
2. Add a configured local fake destination, fixed-size workers, and persisted attempt state.
3. Add idempotency, lease recovery, tenant checks, and an operator replay action.
4. Execute every failure scenario in the project brief and prepare a short operational runbook.

```sh
go test -race ./...
# Implement Dispatch in your own module; see projects/DISPATCH.md.
```


<details><summary>Reference / excerpt</summary>

```
-- Learning schema for a single-node durable webhook service using SQLite.
PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS events (
    tenant_id TEXT NOT NULL,
    id TEXT NOT NULL,
    idempotency_key TEXT NOT NULL,
    body_hash TEXT NOT NULL,
    payload BLOB NOT NULL,
    created_at INTEGER NOT NULL,
    PRIMARY KEY (tenant_id, id),
    UNIQUE (tenant_id, idempotency_key)
);
CREATE TABLE IF NOT EXISTS deliveries (
    tenant_id TEXT NOT NULL,
    id TEXT NOT NULL,
    event_id TEXT NOT NULL,
    destination_id TEXT NOT NULL,
    state TEXT NOT NULL CHECK (state IN ('pending','leased','sent','dead')),
    attempts INTEGER NOT NULL DEFAULT 0,
    next_attempt INTEGER NOT NULL,
    lease_until INTEGER,
    lease_token TEXT,
    PRIMARY KEY (tenant_id, id),
    FOREIGN KEY (tenant_id, event_id) REFERENCES events(tenant_id, id)
);
CREATE INDEX IF NOT EXISTS deliveries_ready ON deliveries(state, next_attempt);
```

</details>

### Checkpoints

- No accepted event is silently lost on process restart.
- The receiver can detect duplicate deliveries.
- Tenant isolation and idempotency conflicts are tested.
- A second developer can deploy locally, observe failures, and recover pending work.

**Watch for:** Do not call a learning implementation production-ready merely because it has retries. Review its security, resource limits, recovery, and operational assumptions first.

**Recall:** What distinguishes this capstone from basic CRUD?

It protects business invariants across crashes, retries, concurrent workers, external failures, and tenant boundaries, with tests and operational procedures for those situations.


Further reading: [Go database guide](https://go.dev/doc/database/) · [Context](https://pkg.go.dev/context)


## 31. Build a testing strategy and a useful CI pipeline

A change passes unit tests but breaks real database behavior. Another accepts input that nobody thought to test.

### Test at several boundaries

Unit tests exercise small decisions. Integration tests cover real database queries, migrations, and HTTP boundaries. End-to-end checks cover a few important user journeys. Use real temporary resources where mocking would hide the behavior you need to verify. Keep tests isolated, deterministic, and independent of external websites.

### Use Go’s complementary tools

The race detector detects exercised unsynchronized memory access. Fuzzing explores input space and preserves failures as regression inputs. Benchmarks measure costs rather than correctness. Static analysis and dependency scanning find another set of issues. None of these alone establishes that business requirements are satisfied.

### Keep CI reproducible

Pin the Go toolchain policy and dependency versions through go.mod/go.sum, build from a clean checkout, run tests and static checks, and produce artifacts from that exact commit. Run a quick suite on each change and deeper integration or load checks where warranted. A flaky test should be investigated rather than repeatedly rerun until green.

### Build it yourself

1. Run the provided unit and integration tests, including race checks.
2. Fuzz LogLens and add any discovered failure to its regression corpus.
3. Add a migration test starting from an old populated schema.
4. Inspect the included CI workflow and add the checks required by your capstone.

```sh
go test ./...
go test -race ./...
go vet ./...
go test ./examples/loglens -fuzz=FuzzScan -fuzztime=10s
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	cases := []struct {
		name, input, needle string
		want                Summary
	}{
		{"empty", "", "ERROR", Summary{}},
		{"mixed", "INFO start\nERROR failed\n", "ERROR", Summary{2, 1}},
		{"no newline", "ERROR failed", "ERROR", Summary{1, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := scan(strings.NewReader(tc.input), tc.needle)
			if err != nil || got != tc.want {
				t.Fatalf("got %+v, %v; want %+v", got, err, tc.want)
			}
		})
	}
}

func TestRunJSONAndUsage(t *testing.T) {
	var out, stderr bytes.Buffer
	code := run([]string{"-json"}, strings.NewReader("ERROR bad\n"), &out, &stderr)
	if code != 0 || out.String() != "{\"lines\":1,\"matches\":1}\n" {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), stderr.String())
	}
	if got := run([]string{"-unknown"}, strings.NewReader(""), &out, &stderr); got != 2 {
		t.Fatalf("usage code=%d", got)
	}
}

func FuzzScan(f *testing.F) {
	f.Add("INFO ok\nERROR failed\n", "ERROR")
	f.Fuzz(func(t *testing.T, input, needle string) {
		got, err := scan(strings.NewReader(input), needle)
		if err == nil && (got.Matches < 0 || got.Matches > got.Lines) {
			t.Fatalf("invalid summary: %+v", got)
		}
	})
}

func BenchmarkScan(b *testing.B) {
	input := strings.Repeat("INFO request complete\nERROR request failed\n", 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := scan(strings.NewReader(input), "ERROR"); err != nil {
			b.Fatal(err)
		}
	}
}
```

</details>

### Checkpoints

- Tests do not depend on developer machine state.
- HTTP input, database behavior, and key business invariants are exercised.
- CI can reproduce the build from a fresh checkout.

**Watch for:** 100% coverage can coexist with weak assertions. Review what failures your tests would actually detect.

**Recall:** Why test a real database even when the repository interface has a fake?

A fake cannot establish that SQL, constraints, transactions, migrations, and driver behavior are correct.


Further reading: [Go fuzzing](https://go.dev/doc/security/fuzz/) · [Testing documentation](https://pkg.go.dev/testing)


## 32. Measure allocations, latency, and bottlenecks

Users report slow exports. You need evidence before changing data structures or adding concurrency.

### Start with a representative question

Define what is slow: median or tail latency, throughput, peak memory, CPU, or allocation rate. Use input sizes and shapes resembling real work. A microbenchmark isolates one operation; a load test exposes queueing and contention across the application. Neither substitutes for a meaningful workload.

### Know the major costs

Slices can retain large backing arrays, string concatenation can allocate, and unnecessary copies can add work. Goroutines consume resources and may wait on locks or I/O. Garbage collection costs depend partly on allocation behavior and live memory. These are hypotheses to measure, not reasons to preemptively pool everything.

### Profile, change, compare

Use go test -bench with -benchmem to record allocation counts. CPU profiles show where CPU is spent; heap profiles show allocations or retained memory; blocking and mutex profiles help explain waiting. Change one thing, rerun equivalent measurements, and keep correctness tests. Protect diagnostic endpoints from untrusted access.

### Build it yourself

1. Run BenchmarkScan with allocations and save the baseline.
2. Capture a CPU profile for the benchmark and inspect it with go tool pprof.
3. Try one improvement, such as reducing an unnecessary allocation, then compare repeated runs.
4. Write a short result that includes workload, before/after data, and trade-offs.

```sh
go test ./examples/loglens -run=^$ -bench=BenchmarkScan -benchmem
go test ./examples/loglens -run=^$ -bench=BenchmarkScan -cpuprofile=/tmp/loglens.cpu
go tool pprof /tmp/loglens.cpu
```


<details><summary>Reference / excerpt</summary>

```
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	cases := []struct {
		name, input, needle string
		want                Summary
	}{
		{"empty", "", "ERROR", Summary{}},
		{"mixed", "INFO start\nERROR failed\n", "ERROR", Summary{2, 1}},
		{"no newline", "ERROR failed", "ERROR", Summary{1, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := scan(strings.NewReader(tc.input), tc.needle)
			if err != nil || got != tc.want {
				t.Fatalf("got %+v, %v; want %+v", got, err, tc.want)
			}
		})
	}
}

func TestRunJSONAndUsage(t *testing.T) {
	var out, stderr bytes.Buffer
	code := run([]string{"-json"}, strings.NewReader("ERROR bad\n"), &out, &stderr)
	if code != 0 || out.String() != "{\"lines\":1,\"matches\":1}\n" {
		t.Fatalf("code=%d out=%q err=%q", code, out.String(), stderr.String())
	}
	if got := run([]string{"-unknown"}, strings.NewReader(""), &out, &stderr); got != 2 {
		t.Fatalf("usage code=%d", got)
	}
}

func FuzzScan(f *testing.F) {
	f.Add("INFO ok\nERROR failed\n", "ERROR")
	f.Fuzz(func(t *testing.T, input, needle string) {
		got, err := scan(strings.NewReader(input), needle)
		if err == nil && (got.Matches < 0 || got.Matches > got.Lines) {
			t.Fatalf("invalid summary: %+v", got)
		}
	})
}

func BenchmarkScan(b *testing.B) {
	input := strings.Repeat("INFO request complete\nERROR request failed\n", 100)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := scan(strings.NewReader(input), "ERROR"); err != nil {
			b.Fatal(err)
		}
	}
}
```

</details>

### Checkpoints

- An optimization is justified by measured evidence.
- Correctness and resource limits remain intact.
- You distinguish CPU time from time blocked on I/O.

**Watch for:** A faster microbenchmark can make the application worse if it increases memory retention or contention. Measure the relevant workload too.

**Recall:** Why can retaining a small subslice retain a large allocation?

The subslice still refers to the original backing array, so that array can remain live until no references to it remain.


Further reading: [Go diagnostics](https://go.dev/doc/diagnostics) · [Profiling Go programs](https://go.dev/blog/pprof)


## 33. Make failures explainable

A delivery is late. Operations needs to distinguish backlog, a slow database, and a failing destination.

### Use logs for events with context

log/slog produces structured key-value records. Include operation, event or delivery ID, result category, and duration where useful. Request IDs help follow a request; tracing can connect work across services. Do not log raw credentials, session tokens, or sensitive payloads by default.

### Use metrics for trends and limits

Measure request count, error rate, latency, in-flight work, queue age, retries, and exhausted jobs. Choose label dimensions with bounded cardinality. Putting unique request IDs into metric labels can overwhelm a metrics system; keep those identifiers in logs or traces instead.

### Turn observations into actions

Define what healthy means for the user: accepted events delivered within an agreed time, for example. Alert on sustained user-impacting symptoms, and link alerts to runbooks. Liveness asks whether the process should be restarted; readiness asks whether it should receive new work. A database outage does not always justify restarting every process.

### Build it yourself

1. Add structured logs around API acceptance and delivery attempts.
2. Record queue depth and oldest pending job age; explain how they differ.
3. Add readiness that becomes false during draining.
4. Write a runbook for rising retry count and one for a growing queue.

<details><summary>Reference / excerpt</summary>

```
slog.InfoContext(ctx, "delivery finished",
    "delivery_id", deliveryID,
    "attempt", attempt,
    "result", "success",
    "duration", time.Since(start))
```

</details>

### Checkpoints

- One event can be traced across acceptance, attempts, and completion.
- Metrics avoid unbounded identifiers as labels.
- An operator can follow a written recovery procedure.

**Watch for:** More logging is not automatically better. High-volume payload logs can add cost, expose data, and hide the useful events.

**Recall:** Why monitor queue age as well as queue length?

A small queue can still contain a job stuck for a long time. Age exposes delayed work that a count alone can miss.


Further reading: [Structured logging](https://pkg.go.dev/log/slog) · [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)


## 34. Build, release, and operate a service

A teammate must deploy the same binary you tested and recover the application if a release fails.

### Separate build-time and runtime concerns

Build artifacts from a known commit with a documented toolchain. Runtime configuration supplies addresses, credentials, database location, and limits. Validate it before accepting traffic. A container packages an artifact and its runtime needs; it does not eliminate operating-system, networking, or security responsibilities.

### Release with state in mind

Run migrations through a controlled process before code relies on new schema. Use compatible changes so old and new versions can overlap during rollout. Store durable data outside an ephemeral container filesystem. Maintain backups and test restoring them; an untested backup is only an assumption.

### Plan the exit and rollback

A deployment platform sends signals and may kill the process after a grace period. On Unix deployments handle SIGTERM as well as interactive interrupts. Stop intake, drain with a deadline, then release resources. Application rollback is not always database rollback; keep a forward-recovery plan when data changes cannot be reversed safely.

### Build it yourself

1. Build a tagged artifact and record the commit and Go version.
2. Run TeamDesk from a different directory with a persistent database path.
3. Add SIGTERM handling for a Unix deployment and test its drain deadline.
4. Write backup, restore, rollout, and rollback procedures; perform a restore into a fresh environment.

```sh
go build -trimpath -o bin/teamdesk ./examples/teamdesk
./bin/teamdesk -addr 127.0.0.1:8080 -db /tmp/teamdesk-release.db
```

### Checkpoints

- The deployed artifact corresponds to the tested commit.
- Configuration and data outlive individual containers.
- Shutdown fits inside the platform’s termination grace period.
- A restore has actually been exercised.

**Watch for:** Using a scratch container requires understanding runtime needs such as CA certificates, time zones, and any dynamically linked libraries.

**Recall:** Why can reverting code be insufficient after a bad deployment?

The release may have changed schema or stored data. Older code may no longer understand that state, so recovery needs a data-aware plan.


Further reading: [Building Go packages](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies) · [Signal handling](https://pkg.go.dev/os/signal)


## 35. Review dependencies, changes, and security

A service will outlive its first developer. Future changes need to be understandable and safe to review.

### Treat dependencies as ongoing commitments

Review what a library adds, its maintenance, licensing, and security history. Pin reproducible versions and understand transitive dependencies. go.sum verifies downloaded module content; it is not a declaration that dependencies are vulnerability-free. Use govulncheck and supported toolchains as part of maintenance.

### Write changes a reviewer can assess

A good change explains the concrete behavior before and after, why it matters, and how it was verified. Keep changes focused, name functions and packages by their role, and document exported APIs where their contract is not obvious. Prefer a short design note for a meaningful trade-off over a large amount of ceremonial architecture.

### Practice incident-driven learning

When a failure occurs, preserve evidence, restore service, identify contributing conditions, and add a regression check or operational improvement. Review interfaces and invariants when requirements change. Avoid reflexive rewrites: small well-observed changes are easier to evaluate and roll back.

### Build it yourself

1. Run go mod tidy and inspect the resulting dependency diff.
2. Scan the module with govulncheck and investigate reachable findings rather than only counting them.
3. Prepare a small pull request adding one feature to your capstone with tests and a behavior-focused description.
4. Ask a reviewer to explain your design back to you and note where it was unclear.

```sh
go mod tidy
go vet ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

### Checkpoints

- Dependencies have a documented update and scanning process.
- A pull request communicates the behavior and validation.
- A regression test or operational check follows a real incident.

**Watch for:** A scanner cannot replace authorization review, safe secret handling, or business-logic analysis. It covers a different category of risks.

**Recall:** What does go.sum establish?

It records expected module-content hashes for reproducibility and integrity checks. It does not prove that the code is safe or correct.


Further reading: [Security best practices](https://go.dev/doc/security/best-practices) · [Vulnerability checking](https://go.dev/doc/tutorial/govulncheck)


## 36. Demonstrate independent professional delivery

Take an unfamiliar requirement and deliver a small, complete Go feature with convincing evidence.

### Use outcomes rather than completion badges

Reading a guide is exposure; implementing a variant without the solution demonstrates understanding. Professional skill grows through repeated design, debugging, code review, and operation of real systems. Completing this path gives you a foundation and a portfolio, not a guarantee of expertise or job readiness.

### Choose a fresh requirement

For the CLI, add a new report format without breaking scripts. For TeamDesk, add a tenant-scoped export with cancellation. For Dispatch, add a replay workflow with authorization and idempotency. Write acceptance cases first, implement the smallest complete version, and inspect its failure paths.

### Defend the trade-offs

Explain your API, package boundaries, data ownership, transaction boundaries, concurrency limits, security assumptions, tests, and operational signals. Use profiles or measurements for performance claims. Be able to name the design’s limits and the next change you would make under a different workload.

### Build it yourself

1. Pick one unseen requirement and write a one-page implementation plan.
2. Implement it without copying a chapter’s solution; consult documentation as needed.
3. Have another person review the code and run it using only your README.
4. Record the feedback, improve the implementation, and write a short incident or failure drill report.
### Checkpoints

- You can build and test a CLI independently.
- You can deliver an HTML web app and JSON API with persistent, authorized data access.
- You can explain cancellation, retries, transaction boundaries, and duplicate handling.
- You can profile, deploy, diagnose, and safely modify the system.

**Watch for:** Checking a lesson as done is a study aid, not an assessment. Revisit a topic whenever you cannot explain or adapt it without the reference.

**Recall:** What is a better measure of progress than how many lessons you read?

The number of unfamiliar requirements you can turn into working, tested, understandable software—and the quality of your decisions when it fails.


Further reading: [Go learning resources](https://go.dev/learn/) · [Go code review comments](https://go.dev/wiki/CodeReviewComments)

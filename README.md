# Go From Fundamentals (For Python Developers)

A simple, step-by-step, hands-on handbook for learning Go from core programming fundamentals to production-ready backend development. Every lesson contrasts Go with Python, highlights Go's golden rules, links to a runnable program in [`lessons/`](./lessons), and includes a quick self-check with a collapsible answer.

---

## Table of Contents

1. [Lesson 1: How Go Works & First Program](#lesson-1-how-go-works--first-program) (`lessons/01-hello`)
2. [Lesson 2: Variables, Types & Zero Values](#lesson-2-variables-types--zero-values) (`lessons/02-variables`)
3. [Lesson 3: Making Decisions (`if / else if / else`)](#lesson-3-making-decisions-if--else-if--else) (`lessons/03-decisions`)
4. [Lesson 4: Loops (`for` is All You Need)](#lesson-4-loops-for-is-all-you-need) (`lessons/04-loops`)
5. [Lesson 5: Lists & Collections (Slices and `range`)](#lesson-5-lists--collections-slices-and-range) (`lessons/05-slices`)
6. [Lesson 6: Maps (Key-Value Dictionaries)](#lesson-6-maps-key-value-dictionaries) (`lessons/06-maps`)
7. [Lesson 7: Functions (`func`) & Error Handling](#lesson-7-functions-func--error-handling) (`lessons/07-functions`)
8. [Lesson 8: Custom Data Types (`struct`) & Methods](#lesson-8-custom-data-types-struct--methods) (`lessons/08-structs`)
9. [Lesson 9: Interfaces (Type-Safe Duck Typing)](#lesson-9-interfaces-type-safe-duck-typing) (`lessons/09-interfaces`)
10. [Lesson 10: Concurrency (`go` & Channels)](#lesson-10-concurrency-go--channels) (`lessons/10-concurrency`)
11. [Lesson 11: Mini Project — Task Tracker](#lesson-11-mini-project--task-tracker) (`lessons/11-mini-project`)
12. [Lesson 12: Pointers & Memory](#lesson-12-pointers--memory) (`lessons/12-pointers`)
13. [Lesson 13: Defer, Panic & Recover](#lesson-13-defer-panic--recover) (`lessons/13-defer-panic-recover`)
14. [Lesson 14: Files, I/O & Working with the OS](#lesson-14-files-io--working-with-the-os) (`lessons/14-files-and-os`)
15. [Lesson 15: Working with JSON](#lesson-15-working-with-json) (`lessons/15-json`)
16. [Lesson 16: Building Web Servers & HTTP Clients](#lesson-16-building-web-servers--http-clients) (`lessons/16-http-server-and-client`)
17. [Lesson 17: Automated Testing & Benchmarks](#lesson-17-automated-testing--benchmarks) (`lessons/17-testing`)
18. [Lesson 18: Context & Timeouts](#lesson-18-context--timeouts) (`lessons/18-context-and-timeouts`)
19. [Lesson 19: Generics in Modern Go](#lesson-19-generics-in-modern-go) (`lessons/19-generics`)
20. [Lesson 20: Capstone Project — Production-Ready Microservice](#lesson-20-capstone-project--production-ready-microservice) (`lessons/20-capstone-service`)
21. [Quick Cheat Sheet: Python vs. Go](#quick-cheat-sheet-python-vs-go)
22. [Additional Reference Examples](#additional-reference-examples)

---

## Lesson 1: How Go Works & First Program

* **Runnable code:** [`lessons/01-hello/main.go`](./lessons/01-hello/main.go)
* **Run it:** `go run ./lessons/01-hello`

### Python vs. Go: How Code Runs
* **Python (Interpreted):** Reads and executes your `.py` file line-by-line at runtime.
* **Go (Compiled):** Translates your code directly into a single, super-fast machine-code binary before running. `go run` compiles and runs in one step!

### Anatomy of Every Go Program
```go
package main // 1. Tells Go this file builds an executable program (not a library)

import "fmt" // 2. Imports the standard "format" package for printing

func main() { // 3. The entry point where execution always starts
	fmt.Println("Hello, my name is Ramu")
	fmt.Println("I am learning Go!")
}
```

> **Golden Rule — Exported Names (Capitalization):**
> Why is it `fmt.Println` with a capital **`P`**?
> In Go, **capitalizing the first letter makes a function or variable public (exported)** outside its package. Lowercase means private!

<details>
<summary><strong>Quick Checkpoint:</strong> What happens if you name the package <code>package helper</code> instead of <code>package main</code> when trying to run a standalone program?</summary>

Go will refuse to run it as an executable! Only `package main` with a `func main()` can be executed directly as a program.
</details>

---

## Lesson 2: Variables, Types & Zero Values

* **Runnable code:** [`lessons/02-variables/main.go`](./lessons/02-variables/main.go)
* **Run it:** `go run ./lessons/02-variables`

### Dynamic vs. Static Typing
* **Python (Dynamic):** A variable can change types anytime (`x = 10` then `x = "hello"`).
* **Go (Static):** Every variable has a fixed type known at compile time. Once `age` is an `int`, it can only ever hold an `int`.

### Basic Types
| Type | Meaning | Example | Zero Value (Default) |
| :--- | :--- | :--- | :--- |
| `string` | Text (always double quotes `""`) | `"Ramu"` | `""` (empty string) |
| `int` | Whole numbers | `30`, `-5` | `0` |
| `float64` | Decimal numbers | `4.9`, `3.14` | `0.0` |
| `bool` | True or False | `true`, `false` | `false` |

### The 3 Ways to Create Variables
```go
// 1. Explicit type (use when declaring without an initial value, or at package level)
var greeting string = "Welcome!"

// 2. Inferred type with var
var year = 2026

// 3. Short declaration := (Used 90% of the time inside functions!)
age := 30
```

### `:=` (Create) vs. `=` (Reassign)
```go
age := 30 // ✅ Creates brand-new variable 'age'
age := 32 // ❌ Compiler Error: no new variables on left side of :=
age = 32  // ✅ Updates existing variable 'age'
```
> **Pro-Tip:** `:=` *can* reuse an existing variable **if** at least one *brand-new* variable is next to it:
> `age, name := 32, "Ramu"` (Allowed because `name` is new!)

> **Golden Rule — No Unused Variables:**
> If you declare a variable inside a function and don't read it, **your code will not compile**. This keeps Go code clean of dead variables.

<details>
<summary><strong>Quick Checkpoint:</strong> If you write <code>var count int</code> without assigning a number, what is the value of <code>count</code>?</summary>

`0`! Go never leaves variables uninitialized—every type gets a clean **Zero Value** (`0`, `0.0`, `""`, or `false`).
</details>

---

## Lesson 3: Making Decisions (`if / else if / else`)

* **Runnable code:** [`lessons/03-decisions/main.go`](./lessons/03-decisions/main.go)
* **Run it:** `go run ./lessons/03-decisions`

### Python vs. Go Comparison
| Feature | Python | Go |
| :--- | :--- | :--- |
| **If / Else If / Else** | `if ...:` / `elif ...:` / `else:` | `if ... {` / `} else if ... {` / `} else {` |
| **Parentheses** | Optional | **Omitted** (`if age >= 18 {`) |
| **Curly Braces `{}`** | None (indentation) | **Mandatory** (even for 1 line) |
| **AND / OR / NOT** | `and` / `or` / `not` | `&&` / `\|\|` / `!` |

### Go's Superpower: `if` with a Short Statement
You can create a temporary variable right inside the `if` header. It only exists inside that `if / else` block:
```go
if score := 85; score >= 80 {
	fmt.Println("Grade: B or higher")
}
// 'score' no longer exists out here!
```

> **Golden Rule — Brace Placement:**
> The opening `{` and `} else {` **must** stay on the same line:
> ```go
> if age >= 18 {
>     fmt.Println("Adult")
> } else {
>     fmt.Println("Minor")
> }
> ```

<details>
<summary><strong>Quick Checkpoint:</strong> What does <code>if temp > 25 && !isRaining</code> mean in plain English?</summary>

"If `temp` is greater than 25 **AND** `isRaining` is **NOT** true (`false`)."
</details>

---

## Lesson 4: Loops (`for` is All You Need)

* **Runnable code:** [`lessons/04-loops/main.go`](./lessons/04-loops/main.go)
* **Run it:** `go run ./lessons/04-loops`

Go has **only ONE loop keyword**: `for`. There is no `while` keyword.

### The 3 Flavors of `for`
```go
// 1. Classic 3-part loop (init; condition; post)
for i := 1; i <= 3; i++ {
	fmt.Println(i)
}

// 2. While-style loop (condition only — replaces Python's 'while')
energy := 3
for energy > 0 {
	energy--
}

// 3. Infinite loop (replaces Python's 'while True:')
for {
	break // exits the loop (use 'continue' to skip to next cycle)
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> How do you write a <code>while count > 0:</code> loop in Go?</summary>

```go
for count > 0 {
    fmt.Println(count)
    count--
}
```
</details>

---

## Lesson 5: Lists & Collections (Slices and `range`)

* **Runnable code:** [`lessons/05-slices/main.go`](./lessons/05-slices/main.go)
* **Run it:** `go run ./lessons/05-slices`

* **Arrays (`[3]string`):** Fixed size that never changes. Rarely used directly.
* **Slices (`[]string`):** Dynamic, resizable lists (like Python's `list`). **Used 99% of the time!**

### Creating, Appending, and Looping
```go
// 1. Create a slice (empty brackets [])
fruits := []string{"apple", "banana"}

// 2. Append (returns the updated slice, so reassign it!)
fruits = append(fruits, "cherry")

// 3. Loop with index and value using 'range' (like Python's enumerate)
for index, fruit := range fruits {
	fmt.Println(index, fruit)
}

// 4. Ignore the index using the Blank Identifier '_'
for _, fruit := range fruits {
	fmt.Println(fruit)
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> Why do we write <code>for _, fruit := range fruits</code> instead of <code>for i, fruit := range fruits</code> when we only print <code>fruit</code>?</summary>

Because Go forbids unused variables! If you don't use `i`, the compiler throws `i declared and not used`. Using `_` tells Go to safely ignore the index.
</details>

---

## Lesson 6: Maps (Key-Value Dictionaries)

* **Runnable code:** [`lessons/06-maps/main.go`](./lessons/06-maps/main.go)
* **Run it:** `go run ./lessons/06-maps`

A Go `map[KeyType]ValueType` is like a Python `dict`, with two major differences:
1. **Strictly typed:** All keys share one type, and all values share one type.
2. **No `KeyError` crashes:** Reading a missing key returns the value type's **zero value** (`""` or `0`).

### Map Operations & The "Comma OK" Idiom
```go
phonebook := map[string]string{
	"Alice": "555-0101",
}

// Add or update
phonebook["Ramu"] = "555-0103"

// Delete
delete(phonebook, "Alice")

// "Comma OK" idiom: Check if a key ACTUALLY exists!
num, ok := phonebook["Charlie"]
if ok {
	fmt.Println("Found:", num)
} else {
	fmt.Println("Charlie is not in the map!")
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> In <code>scores := map[string]int{"Luigi": 0}</code>, what do <code>score, ok := scores["Bowser"]</code> return?</summary>

`score = 0` (the zero value for `int`) and `ok = false` (because `"Bowser"` is not in the map). For `"Luigi"`, it returns `score = 0, ok = true`!
</details>

---

## Lesson 7: Functions (`func`) & Error Handling

* **Runnable code:** [`lessons/07-functions/main.go`](./lessons/07-functions/main.go)
* **Run it:** `go run ./lessons/07-functions`

### Multiple Return Values & Explicit Errors
Go has **no `try / except` exceptions**. Instead, functions that can fail return an `error` as their last return value. When everything succeeds, the error is `nil` (Go's `None`).

```go
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil // nil means NO error!
}

func main() {
	res, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", res)
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> What value does <code>err</code> have when a function succeeds without any problems?</summary>

`nil`! That's why `if err != nil` is the universal check for "did an error happen?"
</details>

---

## Lesson 8: Custom Data Types (`struct`) & Methods

* **Runnable code:** [`lessons/08-structs/main.go`](./lessons/08-structs/main.go)
* **Run it:** `go run ./lessons/08-structs`

Go has **no `class` keyword and no inheritance**. Instead:
1. **`struct`** groups data fields together (like a Python `@dataclass`).
2. **Receivers** attach methods to a struct (like Python's `self`).

### Value Receiver vs. Pointer Receiver (`*`)
```go
type User struct {
	Name string
	Age  int
}

// 1. Value Receiver (u User): gets a COPY (Read-only, cannot modify original)
func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s (%d)", u.Name, u.Age)
}

// 2. Pointer Receiver (u *User): gets a REFERENCE (Modifies the real struct!)
func (u *User) HaveBirthday() {
	u.Age++
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> If <code>HaveBirthday()</code> used <code>(u User)</code> WITHOUT the <code>*</code>, what would happen to <code>user1.Age</code> after calling <code>user1.HaveBirthday()</code>?</summary>

`user1.Age` would stay **30**! Without `*`, Go only increments the temporary copy inside the method. Use a pointer receiver `(u *User)` whenever a method needs to mutate fields.
</details>

---

## Lesson 9: Interfaces (Type-Safe Duck Typing)

* **Runnable code:** [`lessons/09-interfaces/main.go`](./lessons/09-interfaces/main.go)
* **Run it:** `go run ./lessons/09-interfaces`

An **interface** defines a set of methods a type must have.
In Go, you never write `class Email implements Notifier`. If your struct has the required methods, it satisfies the interface **automatically**!

```go
type Notifier interface {
	Send(message string) string
}

type EmailNotifier struct{ Address string }

// Because EmailNotifier has Send(string) string, it IS a Notifier!
func (e EmailNotifier) Send(message string) string {
	return fmt.Sprintf("[Email to %s]: %s", e.Address, message)
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> If a struct defines <code>PostMessage(msg string)</code> instead of <code>Send(msg string)</code>, does it satisfy <code>Notifier</code>?</summary>

No! The method name, parameter types, and return types must match the interface definition exactly.
</details>

---

## Lesson 10: Concurrency (`go` & Channels)

* **Runnable code:** [`lessons/10-concurrency/main.go`](./lessons/10-concurrency/main.go)
* **Run it:** `go run ./lessons/10-concurrency`

Concurrency is Go's standout capability:
* **Goroutine (`go someFunc()`):** Starts a lightweight background task (takes only a few kilobytes of memory; you can run thousands at once).
* **Channel (`chan Type`):** A safe pipe for sending messages between goroutines (`ch <- msg` to send, `msg := <-ch` to receive).

```go
results := make(chan string)

// Run concurrently
go func() {
    results <- "Service is UP"
}()

// Receive result
msg := <-results
fmt.Println(msg)
```

<details>
<summary><strong>Quick Checkpoint:</strong> What keyword turns any normal function call into an asynchronous background task in Go?</summary>

The `go` keyword! For example, `go checkServer()`.
</details>

---

## Lesson 11: Mini Project — Task Tracker

* **Runnable code:** [`lessons/11-mini-project/main.go`](./lessons/11-mini-project/main.go)
* **Run it:** `go run ./lessons/11-mini-project`

Combines **Structs**, **Slices**, **Maps**, **Pointer Methods**, and **Error Handling** into a complete in-memory Task Tracker CLI.

---

## Lesson 12: Pointers & Memory

* **Runnable code:** [`lessons/12-pointers/main.go`](./lessons/12-pointers/main.go)
* **Run it:** `go run ./lessons/12-pointers`

### Python vs. Go: Memory Model
* **Python:** Everything is an object, and every variable holds a reference behind the scenes.
* **Go:** Variables hold actual values by default. Go gives you explicit control over whether to pass by **value** (copy) or by **reference** (pointer).

### The Two Operators
1. **`&` (Address-Of):** Finds where a variable lives in memory.
2. **`*` (Dereference):** Follows the pointer to read or write the actual value at that memory address.

```go
score := 42
ptr := &score  // ptr holds memory address: 0x...
*ptr = 100     // Updates score directly to 100!
```

> **Golden Rule — Safe Pointers:**
> Go has pointers, but **NO pointer arithmetic** (`ptr++` is illegal). You get the performance and mutability of pointers without C/C++ memory corruption risks!

<details>
<summary><strong>Quick Checkpoint:</strong> If you pass <code>x</code> into <code>func update(val int)</code> and set <code>val = 50</code>, does the caller's <code>x</code> change?</summary>

No! Go passes arguments by value (copy) by default. To change the caller's variable, pass the pointer <code>&x</code> into <code>func update(val *int)</code> and write <code>*val = 50</code>.
</details>

---

## Lesson 13: Defer, Panic & Recover

* **Runnable code:** [`lessons/13-defer-panic-recover/main.go`](./lessons/13-defer-panic-recover/main.go)
* **Run it:** `go run ./lessons/13-defer-panic-recover`

### Python `with` vs. Go `defer`
In Python, you use context managers:
```python
with open("file.txt") as f:
    process(f)  # f is closed automatically
```
In Go, you use **`defer`** immediately after acquiring a resource:
```go
file, err := os.Open("file.txt")
if err != nil {
    return err
}
defer file.Close() // Guaranteed to run when the surrounding function exits!
```

### LIFO Order (Last-In, First-Out)
Multiple deferred calls run in reverse order (stacked):
```go
defer fmt.Println("1")
defer fmt.Println("2")
// Output at exit: 2 then 1
```

### `panic` and `recover`
* `panic()`: Go's fatal error stop (like an unhandled exception in Python).
* `recover()`: Intercepts an active panic inside a deferred function, preventing the program from crashing.

<details>
<summary><strong>Quick Checkpoint:</strong> Should you use <code>panic</code> for normal errors (like invalid user input)?</summary>

No! Idiomatic Go returns <code>error</code> values for expected failures. <code>panic</code> is reserved only for fatal, unrecoverable situations (e.g. corrupted configuration or missing system dependencies).
</details>

---

## Lesson 14: Files, I/O & Working with the OS

* **Runnable code:** [`lessons/14-files-and-os/main.go`](./lessons/14-files-and-os/main.go)
* **Run it:** `go run ./lessons/14-files-and-os`

### Reading and Writing Files
Go provides fast, direct standard library tools in `os` and `bufio`:

```go
// 1. One-shot write
err := os.WriteFile("notes.txt", []byte("Hello Go!"), 0644)

// 2. Line-by-line reading with bufio.Scanner
file, err := os.Open("notes.txt")
if err != nil {
    return err
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
```

### Environment Variables & Arguments
* Arguments: `os.Args` (slice of CLI arguments, where `os.Args[0]` is the program name).
* Environment: `os.Getenv("PORT")`.

<details>
<summary><strong>Quick Checkpoint:</strong> Why must <code>defer file.Close()</code> come AFTER <code>if err != nil</code> rather than before?</summary>

Because if <code>os.Open()</code> fails, <code>file</code> will be <code>nil</code>. Calling <code>nil.Close()</code> would cause a nil-pointer panic!
</details>

---

## Lesson 15: Working with JSON

* **Runnable code:** [`lessons/15-json/main.go`](./lessons/15-json/main.go)
* **Run it:** `go run ./lessons/15-json`

### Python vs. Go: JSON Mapping
* **Python:** `json.dumps(obj)` and `json.loads(text)`.
* **Go:** `json.Marshal(v)` (Struct -> `[]byte`) and `json.Unmarshal(data, &v)` (`[]byte` -> Struct).

### Struct Tags
Struct tags map Go PascalCase fields to JSON snake_case or camelCase keys:
```go
type ServerConfig struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Secret   string `json:"-"` // Omitted from JSON completely
}
```

> **Golden Rule — Exported Fields in JSON:**
> The `encoding/json` package can **only** read and write fields that start with a capital letter (exported)!

<details>
<summary><strong>Quick Checkpoint:</strong> What struct tag tells Go to exclude a sensitive field (like a password) from JSON serialization?</summary>

<code>`json:"-"`</code>!
</details>

---

## Lesson 16: Building Web Servers & HTTP Clients

* **Runnable code:** [`lessons/16-http-server-and-client/main.go`](./lessons/16-http-server-and-client/main.go)
* **Run it:** `go run ./lessons/16-http-server-and-client`

Unlike Python, where you typically install external frameworks like FastAPI or Flask, Go's standard library `net/http` is production-grade out of the box!

### Creating an HTTP API Endpoint
```go
mux := http.NewServeMux()

mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go!"})
})

http.ListenAndServe(":8080", mux)
```

### Making HTTP Client Requests
```go
resp, err := http.Get("https://api.example.com/status")
if err != nil {
    return err
}
defer resp.Body.Close() // Always close response body!

body, _ := io.ReadAll(resp.Body)
```

<details>
<summary><strong>Quick Checkpoint:</strong> What line must you always include right after successfully executing an <code>http.Get()</code> call?</summary>

<code>defer resp.Body.Close()</code> to prevent leaking network socket connections!
</details>

---

## Lesson 17: Automated Testing & Benchmarks

* **Runnable code:** [`lessons/17-testing/main.go`](./lessons/17-testing/main.go) & [`lessons/17-testing/calc_test.go`](./lessons/17-testing/calc_test.go)
* **Run tests:** `go test -v ./lessons/17-testing`
* **Run benchmarks:** `go test -bench=. ./lessons/17-testing`

Go includes a built-in test runner without needing `pytest` or `unittest`:
* Test files end with `_test.go`.
* Test functions start with `Test...` and accept `t *testing.T`.
* Benchmark functions start with `Benchmark...` and accept `b *testing.B`.

### Idiomatic Go: Table-Driven Testing
```go
func TestCalculateDiscount(t *testing.T) {
    tests := []struct {
        name     string
        price    float64
        discount int
        expected float64
    }{
        {name: "20% off $100", price: 100.0, discount: 20, expected: 80.0},
        {name: "0% off $50", price: 50.0, discount: 0, expected: 50.0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got, _ := CalculateDiscount(tc.price, tc.discount)
            if got != tc.expected {
                t.Errorf("got %.2f, want %.2f", got, tc.expected)
            }
        })
    }
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> What command runs all tests across all packages in a Go project?</summary>

<code>go test ./...</code>!
</details>

---

## Lesson 18: Context & Timeouts

* **Runnable code:** [`lessons/18-context-and-timeouts/main.go`](./lessons/18-context-and-timeouts/main.go)
* **Run it:** `go run ./lessons/18-context-and-timeouts`

In distributed systems and microservices, requests must not hang forever. `context.Context` is Go's standard mechanism for deadlines, timeouts, and cancellations.

```go
// Create a context that automatically expires after 100 milliseconds
ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
defer cancel()

select {
case result := <-doWork():
    fmt.Println("Work finished:", result)
case <-ctx.Done():
    fmt.Println("Aborted due to timeout:", ctx.Err())
}
```

<details>
<summary><strong>Quick Checkpoint:</strong> What channel closes inside a context when its timeout expires or cancel is called?</summary>

<code>ctx.Done()</code>!
</details>

---

## Lesson 19: Generics in Modern Go

* **Runnable code:** [`lessons/19-generics/main.go`](./lessons/19-generics/main.go)
* **Run it:** `go run ./lessons/19-generics`

Python functions can accept any type dynamically. In Go 1.18+, **Generics** let you write type-safe reusable functions and data structures without runtime type assertions:

```go
// [T comparable] allows using == and !=
func Contains[T comparable](slice []T, target T) bool {
    for _, item := range slice {
        if item == target {
            return true
        }
    }
    return false
}

// Works for slices of strings, ints, or custom structs!
Contains([]int{1, 2, 3}, 2)       // true
Contains([]string{"a", "b"}, "c") // false
```

### Constraints: `any`, `comparable`, and Unions
* `any`: accepts absolutely any type.
* `comparable`: types that can be compared with `==` and `!=`.
* `int | float64`: custom union interface for numeric operations.

<details>
<summary><strong>Quick Checkpoint:</strong> What constraint is required if a generic function needs to compare items with <code>==</code> or use them as map keys?</summary>

<code>comparable</code>!
</details>

---

## Lesson 20: Capstone Project — Production-Ready Microservice

* **Runnable code:** [`lessons/20-capstone-service/main.go`](./lessons/20-capstone-service/main.go)
* **Run it:** `go run ./lessons/20-capstone-service`

Ties all 20 lessons together into a complete, clean, production-grade Go microservice for managing home-lab devices:
1. **Model:** `Device` struct with JSON tags (Lessons 8 & 15).
2. **Interface:** `DeviceStore` contract (Lesson 9).
3. **Thread-Safe Store:** `InMemoryStore` using `sync.RWMutex`, slices, maps, and pointers (Lessons 5, 6, 8, 10, 12, 13).
4. **REST Endpoints:** `GET /health`, `GET /api/devices`, `POST /api/devices` (Lesson 16).
5. **Robust Concurrency:** Background HTTP server with automated client requests (Lesson 10).
6. **Graceful Shutdown:** Context timeout shutdown with zero dropped connections (Lesson 18).

---

## Quick Cheat Sheet: Python vs. Go

| Concept | Python | Go |
| :--- | :--- | :--- |
| **Create variable** | `x = 10` | `x := 10` (or `var x int = 10`) |
| **Update variable** | `x = 20` | `x = 20` |
| **None / Null** | `None` | `nil` |
| **Boolean values** | `True`, `False` | `true`, `false` |
| **Logical operators** | `and`, `or`, `not` | `&&`, `\|\|`, `!` |
| **List / Slice** | `items = ["a", "b"]` | `items := []string{"a", "b"}` |
| **Append to list** | `items.append("c")` | `items = append(items, "c")` |
| **Dictionary / Map** | `d = {"a": 1}` | `d := map[string]int{"a": 1}` |
| **Check key in map** | `if "a" in d:` | `if val, ok := d["a"]; ok {` |
| **For-each loop** | `for i, v in enumerate(x):` | `for i, v := range x {` |
| **While loop** | `while cond:` | `for cond {` |
| **Error handling** | `try: ... except:` | `res, err := fn(); if err != nil {` |
| **Class / Struct** | `class User:` | `type User struct { ... }` |
| **Pass by pointer** | Handled implicitly | `&val` (address) / `*val` (dereference) |
| **Resource cleanup** | `with open(...) as f:` | `f, _ := os.Open(...); defer f.Close()` |
| **JSON encode** | `json.dumps(obj)` | `json.Marshal(struct)` |
| **JSON decode** | `json.loads(text)` | `json.Unmarshal(bytes, &struct)` |
| **HTTP server** | FastAPI / Flask / Django | Standard library `net/http` |
| **Unit tests** | `pytest` / `unittest` | Standard library `testing` (`go test ./...`) |
| **Async / Threading** | `asyncio` / `threading` | Goroutines (`go worker()`) & Channels |
| **Timeouts / Cancellation** | `asyncio.wait_for` | `context.WithTimeout(ctx, duration)` |
| **Generics** | Dynamic (any type) | `func Process[T any](item T)` |

---

## Additional Reference Examples

When you want to inspect larger reference applications and test suites in [`examples/`](./examples):

```sh
# Run all unit tests across the repository
go test -race ./...

# LogLens CLI log analyzer
go run ./examples/loglens -json examples/loglens/sample.log

# Sitecheck concurrent URL health checker
go run ./examples/sitecheck https://example.com

# TeamDesk SQLite API & Web app
go run ./examples/teamdesk -db /tmp/teamdesk.db
```

See also [`CHEATSHEET.md`](./CHEATSHEET.md) and [`EXERCISES.md`](./EXERCISES.md) for extra practice ideas.

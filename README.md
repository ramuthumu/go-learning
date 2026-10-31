# Go From Fundamentals (For Python Developers)

A simple, step-by-step, hands-on handbook for learning Go from core programming fundamentals. Every lesson contrasts Go with Python, highlights Go's golden rules, links to a runnable program in [`lessons/`](./lessons), and includes a quick self-check with a collapsible answer.

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
12. [Quick Cheat Sheet: Python vs. Go](#quick-cheat-sheet-python-vs-go)
13. [Additional Reference Examples](#additional-reference-examples)

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

---

## Lesson 10: Concurrency (`go` & Channels)

* **Runnable code:** [`lessons/10-concurrency/main.go`](./lessons/10-concurrency/main.go)
* **Run it:** `go run ./lessons/10-concurrency`

Concurrency is Go's superpower:
* **Goroutine (`go someFunc()`):** Starts a lightweight background task (takes only a few kilobytes of memory; you can run thousands at once).
* **Channel (`chan Type`):** A safe pipe for sending messages between goroutines (`ch <- msg` to send, `msg := <-ch` to receive).

---

## Lesson 11: Mini Project — Task Tracker

* **Runnable code:** [`lessons/11-mini-project/main.go`](./lessons/11-mini-project/main.go)
* **Run it:** `go run ./lessons/11-mini-project`

Combines **Structs**, **Slices**, **Maps**, **Pointer Methods**, and **Error Handling** into a complete in-memory Task Tracker CLI.

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

---

## Additional Reference Examples

When you are ready to explore larger real-world programs and unit tests in [`examples/`](./examples):

```sh
# Run all unit tests across the repository
go test -race ./...

# LogLens CLI log analyzer
go run ./examples/loglens -json examples/loglens/sample.log

# Sitecheck concurrent URL health checker
go run ./examples/sitecheck https://example.com
```

See also [`CHEATSHEET.md`](./CHEATSHEET.md) and [`EXERCISES.md`](./EXERCISES.md) for extra practice ideas.

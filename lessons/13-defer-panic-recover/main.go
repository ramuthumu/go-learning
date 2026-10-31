package main

import "fmt"

func demonstrateDefer() {
	fmt.Println("--- 1. How defer works ---")
	// In Python: you use 'with ...' or 'try...finally' for cleanup.
	// In Go: 'defer' runs a function right before the surrounding function finishes.
	defer fmt.Println("Deferred cleanup 1 (runs last!)")
	defer fmt.Println("Deferred cleanup 2 (runs second)")
	defer fmt.Println("Deferred cleanup 3 (runs first: LIFO order)")

	fmt.Println("Executing core logic inside demonstrateDefer()...")
	// Defers are stacked: Last-In, First-Out (LIFO)!
}

func safeWorker(shouldPanic bool) {
	// 2. recover() catches a panic, preventing the program from crashing
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Printf("\n--- 2. Running safeWorker(shouldPanic=%t) ---\n", shouldPanic)
	if shouldPanic {
		// panic is like raising an unhandled Exception in Python
		// Idiomatic Go uses 'error' return values; panic is ONLY for fatal/unrecoverable states
		panic("critical database connection lost!")
	}
	fmt.Println("safeWorker completed normally.")
}

func main() {
	demonstrateDefer()

	// Safe worker with normal execution
	safeWorker(false)

	// Safe worker triggering panic, but recovered gracefully!
	safeWorker(true)

	fmt.Println("\nProgram continued running successfully after recover!")
}

package main

import (
	"errors"
	"fmt"
)

// 1. Basic function: takes two ints, returns one int
func add(a int, b int) int {
	return a + b
}

// Shortcut: if parameters share the same type, you only need to write the type once
func multiply(a, b int) int {
	return a * b
}

// 2. Multiple return values: returns quotient and remainder
func divideWithRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// 3. The Go Way: Returning a result and an error
// (Go does NOT have try/except; it returns errors as values!)
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil // 'nil' means "no error" (like None in Python)
}

func main() {
	// Calling basic functions
	sum := add(10, 5)
	product := multiply(10, 5)
	fmt.Printf("10 + 5 = %d\n", sum)
	fmt.Printf("10 * 5 = %d\n", product)

	// Calling multiple-return function
	q, r := divideWithRemainder(10, 3)
	fmt.Printf("10 / 3 = %d with remainder %d\n", q, r)

	// Calling function with error check
	fmt.Println("\n--- Safe Division Examples ---")

	result1, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("10 / 2 =", result1)
	}

	result2, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("10 / 0 =", result2)
	}
}

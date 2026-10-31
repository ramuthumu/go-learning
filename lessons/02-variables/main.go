package main

import "fmt"

func main() {
	// 1. Explicit declaration with var and type
	var greeting string = "Welcome to Go!"

	// 2. Type inference with var
	var year = 2026

	// 3. Short declaration := (The most common way inside functions!)
	// ':=' creates a brand-new variable and infers its type automatically
	name := "Ramu"
	age := 30
	isCoding := true
	rating := 4.9

	// 4. Reassigning with '=' (updates an existing variable)
	age = 31

	fmt.Println(greeting, year)
	fmt.Printf("Name: %s | Age: %d | Coding: %t | Rating: %.1f\n", name, age, isCoding, rating)

	// 5. Zero Values (Default values when declared without initializing)
	// Go NEVER has uninitialized garbage values or 'undefined'
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultString string // "" (empty string)
	var defaultBool bool     // false

	fmt.Printf("Zero values -> int: %d, float: %.1f, string: %q, bool: %t\n",
		defaultInt, defaultFloat, defaultString, defaultBool)
}

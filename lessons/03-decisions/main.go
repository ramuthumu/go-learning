package main

import "fmt"

func main() {
	// 1. Basic if / else if / else
	age := 20

	if age >= 18 {
		fmt.Println("Eligible to vote!")
	} else if age >= 16 {
		fmt.Println("Almost there, can drive but not vote yet.")
	} else {
		fmt.Println("Too young to vote.")
	}

	// 2. Logical Operators: && (AND), || (OR), ! (NOT)
	hasLicense := true
	hasInsurance := false

	if hasLicense && hasInsurance {
		fmt.Println("You can drive the car legally.")
	} else if hasLicense && !hasInsurance {
		fmt.Println("You have a license, but you need insurance!")
	}

	// 3. Go's special feature: "if with a short statement"
	// We can declare a variable right before checking the condition.
	// Syntax: if <statement>; <condition> { ... }
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}
	// Note: 'score' only exists inside the if/else block above!
}

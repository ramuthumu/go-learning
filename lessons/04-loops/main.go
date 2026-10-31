package main

import "fmt"

func main() {
	// 1. Classic 3-part loop: init; condition; post
	fmt.Println("--- 1. Classic 3-part loop ---")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// 2. While-style loop (condition only)
	fmt.Println("\n--- 2. While-style loop ---")
	energy := 3
	for energy > 0 {
		fmt.Printf("Energy remaining: %d\n", energy)
		energy-- // Decrement by 1
	}
	fmt.Println("Tired! Need sleep.")

	// 3. Infinite loop with break and continue
	fmt.Println("\n--- 3. Loop with break & continue ---")
	num := 0
	for {
		num++

		if num == 2 {
			fmt.Println("Skipping 2 using continue")
			continue // Skip rest of this iteration
		}

		if num > 4 {
			fmt.Println("Stopping loop using break")
			break // Exit the loop entirely
		}

		fmt.Printf("Number: %d\n", num)
	}
}

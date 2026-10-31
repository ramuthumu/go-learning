package main

import "fmt"

func main() {
	// 1. Slices: Dynamic lists (like Python lists)
	// Notice: [] has NO number inside (unlike fixed-size arrays)
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("Initial slice:", fruits)
	fmt.Printf("Length: %d\n", len(fruits))

	// 2. Appending elements: append(slice, newElement)
	// In Go, append returns the updated slice, so we reassign it:
	fruits = append(fruits, "orange")
	fmt.Println("After append:", fruits)

	// 3. Accessing items by 0-based index
	fmt.Println("First fruit:", fruits[0])

	// 4. Looping over a slice using 'range' (Like Python's 'enumerate')
	fmt.Println("\n--- Looping with index and value ---")
	for index, fruit := range fruits {
		fmt.Printf("Index %d has: %s\n", index, fruit)
	}

	// 5. If you only care about the value (ignore index with _)
	fmt.Println("\n--- Looping values only ---")
	for _, fruit := range fruits {
		fmt.Printf("Fruit: %s\n", fruit)
	}
}

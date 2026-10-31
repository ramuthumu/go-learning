package main

import "fmt"

// modifyValue receives a COPY of the number (Python-like immutable primitive behavior)
func modifyValue(val int) {
	val = 999
}

// modifyPointer receives the MEMORY ADDRESS of the number (allows modifying caller's variable)
func modifyPointer(ptr *int) {
	*ptr = 999 // The * dereferences the pointer: "go to this address and write 999"
}

func main() {
	// 1. Value vs Address (& operator)
	score := 42
	fmt.Printf("1. score value: %d\n", score)
	fmt.Printf("   score memory address (&score): %p\n", &score)

	// 2. Pointer variables (*Type)
	// A pointer variable stores a memory address
	var ptr *int = &score
	fmt.Printf("2. ptr stores address: %p\n", ptr)
	fmt.Printf("   *ptr (dereference) reads value at address: %d\n", *ptr)

	// Changing value through pointer
	*ptr = 100
	fmt.Printf("   After *ptr = 100, score is now: %d\n", score)

	// 3. Passing by Value vs Pointer to functions
	fmt.Println("\n--- 3. Function arguments ---")
	num := 10
	fmt.Println("Initial num:", num)

	modifyValue(num)
	fmt.Println("After modifyValue(num):", num) // Still 10!

	modifyPointer(&num)                            // Pass the memory address using &
	fmt.Println("After modifyPointer(&num):", num) // Changed to 999!

	// 4. Why Go pointers are safe
	// In C/C++, you can do pointer arithmetic (ptr + 1) which causes crashes/exploits.
	// In Go, pointer arithmetic is FORBIDDEN. Pointers are only for sharing and mutating data safely!
}

package main

import "fmt"

func main() {
	// 1. Creating a Map: map[KeyType]ValueType (like a Python dict)
	// All keys must be the same type; all values must be the same type.
	phonebook := map[string]string{
		"Alice": "555-0101",
		"Bob":   "555-0102",
	}

	// 2. Adding / Updating a key
	phonebook["Ramu"] = "555-0103"
	fmt.Println("Phonebook:", phonebook)

	// 3. Deleting a key
	delete(phonebook, "Bob")
	fmt.Println("After deleting Bob:", phonebook)

	// 4. The "Comma OK" idiom (Crucial Go Pattern!)
	// In Python, accessing a missing key raises a KeyError.
	// In Go, it returns the zero value ("" for strings, 0 for ints).
	// To check if a key ACTUALLY exists:
	number, exists := phonebook["Charlie"]
	if exists {
		fmt.Println("Charlie's number:", number)
	} else {
		fmt.Println("Charlie is not in the phonebook!")
	}

	// 5. Looping through a map with 'range'
	fmt.Println("\n--- Phonebook Directory ---")
	for name, num := range phonebook {
		fmt.Printf("%s -> %s\n", name, num)
	}
}

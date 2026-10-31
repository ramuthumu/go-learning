package main

import "fmt"

// 1. Generic Function with 'any' constraint
// T can be any type: int, string, struct, etc.
// In Python: def first(items): return items[0] if items else None
func First[T any](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T // Returns the zero value of type T
		return zero, false
	}
	return items[0], true
}

// 2. Generic Function with 'comparable' constraint
// 'comparable' allows using == and !=
func Contains[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// 3. Custom Constraint using Union Interface
// Accepts int OR float64
type Number interface {
	int | int64 | float64
}

func Sum[T Number](numbers []T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

// 4. Generic Struct: Safe Result Container (Result[T])
type ApiResponse[T any] struct {
	Status  int    `json:"status"`
	Payload T      `json:"payload"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// Testing First[T] with ints and strings
	intList := []int{10, 20, 30}
	strList := []string{"Go", "Python", "Rust"}

	firstInt, _ := First(intList)
	firstStr, _ := First(strList)
	fmt.Printf("1. First int: %d\n", firstInt)
	fmt.Printf("   First string: %s\n", firstStr)

	// Testing Contains[T]
	fmt.Println("\n--- 2. Contains[T comparable] ---")
	fmt.Println("strList contains 'Go':", Contains(strList, "Go"))
	fmt.Println("strList contains 'Java':", Contains(strList, "Java"))
	fmt.Println("intList contains 20:", Contains(intList, 20))

	// Testing Sum[T Number]
	fmt.Println("\n--- 3. Number constraint ---")
	intSum := Sum([]int{1, 2, 3, 4, 5})
	floatSum := Sum([]float64{1.5, 2.5, 3.0})
	fmt.Printf("Integer sum: %d\n", intSum)
	fmt.Printf("Float sum: %.2f\n", floatSum)

	// Testing Generic Struct
	fmt.Println("\n--- 4. Generic Struct ---")
	respString := ApiResponse[string]{
		Status:  200,
		Payload: "User account verified",
	}
	respSlice := ApiResponse[[]int]{
		Status:  200,
		Payload: []int{101, 102, 103},
	}
	fmt.Printf("API String Response: %+v\n", respString)
	fmt.Printf("API Slice Response:  %+v\n", respSlice)
}

package main

import (
	"errors"
	"fmt"
)

// CalculateDiscount applies a percentage discount (0-100) to an amount.
func CalculateDiscount(price float64, discountPercent int) (float64, error) {
	if price < 0 {
		return 0, errors.New("price cannot be negative")
	}
	if discountPercent < 0 || discountPercent > 100 {
		return 0, errors.New("discount percent must be between 0 and 100")
	}
	discount := price * float64(discountPercent) / 100.0
	return price - discount, nil
}

func main() {
	finalPrice, err := CalculateDiscount(100.0, 20)
	if err != nil {
		fmt.Printf("Calculation error: %v\n", err)
		return
	}
	fmt.Printf("Original $100 with 20%% discount = $%.2f\n", finalPrice)
	fmt.Println("Run `go test -v ./lessons/17-testing` to run the automated unit tests!")
}

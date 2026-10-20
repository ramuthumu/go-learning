package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrInvalidQuantity = errors.New("quantity must be positive")

func parseQuantity(raw string) (int, error) {
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parse quantity %q: %w", raw, err)
	}
	if n < 1 {
		return 0, ErrInvalidQuantity
	}
	return n, nil
}

func main() {
	for _, raw := range []string{"4", "0", "oops"} {
		n, err := parseQuantity(raw)
		if errors.Is(err, ErrInvalidQuantity) {
			fmt.Println("please order at least one item")
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("accepted", n)
	}
}

package main

import "fmt"

func total(unitCents, quantity int) int { return unitCents * quantity }

func main() {
	cents := total(1250, 3)
	label := "small order"
	if cents >= 3000 {
		label = "large order"
	}
	fmt.Printf("%s: $%d.%02d\n", label, cents/100, cents%100)
	for i := 1; i <= 3; i++ {
		fmt.Println("package", i)
	}
}

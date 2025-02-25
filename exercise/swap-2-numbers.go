package main

import "fmt"

func swapNumber(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("Swap two numbers using pointer.")

	a, b := 5, 10

	// Print value before swapping
	fmt.Print("Value before swapping: ")
	fmt.Printf("a = %d, b = %d\n", a, b)

	swapNumber(&a, &b)

	// Print value after swapping
	fmt.Print("Value after swapping: ")
	fmt.Printf("a = %d, b = %d\n", a, b)
}

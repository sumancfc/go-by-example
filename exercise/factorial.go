package main

import "fmt"

func main() {
	var number int
	factorial := 1

	fmt.Printf("Enter Number: ")
	_, err := fmt.Scanf("%d", &number)

	if err != nil {
		return
	}

	for i := 1; i <= number; i++ {
		factorial *= i
	}

	fmt.Printf("Factorial of %d is %d\n", number, factorial)
}

package main

import "fmt"

func main() {
	var number int
	a, b := 0, 1

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanf("%d", &number)

	if err != nil {
		return
	}

	fmt.Print("Fibonacci Series: ", a, " ", b, " ")

	for i := 3; i <= number; i++ {
		c := a + b

		fmt.Print(c, " ")

		a = b
		b = c
	}
}

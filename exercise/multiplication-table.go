package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter number:")
	_, err := fmt.Scanf("%d", &number)

	if err != nil {
		return
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
	}
}

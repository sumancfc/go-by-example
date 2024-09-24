package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Infinite Loop
	i := 0
	for {
		fmt.Println("Infinite loop, iteration:", i)
		i++
		if i == 10 {
			break
		}
	}

	// for as a while Loop
	for i < 20 {
		fmt.Println("While-like loop, iteration:", i)
		i++
	}

	// for Loop with range
	numbers := []int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Printf("Index: %d, value: %d \n", index, value)
	}

	// Iterating Over a String
	word := "Hello"

	for index, char := range word {
		fmt.Printf("Index: %d, Character: %c \n", index, char)
	}

	// Nested for Loop
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}

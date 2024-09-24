package main

import "fmt"

func main() {
	// Initializing and using variables in if
	if age := 18; age >= 18 {
		fmt.Println("You are eligible to Drive")
	}

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
}

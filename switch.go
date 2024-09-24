package main

import "fmt"

func main() {
	day := "Tuesday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a Weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's a Weekend.")
	default:
		fmt.Println("Invalid Day.")
	}

	// Switch Statement with no condition
	num := 10
	switch { // switch without expression acts like multiple if-else if condition
	case num < 0:
		fmt.Println("Negative Number.")
	case num == 0:
		fmt.Println("Zero Number.")
	case num > 0:
		fmt.Println("Positive Number.")
	}
}

package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter a year:")
	_, err := fmt.Scanln(&year)

	if err != nil {
		return
	}

	// If-else condition
	if year%400 == 0 {
		fmt.Println(year, "is a leap year.")
	} else if year%100 == 0 {
		fmt.Println(year, "is not a leap year")
	} else if year%4 == 0 {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}

	// Switch condition
	switch {
	case year%400 == 0:
		fmt.Println("It is a leap year.")
	case year%100 == 0:
		fmt.Println("It is not a leap year.")
	case year%4 == 0:
		fmt.Println("It is a leap year.")
	default:
		fmt.Println("It is not a leap year.")
	}

}

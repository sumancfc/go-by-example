package main

import "fmt"

func main() {
	var marks int
	fmt.Print("Enter a number between 0 - 100: ")
	_, err := fmt.Scanf("%d", &marks)

	if err != nil {
		return
	}

	fmt.Println("If-else Condition")
	// if-else condition
	if marks >= 90 && marks <= 100 {
		fmt.Println("Grade: A")
	} else if marks >= 80 && marks < 90 {
		fmt.Println("Grade: B")
	} else if marks >= 70 && marks < 80 {
		fmt.Println("Grade: C")
	} else if marks >= 60 && marks < 70 {
		fmt.Println("Grade: D")
	} else if marks >= 0 && marks < 60 {
		fmt.Println("Fail")
	} else {
		fmt.Println("Invalid Marks. Please enter a value between 0 and 100.")
	}

	fmt.Println("Switch Condition")
	// Switch Condition
	switch {
	case marks >= 90 && marks <= 100:
		fmt.Println("Grade: A")
	case marks >= 80 && marks < 90:
		fmt.Println("Grade: B")
	case marks >= 70 && marks < 80:
		fmt.Println("Grade: C")
	case marks >= 60 && marks < 70:
		fmt.Println("Grade: D")
	case marks >= 0 && marks < 60:
		fmt.Println("Fail")
	default:
		fmt.Println("Invalid Marks. Please enter a value between 0 and 100.")
	}
}

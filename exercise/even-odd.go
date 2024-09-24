package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}

	if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}

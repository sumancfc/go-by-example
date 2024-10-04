package main

import "fmt"

// change by reference
func changeNumber(num *int) {
	*num = 5
	fmt.Println("Change in Number", *num)
}

func main() {
	num := 1
	changeNumber(&num)
	fmt.Println("After Change in Number", num)
}

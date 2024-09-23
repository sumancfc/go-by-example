package main

import (
	"fmt"
)

/**
Array syntax in GO ==> var arrayName [size]elementType
- var is used to declare a variable.
- arrayName is the name of the array.
- [size] specifies the number of elements in the array (fixed size).
- elementType specifies the type of the elements in the array (e.g., int, string, etc.).
- Array: Fixed size, cannot be resized.
**/

func main() {
	var arr [5]int
	// Assign values to the array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	fmt.Println("Array Elements:", arr)

	// Iterating over an Array
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr[%d] = %d\n", i, arr[i])
	}

	// Declare and initialize an array with values
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruit Array:", fruits)

	// short-hand declaration of an array
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8} // declare any length of array with [...]
	fmt.Println("Number Array:", nums)
	fmt.Println("Lenght of Array:", len(nums))
}

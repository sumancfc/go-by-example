package main

import "fmt"

func main() {
	//A slice of Integers
	//slice := []int{1, 2,3}

	// A slice from existing array
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:]

	//// Using Make()
	slice1 := make([]int, 5)    // length 5, capacity 5
	slice2 := make([]int, 3, 6) // length 3, capacity 6
	//
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(slice1, slice2)

	numbers := []int{10, 20, 30, 40, 50}

	// Access elements in a slice
	fmt.Println("First Element:", numbers[0])
	fmt.Println("Slice Length:", len(numbers))
	fmt.Println("Slice Capacity:", cap(numbers))

	// Modify an element in a slice
	numbers[1] = 25
	fmt.Println("Modified Slice:", numbers)

	// Slice a slice
	newSlice := numbers[1:4]
	fmt.Println("Sliced Slice:", newSlice)

	// Append to a slice
	numbers = append(numbers, 60, 70, 80)
	fmt.Println("Extended Slice:", numbers)

	// Append another slice to a slice
	moreNumbers := []int{90, 100}
	numbers = append(numbers, moreNumbers...)

	fmt.Println("After appending another slice:", numbers)

	// Copy element in a slice
	source := []int{1, 2, 3, 4, 5}
	dest := make([]int, 5)

	// Copy element from source to dest
	count := copy(dest, source)

	fmt.Println("Copied elements:", dest)
	fmt.Println("Number of elements copied:", count)
}

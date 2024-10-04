package main

import (
	//"errors"
	"fmt"
)

// Function that can return an error
//func divide(a, b float64) (float64, error) {
//	if b == 0 {
//		// Creating an error using errors.New
//		return 0, errors.New("cannot divide by zero")
//	}
//
//	return a / b, nil
//}

// DivideError custom error type
type DivideError struct {
	dividend float64
	divisor  float64
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by %.2f", e.dividend, e.divisor)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Returning an instance of DivideError
		return 0, &DivideError{a, b}
	}

	return a / b, nil
}

func main() {
	result, err := divide(4, 0) // Division by zero

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

}

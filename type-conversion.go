package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Integer to String
	intVal := 123
	strVal := strconv.Itoa(intVal)

	// Float to String
	floatVal := 123.456
	strFloatVal := strconv.FormatFloat(floatVal, 'f', 1, 64)

	fmt.Printf("Integer: %d, String: %s\n", intVal, strVal)
	fmt.Printf("Float: %.2f, String: %s\n", floatVal, strFloatVal)

	// String to Integer
	stringValue := "123"
	integerValue, err := strconv.Atoi(stringValue)

	if err != nil {
		fmt.Println("Conversion Error:", err)
	} else {
		fmt.Printf("Integer: %d, String: %s\n", integerValue, stringValue)
	}

	// String to Float
	stringFloatVal := "123.456"
	floatValFloat, err := strconv.ParseFloat(stringFloatVal, 64)

	if err != nil {
		fmt.Println("Conversion Error:", err)
	} else {
		fmt.Printf("Float: %.2f, String: %s\n", floatValFloat, stringFloatVal)
	}

	// String to Boolean
	strBoolVal := "true"
	boolVal, err := strconv.ParseBool(strBoolVal)

	if err != nil {
		fmt.Println("Conversion Error:", err)
	} else {
		fmt.Printf("Boolean: %t\n", boolVal)
	}

	// Complex Number in GO
	comp := complex(5, 7)
	fmt.Printf("Complex: %v\n", comp)
}

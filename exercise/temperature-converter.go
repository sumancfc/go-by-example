package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32.0
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var temperature float64
	var unit string

	// Temperature input from user
	fmt.Print("Enter the value of Temperature: ")
	_, err := fmt.Scanln(&temperature)
	if err != nil {
		return
	}

	// Unit input from user
	fmt.Print("Is this Celsius or Fahrenheit (C/F)? ")
	_, err = fmt.Scanln(&unit)
	if err != nil {
		return
	}

	if unit == "C" || unit == "c" {
		fahrenheit := celsiusToFahrenheit(temperature)
		fmt.Printf("The Temperature in Fahrenheit is: %.2f\n", fahrenheit)
	} else if unit == "F" || unit == "f" {
		celsius := fahrenheitToCelsius(temperature)
		fmt.Printf("The Temperature in Celsius is: %.2f\n", celsius)
	} else {
		fmt.Println("Please enter a valid unit")
	}

}

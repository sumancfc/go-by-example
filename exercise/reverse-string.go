package main

import "fmt"

func main() {
	str := "Hello World"
	reversed := ""

	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	fmt.Println("Reversed string is:", reversed)
}

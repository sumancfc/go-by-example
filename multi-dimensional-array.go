package main

import "fmt"

func main() {
	var matrix [2][2]int

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 4
	matrix[1][1] = 5

	fmt.Println("Matrix")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

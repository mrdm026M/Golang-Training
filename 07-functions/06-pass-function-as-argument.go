package main

import "fmt"

// Pass function as an argument to another function in Go

func main() {
	// passing area and sum function as argument in print function
	print(area, 2, 4)
	print(sum, 2, 4)
}

func print(result func(int, int) int, a int, b int) {
	fmt.Println(result(a, b))
}

func area(a int, b int) int {
	return a * b
}

func sum(a int, b int) int {
	return a + b
}

package main

import "fmt"

// Function as Value or Anonymous functions
// It is also called as anonymous functions because a function is not named and can be assigned to a variable and passed around.

var max = func(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	result := max(10, 15)
	fmt.Print(result)
}

package main

import "fmt"

// Variadic Function
// A function that can accept a dynamic number of arguments is called a Variadic function.
// Three dots are used as a prefix before type.

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))
}

func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

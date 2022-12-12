package main

import "fmt"

func main() {
	result := sum(10, 15)
	fmt.Print(result)

	// returning multiple values

}

// func sum_avg(a, b int) (int, int) -> return multiple values
// func sum(a, b int) (int, error) -> error can also be returned
// result, err := sum(2, 3)

func sum_avg(a int, b int) (int, int) {
	add := a + b
	avg := add / 2

	return add, avg
}

func sum(a int, b int) int {
	return a + b
}

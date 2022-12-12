// Check if a string is a number in GO
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "12"
	val, err := strconv.Atoi(x)

	if err != nil {
		fmt.Printf("X: %s is not a number", x)
	} else {
		fmt.Printf("X: %d is a number", val)
	}
}

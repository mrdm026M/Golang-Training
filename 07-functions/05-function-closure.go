package main

import "fmt"

// Function closures are nothing but an anonymous function that can access variables declared outside the function and
// also retain the current value of those variables between different function calls.

// A closure happens
// 1. when a function is defined within a different function and
// 2. the inner function can access the variable of the outer function.

func main() {
	modulus := getModulus() // getModulus() return closure and assign to variable modulus.
	modulus(-1)
	modulus(2)
	modulus(-5)

	// Example 2 -
	count := 0
	for i := 1; i <= 5; i++ {
		func() {
			count++
			fmt.Println(count)
		}()
	}
}

func getModulus() func(int) int {
	count := 0               // variable outside the inner function
	return func(x int) int { // function inside another function
		count = count + 1 // inner function access the count variable.
		fmt.Printf("modulus function called %d times\n", count)
		if x < 0 {
			x = x * -1
		}
		return x
	}
}

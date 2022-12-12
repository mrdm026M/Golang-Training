package main

import "fmt"

// Only if
// If Else
// If Else Ladder
// Nested if-else
// If with a short statement
// no ternary operator in Go

func main() {
	a := 10

	// only if
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	// if-else
	if a < 5 {
		fmt.Println("a is smaller than 5")
	} else {
		fmt.Println("a is greater than 5")
	}

	// if-else-ladder
	if a < 5 {
		fmt.Println("a is smaller than 5")
	} else if a == 5 {
		fmt.Println("a is equal to 5")
	} else {
		fmt.Println("a is greater than 5")
	}

	// nested if-else
	b := 1
	c := 2
	d := 3

	if b > c {
		if b > d {
			fmt.Println("Biggest is b")
		} else {
			fmt.Println("Biggest is d")
		}
	} else if c > d {
		fmt.Println("Biggest is c")
	} else {
		fmt.Println("Biggest is d")
	}

	// if short version
	if e := 5; e == 5 {
		fmt.Println("E is equal to 5")
	}
}

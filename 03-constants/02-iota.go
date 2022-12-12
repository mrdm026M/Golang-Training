package main

import "fmt"

// Iota is an identifier which is used with constant and which can simplify constant definitions that use auto increment numbers.
// The IOTA keyword represent integer constant starting from zero.
// They can also be used to create enum in Go.

// IOTA is
// 1. A counter which starts with zero
// 2. Increases by 1 after each line
// 3. Is only used with constant

// creating enum in go using iota
type Size int

const (
	small Size = iota
	medium
	large
	extraLarge
)

func main() {
	const (
		a = 0
		b = 1
		c = 2
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// auto increament values from 0
	const (
		d = iota
		e
		f
	)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	// prinitng enum
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)
}

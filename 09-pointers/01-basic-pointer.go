package main

import "fmt"

// Pointer is a variable that holds a memory address of another variable.

// The * operator can be used to:
// 1. Dereference a pointer which means getting the value at address stored in the pointer.
// 2. Change the value at that pointer location as well

func main() {
	// Declare a pointer
	var b *int
	var c int
	a := 2
	b = &a // address of a
	c = *b // value at the address stored in b

	fmt.Println("Value of a: ", a)
	fmt.Println("Value of b: ", b)
	fmt.Println("Value of *b: ", *b)
	fmt.Println("Value of c: ", c)
	fmt.Println("-----------------")

	b = new(int)
	*b = 10
	fmt.Println("Value of b: ", b)
	fmt.Println("Value of *b: ", *b)
	fmt.Println("Value of &b: ", &b)
	fmt.Println("-----------------")

	d := 2
	e := &d
	fmt.Println("Value of d: ", d)
	fmt.Println("Value of *e: ", *e)

	*e = 3
	fmt.Println("Value of d: ", d)
	fmt.Println("Value of *e: ", *e)

	d = 4
	fmt.Println("Value of d: ", d)
	fmt.Println("Value of *e: ", *e)
}

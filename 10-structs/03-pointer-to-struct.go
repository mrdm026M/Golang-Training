// There are two ways of creating a pointer to the struct
// Using the & operator
// Using the new keyword

// 1. Using the & operator
// The & operator can be used to get the pointer to a struct variable.
package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	// 1
	emp := employee{name: "Sam", age: 31, salary: 2000}
	empP := &emp
	fmt.Printf("Emp: %+v\n", empP)
	empP = &employee{name: "John", age: 30, salary: 3000}
	fmt.Printf("Emp: %+v\n", empP)

	// 2
	empN := new(employee)
	fmt.Printf("Emp Pointer Address: %p\n", empN)
	fmt.Printf("Emp Pointer: %+v\n", empN)
	fmt.Printf("Emp Value: %+v\n", *empN)
}

// 2. Using the  new() keyword will:
// Create the struct
// Initialize all the field to the zero default value of their type
// Return the pointer to the newly created struct

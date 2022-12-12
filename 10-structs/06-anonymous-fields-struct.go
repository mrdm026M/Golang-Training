// A struct can have anonymous fields as well, meaning a field having no name.
// The type will become the field name.

package main

import "fmt"

type employee struct {
	string // anonymous field
	age    int
	salary int
}

func main() {
	emp := employee{string: "Sam", age: 31, salary: 2000}

	//Accessing a struct field
	n := emp.string
	fmt.Printf("Current name is: %s\n", n)

	//Assigning a new value
	emp.string = "John"
	fmt.Printf("New name is: %s\n", emp.string)
}

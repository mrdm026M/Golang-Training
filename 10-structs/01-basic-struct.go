// GO struct is named collection of data fields which can be of different types.
// Struct acts as a container that has different heterogeneous data types which together represents an entity.

package main

import "fmt"

// example - struct
type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	// empty struct - fields will get default value of their datatypes
	emp1 := employee{}
	fmt.Printf("Emp1: %+v\n", emp1)

	emp2 := employee{
		name:   "Sam",
		age:    31,
		salary: 2000,
	}
	fmt.Printf("Emp3: %+v\n", emp2)

	emp4 := employee{
		name: "Sam",
		age:  31,
	}
	fmt.Printf("Emp4: %+v\n", emp4)

	// other correct ways to initialize struct
	// if no comma is used at the last field it will raise error
	emp := employee{
		name:   "Sam",
		age:    31,
		salary: 2000}
	fmt.Printf("Emp: %+v\n", emp)

	// we can also initialize the struct with no field name, but then we have to give values in sequence order.
	emp5 := employee{"Sam", 31, 2000}
	fmt.Printf("Emp5: %+v\n", emp5)
}

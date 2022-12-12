package main

import "fmt"

// Method on struct
type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) setNewNameNoUpdate(newName string) {
	e.name = newName
}

func (e *employee) setNewNameWillUpdate(newName string) {
	e.name = newName
}

func main() {
	// will not update since a copy of previous value is kept in the receiver
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.setNewNameNoUpdate("John")
	fmt.Printf("Name: %s\n", emp.name)

	// will update the name value
	emp.setNewNameWillUpdate("John")
	fmt.Printf("Name: %s\n", emp.name)

	(&emp).setNewNameWillUpdate("Mike")
	fmt.Printf("Name: %s\n", emp.name)
}

// Method Chaining
// For method chaining to be possible the methods in the chain should return the receiver.
// Returning the receiver for the last method inn the chain is optional.

package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) printName() employee {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("Salary: %d\n", e.salary)
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}

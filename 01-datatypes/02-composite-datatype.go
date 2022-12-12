package main

import "fmt"

// Declare a struct
type employee struct {
	name   string
	age    int
	salary float64
}

func main() {
	// non reference type
	// array
	sample := [3]string{"a", "b", "c"}
	len := len(sample)

	fmt.Println("sample array := ", sample)
	fmt.Println("len := ", len)

	// struct
	// Initialize a struct without named fields
	employee1 := employee{"John", 21, 1000}

	//Initialize a struct with named fields
	employee2 := employee{
		name:   "Sam",
		age:    22,
		salary: 1100,
	}

	//Initializing only some fields. Other values are initialized to default zero value of that type
	employee3 := employee{name: "Tina", age: 24}

	fmt.Println("Employee 1 := ", employee1)
	fmt.Println("Employee 2 := ", employee2)
	fmt.Println("Employee 3 := ", employee3)

	// reference type
	// slice
	s := make([]string, 2, 3)
	p := []string{"a", "b", "c"} //Direct intialization

	fmt.Println(s)
	fmt.Println(p)

	// channel
	// maps
	// pointer
	// function & method
	// interface
}

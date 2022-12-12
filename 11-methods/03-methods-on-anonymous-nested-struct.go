package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
	address
}

type address struct {
	city    string
	country string
}

func (e employee) empDetails() {
	fmt.Printf("Employee: %+v\n", e)
}

func (a address) details() {
	fmt.Printf("City: %s\n", a.city)
	fmt.Printf("Country: %s\n", a.country)
}

func main() {
	address := address{city: "London", country: "UK"}

	emp := employee{name: "Sam", age: 31, salary: 2000, address: address}

	emp.empDetails()

	// method of anonymous struct can be accessed in two ways-
	emp.details()
	emp.address.details()
}

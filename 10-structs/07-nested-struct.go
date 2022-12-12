package main

import (
	"fmt"
)

type employee struct {
	name    string
	age     int
	salary  int
	address address
}

type address struct {
	city    string
	country string
}

func main() {
	ad := address{
		city:    "London",
		country: "UK",
	}

	emp := employee{
		name:    "Dhruv",
		age:     21,
		salary:  210,
		address: ad,
	}

	fmt.Println(emp.address.city)
}

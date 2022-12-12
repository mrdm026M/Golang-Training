// A struct in go also allows adding metadata to its fields.
// These meta fields can be used to encode decode into different forms, doing some forms of validations on struct fields, etc.
// Meta-data is a string literal i.e it is enclosed in backquotes

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func main() {
	emp := employee{Name: "Sam", Age: 31, Salary: 2000}

	//Converting to json
	empJSON, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(empJSON))
}

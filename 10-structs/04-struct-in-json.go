// Printing struct in JSON format
// Marshal and MarshalIndent function of encoding/json package can be used to print a struct in JSON format.

// Difference -
// 1. Marshal – This function returns the JSON encoding of v by traversing the value recursively.
// 2. MarshalIndent– It is similar to Marshal function but applies Indent to format the output. So it can be used to pretty print a struct.

// It is to be noted that both Marshal and MarshalIndent function can only access the exported fields of a struct, which means that only the capitalized fields can be accessed and encoded in JSON form.

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string
	Age    int
	salary int // The salary field is not printed in the output because it begins with a lowercase letter and is not exported
}

func main() {
	emp := employee{Name: "Sam", Age: 31, salary: 2000}

	//Marshal
	empJSON, err := json.Marshal(emp)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Marshal: %s\n", string(empJSON))

	//MarshalIndent
	empJSON, err = json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent: %s\n", string(empJSON))
}

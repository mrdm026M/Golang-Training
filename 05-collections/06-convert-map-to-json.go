package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string
	Age  int
}

func main() {
	// converting map to json - example 1
	newMap := make(map[int]string)
	newMap[1] = "John"

	j, err := json.Marshal(newMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	// converting map to json - example 2 (Struct)
	a := make(map[string]employee)
	a["1"] = employee{Name: "John", Age: 21}
	l, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(l))
	}

	// converting json to map
	var b map[int]string
	json.Unmarshal(j, &b)

	fmt.Println(b)
}

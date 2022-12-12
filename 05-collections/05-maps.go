package main

import "fmt"

func main() {
	// 1. creating map -
	// a. using map<key datatype>value datatype{}
	// b. using make

	// 1.a using map<key datatype>value datatype{}
	// Declare
	employeeSalary := map[string]int{}

	//Intialize using map lieteral
	employeeSalary = map[string]int{
		"John": 1000,
		"Sam":  1200,
	}

	//Adding a key value
	employeeSalary["Tom"] = 2000
	fmt.Println(employeeSalary)

	// 1.b using make
	employeeSalaryMake := make(map[string]int)

	//Adding a key value
	employeeSalaryMake["Tom"] = 2000
	fmt.Println(employeeSalaryMake)

	fmt.Println("----------------------------------------")

	// Map Operations
	// The below operations are applicable for map

	// Add a key-value pair
	// Update a key-value pair
	// Get the value corresponding to a key
	// Delete a key-value pair
	// Check if a key exists

	newMap := make(map[string]int)
	// 1. Add a key-value pair
	newMap["Hello"] = 1
	fmt.Println("Adding key-value to newMap: ", newMap)

	// 2. Update a key-value pair
	newMap["Hello"] = 2
	fmt.Println("Updating key-value in newMap: ", newMap)

	// 3. Get the value corresponding to a key
	num := newMap["Hello"]
	fmt.Println("Fetching value from newMap using key: ", num)

	// 4. Delete a key-value pair
	delete(newMap, "Hello")
	fmt.Println("Deleteing a key-value pair in newMap: ", newMap)

	// 5. Check if a key exists
	newMap["Hello"] = 2
	val, exists := newMap["Hello"]
	fmt.Println("Checking if a key exists or not")
	fmt.Printf("Val: %d | exists: %t\n", val, exists)

	fmt.Println("----------------------------------------")

	// Iterate over a map
	// 1. Iterating over all keys and values
	// 2. Iterating over only keys
	// 3. Iterating over only values
	// 4. Get list of all keys

	sampleMap := map[string]string{
		"Dhruv": "Maheshwari",
		"Hello": "World",
	}

	// Iterating over all keys and values
	for k, v := range sampleMap {
		fmt.Printf("Key: %s | Value: %s | Full String: %s\n", k, v, k+" "+v)
	}

	// Iterating over only keys
	for k := range sampleMap {
		fmt.Printf("Key: %s\n", k)
	}

	// Iterating over only values
	for _, v := range sampleMap {
		fmt.Printf("Value: %s\n", v)
	}

	// Get list of all keys
	keys := getAllKeys(sampleMap)
	fmt.Println(keys)
}

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}

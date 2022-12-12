package main

import "fmt"

// for-range loop is used to iterate over different collection data structures in golang such as
// 1. array or slice
// 2. string
// 3. maps
// 4. channel

// Both index and value are optional in for-range when using with array/slice.
// Possible uses -
// 1. With index and value
// 2. With value only
// 3. With index only
// 4. Without index and value

func main() {
	names := []string{"dhruv", "nikhil", "amish"}

	// 1. With index and value
	fmt.Println("Both Index and Value")
	for i, name := range names {
		fmt.Printf("Index: %d | Value: %s\n", i, name)
	}

	// 2. Only value
	fmt.Println("Only value")
	for _, name := range names {
		fmt.Printf("Value: %s\n", name)
	}

	// 3. Only index
	fmt.Println("Only Index")
	for i := range names {
		fmt.Printf("Index: %d\n", i)
	}

	// 4. Without index and value. Just print array values
	fmt.Println("Without Index and Value")
	i := 0
	for range names {
		fmt.Printf("Index: %d | Value: %s\n", i, names[i])
		i++
	}
}

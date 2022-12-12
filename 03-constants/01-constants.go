package main

import "fmt"

const e = 123        // outer scope
const name = "Dhruv" // global const variable

// go doesn't support const map, struct, array & slice
func main() {
	// local const variables
	const a string = "10"
	const b = 10
	const (
		c = "String"
		d = 2.43
	)

	const e = 456 // inner scope
	testGlobal()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e) // output - 456 of inner scope constant
}

func testGlobal() {
	fmt.Println(name)
}

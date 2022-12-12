package main

import "fmt"

// Defer as the name suggests is used to defer the cleanup activities in a function.
// These cleanup activities will be performed at the end of the function.
// This cleanup activities will be done in a different function called by defer.
// This different function is called at the end of the surrounding function before it returns.

func main() {
	defer test()
	fmt.Println("Main function")
}

func test() {
	fmt.Println("Defer Test function")
}

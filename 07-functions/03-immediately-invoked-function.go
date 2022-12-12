package main

import "fmt"

// Immediately Invoked Function - IIF or Immediately Invoked Function are those function which can be defined and executed at the same time.
func main() {
	result := func() int {
		return 2 + 2
	}()

	fmt.Print(result)

}

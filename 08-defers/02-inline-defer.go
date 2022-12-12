package main

import "fmt"

func main() {
	defer func() { fmt.Println("Inline Defer Test Function") }()
	fmt.Println("Main Function")
}

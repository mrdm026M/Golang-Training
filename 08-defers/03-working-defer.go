package main

import "fmt"

// When the the compiler encounter a defer statement in a function it pushes it onto a list.
// This list internally implements a stack data structure.
// All the encountered defer statement in the same function are pushed onto this list.
// When the surrounding function returns then all the functions in the stack starting from top to bottom are executed before execution can begin in the calling function.
// Now same thing will happen in the calling function as well.

func main() {
	defer fmt.Println("Main Defer Function")
	fmt.Println("-------------------------")
	fmt.Println("Main Function Start")
	f1()
	fmt.Println("-------------------------")
	fmt.Println("Main Function Ends")
	fmt.Println("-------------------------")
}

func f1() {
	defer fmt.Println("F1 Defer Function")
	fmt.Println("-------------------------")
	fmt.Println("F1 Function Start")
	f2()
	fmt.Println("-------------------------")
	fmt.Println("F1 Function Ends")
	fmt.Println("-------------------------")
}

func f2() {
	defer fmt.Println("F2 Defer Function")
	fmt.Println("-------------------------")
	fmt.Println("F2 Function Start")
	fmt.Println("-------------------------")
	fmt.Println("F2 Function Ends")
	fmt.Println("-------------------------")
}

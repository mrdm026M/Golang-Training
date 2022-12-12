package main

import (
	"fmt"
)

func main() {
	// simple for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for as while loop
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	// infinite loop in for
	// b := 0
	// for {
	// 	fmt.Println(b)
	// 	b++
	// 	time.Sleep(time.Second * 1)
	// }

	// nested for loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// function call and assignment in init part of for loop
	i := 1
	for testFunc(); i < 3; i++ {
		fmt.Println(i)
	}

	for i := 2; i < 3; i++ {
		fmt.Println(i)
	}
}

func testFunc() {
	fmt.Println("Hello Test function in for loop")
}

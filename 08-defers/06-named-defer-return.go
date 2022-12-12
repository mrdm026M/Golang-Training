package main

import "fmt"

func main() {
	s := test()
	fmt.Println("Main Function")
	fmt.Println(s)
}

func test() (size int) {
	defer func() { size = 20 }()
	size = 30
	fmt.Println("Test Function")
	return
}

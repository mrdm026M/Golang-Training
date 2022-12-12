package main

import "fmt"

// Return a function from a function

func main() {
	areaF := getAreaFunc()
	res := areaF(2, 4)
	fmt.Println(res)
}

func getAreaFunc() func(int, int) int {
	return func(x, y int) int { // this function will be return and assigned to areaF variable in main function
		return x * y
	}
}

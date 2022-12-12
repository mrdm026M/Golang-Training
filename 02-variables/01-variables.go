package main

import "fmt"

var i int // global variable

func main() {
	// inner variable
	var a int = 10        // basic variable declaration and initailization
	var b, c int = 10, 11 // multiple variable declaration and initailization
	var (
		d int
		e int    = 12
		f string = "hello"
	)
	var g = 10 // no type defined - will inferred type automatically
	h := 3.4   // short hand version of variable declaration and intialization - type is inferred automatically

	if 10 == 10 {
		i = 10
	}

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

}

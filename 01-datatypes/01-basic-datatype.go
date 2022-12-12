package main

import "fmt"

func main() {
	// int
	var a int = 10
	b := 10

	fmt.Println("A := ", a)
	fmt.Println("B := ", b)

	// float
	var c float64 = 2.5
	d := 4.5

	fmt.Println("C := ", c)
	fmt.Println("D := ", d)

	// complex no
	var e complex128 = complex(c, d)
	f := 10 + 6i

	fmt.Println("E := ", e)
	fmt.Println("F := ", f)

	// byte
	var g byte = 'a'
	h := "abc"

	fmt.Println("G := ", g)
	fmt.Println("h := ", []byte(h))

	// String
	var i string = "hello"
	j := "world"

	fmt.Println(i + " " + j)

	// boolean
	var k bool
	l := (1 < 2)

	fmt.Println("K := ", k)
	fmt.Println("L := ", l)

	// rune
}

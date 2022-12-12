// Interface is a type in Go which is a collection of method signatures.
// These collections of method signatures are meant to represent certain behaviour.
// The interface declares only the method set and any type which implements all methods of the interface is of that interface type.

package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

type dog struct {
	age int
}

func (l dog) breathe() {
	fmt.Println("Dog breathes")
}

func (l dog) walk() {
	fmt.Println("Dog walk")
}

func main() {
	var a animal

	a = lion{age: 10}
	a.breathe()
	a.walk()

	a = dog{age: 5}
	a.breathe()
	a.walk()
}

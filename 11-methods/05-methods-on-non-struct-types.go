// Methods can also be defined on a non-struct custom type.
// Non-struct custom types can be created through type definition.
package main

import (
	"fmt"
	"math"
)

type myFloat float64

func (m myFloat) ceil() float64 {
	return math.Ceil(float64(m))
}

func main() {
	num := myFloat(1.4)
	fmt.Println(num.ceil())
}

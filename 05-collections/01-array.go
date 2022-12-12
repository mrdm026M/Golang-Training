package main

import "fmt"

func main() {
	// Declaration
	// Different ways of creating arrays in go -
	// 1. Both number of elements and actual elements
	num := [2]int{1, 2}
	fmt.Printf("Num: Len: %d | Value: %v\n", len(num), num)

	// 2. Only actual elements
	num2 := [...]int{2, 3}
	fmt.Printf("Num2: Len: %d | Value: %v\n", len(num2), num2)

	// 3. Only number of elements - no elements then default value 0 is used
	num3 := [2]int{}
	fmt.Printf("Num3: Len: %d | Value: %v\n", len(num3), num3)

	// 4. Without both number of elements and actual elements
	num4 := [...]int{}
	fmt.Printf("num4: Len: %d | Value: %v\n", len(num4), num4)

	// Accessing array element -
	fmt.Printf("Num elemet at 0: %v\n", num[0])

	// Different ways of iterating an array
	// 1. using for loop
	fmt.Printf("For Loop\n")
	for i := 0; i < len(num); i++ {
		fmt.Printf("Num elemet at %d: %v\n", i, num[i])
	}

	// 2. using for-range loop
	fmt.Printf("For-Range Loop\n")
	for i, value := range num {
		fmt.Printf("Num elemet at %d: %v\n", i, value)
	}

	// MultiDimensional Arrays
	// declaration
	multiArr := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// accessing multiDim array elements
	fmt.Println("MultiDimensional Arrays")
	for _, row := range multiArr {
		for _, value := range row {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}

}

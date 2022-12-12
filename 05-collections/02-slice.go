package main

import "fmt"

func main() {
	// Ways of creating slice
	// 1. Using the []<type>{} format
	// 2. Creating a slice from another slice or array
	// 3. Using make
	// 4. Using new

	// 1. Using the []<type>{} format
	num := []int{}
	num2 := []int{1, 2}
	fmt.Printf("Num: Len: %d | Cap: %d | Value: %v\n", len(num), cap(num), num)
	fmt.Printf("Num 2: Len: %d | Cap: %d | Value: %v\n", len(num2), cap(num2), num2)

	// 2. Creating a slice from another slice or array
	// imp point -
	// 1. length of newly created slice = (end–start)
	// 2. capacity of newly created slice = (length_of_array–start)
	nums := [5]int{1, 2, 3, 4, 5}

	// Both start and end
	num3 := nums[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("Num 3: Len: %d | Cap: %d | Value: %v\n", len(num3), cap(num3), num3)

	// Only start
	num4 := nums[2:]
	fmt.Println("Only start")
	fmt.Printf("Num 4: Len: %d | Cap: %d | Value: %v\n", len(num4), cap(num4), num4)

	// Only end
	num5 := nums[:3]
	fmt.Println("Only end")
	fmt.Printf("Num 5: Len: %d | Cap: %d | Value: %v\n", len(num5), cap(num5), num5)

	// None
	num6 := nums[:]
	fmt.Println("No start or end")
	fmt.Printf("Num 6: Len: %d | Cap: %d | Value: %v\n", len(num6), cap(num6), num6)

	// 3. Using make
	num7 := make([]int, 3, 5)
	fmt.Printf("Num 7: Len: %d | Cap: %d | Value: %v\n", len(num7), cap(num7), num7)

	// Different ways of iterating an array
	// 1. using for loop
	fmt.Printf("For Loop\n")
	for i := 0; i < len(num2); i++ {
		fmt.Printf("Num2 element at %d: %v\n", i, num2[i])
	}

	// 2. using for-range loop
	fmt.Printf("For-Range Loop\n")
	for i, value := range num2 {
		fmt.Printf("Num2 element at %d: %v\n", i, value)
	}

	// Appending two slices
	// ... - operator which means that the argument is a variadic parameter (accepts zero or more values of a specified type).
	num8 := []int{1, 2}
	num9 := []int{3, 4}
	numbers := append(num8, num9...)
	fmt.Println("Appending two slices")
	fmt.Printf("Numbers: Len: %d | Cap: %d | Value: %v\n", len(numbers), cap(numbers), numbers)

	// MultiDimensional Slices
	// declaration
	multiDSlice := make([][]int, 3)

	// for i := range multiDSlice {
	// 	multiDSlice[i] = make([]int, 3)
	// }

	multiDSlice[0] = []int{1, 2, 3}
	multiDSlice[1] = []int{4, 5, 6}
	multiDSlice[2] = []int{7, 8, 9}

	// accessing multiDim array elements
	fmt.Printf("MultiDimensional Slice: Rows: %d | Columns: %d | Total Elements: %d\n", len(multiDSlice), len(multiDSlice[0]), len(multiDSlice)*len(multiDSlice[0]))
	for _, row := range multiDSlice {
		for _, value := range row {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}

}

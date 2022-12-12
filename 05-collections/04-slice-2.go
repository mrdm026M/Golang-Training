package main

import (
	"fmt"
	"sort"
)

type employee struct {
	name string
	age  int
}

func main() {
	// creating slice of struct
	employees := make([]employee, 3)
	employees[0] = employee{name: "Dhruv", age: 21}
	employees[1] = employee{name: "Amish", age: 21}
	employees[2] = employee{name: "Mukesh", age: 25}

	for _, e := range employees {
		fmt.Printf("Employee Name: %s | Age: %d\n", e.name, e.age)
	}

	fmt.Printf("-----------------------------------------\n")

	// creating slice of map
	maps := make([]map[string]string, 3)
	map1 := make(map[string]string)
	map1["1"] = "a"
	map2 := make(map[string]string)
	map2["2"] = "b"
	map3 := make(map[string]string)
	map3["3"] = "c"

	maps[0] = map1
	maps[1] = map2
	maps[2] = map3

	for _, m := range maps {
		fmt.Println(m)
	}

	fmt.Printf("-----------------------------------------\n")

	// creating slice/array of bool
	boolExample := make([]bool, 3)
	boolExample[0] = false
	boolExample[1] = true
	boolExample[2] = false

	for _, c := range boolExample {
		fmt.Println(c)
	}

	fmt.Printf("-----------------------------------------\n")

	// Sorting a slice
	// 1. Full slice sorting in aesc order
	nums := []int{3, 2, 4, 1}

	fmt.Printf("Slice Nums before sorting: %d\n", nums)

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Printf("Slice Nums after sorting (ascesnding): %d\n", nums)

	// 2. Full slice sorting in desc order
	descNums := []int{3, 2, 4, 1}

	fmt.Printf("Slice descNums before sorting: %d\n", descNums)

	sort.Slice(descNums, func(i, j int) bool {
		return descNums[i] > descNums[j]
	})

	fmt.Printf("Slice descNums after sorting (descending): %d\n", descNums)

	// 3. partial sorting
	partialNums := []int{3, 4, 2, 1}
	fmt.Printf("Slice partialNums before sorting: %d\n", partialNums)
	start := 2
	end := 4

	sort.Slice(partialNums[start:end], func(i, j int) bool {
		return partialNums[start+i] < partialNums[start+j]
	})

	fmt.Printf("Slice partialNums after sorting: %d\n", partialNums)
}

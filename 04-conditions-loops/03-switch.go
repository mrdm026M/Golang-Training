package main

import "fmt"

// both switch statement and expression are optional
// possible scenario -
// 1. both switch statement and expression present
// 2. only switch statement present
// 3. only switch expression present
// 4. both switch statement and expression are not present

func main() {
	// 1. both switch statement and expression are present
	switch ch := 'b'; ch {
	case 'a':
		fmt.Println("a")
	case 'b', 'c':
		fmt.Println("b or c")
	default:
		fmt.Println("no matching character")
	}

	// 2. switch statement is present
	switch age := 25; {
	case age < 10:
		fmt.Println("a")
	case age > 10:
		fmt.Println("b")
	default:
		fmt.Println("no matching character")
	}

	// 3. switch expression is present
	ch := 'a'
	switch ch {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	default:
		fmt.Println("No matching character")
	}

	// 4. Both switch statement and expression are absent
	age := 25
	switch {
	case age < 10:
		fmt.Print(age - 10)
	case age > 10:
		fmt.Print(age + 10)
	default:
		fmt.Print(age)
	}

}
